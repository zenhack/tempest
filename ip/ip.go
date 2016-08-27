package ip

import (
	"encoding/binary"
	"golang.org/x/net/context"
	"io"
	"net"
	capnp "zenhack.net/go/sandstorm/capnp/ip"
	util_capnp "zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/internal/iocommon"
	"zenhack.net/go/sandstorm/util"
)

type IpNetworkDialer struct {
	Ctx       context.Context
	IpNetwork capnp.IpNetwork
}

func (d *IpNetworkDialer) Dial(network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	portNum, err := net.LookupPort(network, port)
	if err != nil {
		return nil, err
	}
	if portNum < 0 || portNum >= 1<<16 {
		// Invaid port number. TODO: return an error.
	}
	ipAddr := net.ParseIP(host)
	var hostPromise capnp.IpRemoteHost
	if ipAddr == nil {
		// not a valid address; assume it's a hostname.
		hostPromise = d.IpNetwork.GetRemoteHostByName(d.Ctx,
			func(p capnp.IpNetwork_getRemoteHostByName_Params) error {
				p.SetAddress(host)
				return nil
			}).Host()
	} else {
		ipAddr = ipAddr.To16()
		hostPromise = d.IpNetwork.GetRemoteHost(d.Ctx,
			func(p capnp.IpNetwork_getRemoteHost_Params) error {
				capnpAddr, err := p.NewAddress()
				if err != nil {
					return err
				}
				capnpAddr.SetUpper64(binary.BigEndian.Uint64(ipAddr[:8]))
				capnpAddr.SetLower64(binary.BigEndian.Uint64(ipAddr[8:]))
				return nil
			}).Host()
	}
	switch network {
	case "tcp", "tcp4", "tcp6":
		portPromise := hostPromise.GetTcpPort(d.Ctx,
			func(p capnp.IpRemoteHost_getTcpPort_Params) error {
				p.SetPortNum(uint16(portNum))
				return nil
			}).Port()
		fromUpstream, toDownstream := io.Pipe()
		toUpstream := util.ByteStreamWriteCloser{
			Ctx: d.Ctx,
			Obj: portPromise.Connect(d.Ctx,
				func(p capnp.TcpPort_connect_Params) error {
					downstream := util_capnp.ByteStream_ServerToClient(
						&util.WriteCloserByteStream{toDownstream},
					)
					p.SetDownstream(downstream)
					return nil
				}).Upstream(),
		}
		return &iocommon.RWCConn{
			ReadWriteCloser: iocommon.MergedRWC{
				Reader: fromUpstream,
				Writer: toUpstream,
				Closer: iocommon.MultiCloser(fromUpstream, toUpstream),
			},
			Local: &iocommon.HardCodedAddr{
				Net:  network,
				Addr: "localhost",
			},
			Remote: &iocommon.HardCodedAddr{
				Net:  network,
				Addr: addr,
			},
		}, nil
	default:
		panic("TODO")
	}
}
