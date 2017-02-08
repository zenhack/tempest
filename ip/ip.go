package ip

import (
	"encoding/binary"
	"golang.org/x/net/context"
	"io"
	"net"
	capnp "zenhack.net/go/sandstorm/capnp/ip"
	capnp_util "zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/util"
)

type IpNetworkDialer struct {
	Ctx       context.Context
	IpNetwork capnp.IpNetwork
}

func (d *IpNetworkDialer) Dial(network, addr string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		hostPromise, portNum, err := d.getAddr(network, addr)
		if err != nil {
			return nil, err
		}
		portPromise := hostPromise.GetTcpPort(d.Ctx,
			func(p capnp.IpRemoteHost_getTcpPort_Params) error {
				p.SetPortNum(portNum)
				return nil
			}).Port()
		return connect(d.Ctx, portPromise), nil

	default:
		panic("TODO")
	}
}

func (d *IpNetworkDialer) getAddr(network, addr string) (capnp.IpRemoteHost, uint16, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return capnp.IpRemoteHost{}, 0, err
	}
	portNum, err := net.LookupPort(network, port)
	if err != nil {
		return capnp.IpRemoteHost{}, 0, err
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
	return hostPromise, uint16(portNum), nil
}

func connect(ctx context.Context, port capnp.TcpPort) net.Conn {
	clientConn, serverConn := net.Pipe()
	toServerBS := port.Connect(ctx, func(p capnp.TcpPort_connect_Params) error {
		p.SetDownstream(capnp_util.ByteStream_ServerToClient(
			&util.WriteCloserByteStream{WC: serverConn},
		))
		return nil
	}).Upstream()
	go io.Copy(
		&util.ByteStreamWriteCloser{Ctx: ctx, Obj: toServerBS},
		clientConn,
	)
	return clientConn
}
