package ip

import (
	"bytes"
	"golang.org/x/net/context"
	"io"
	"testing"
	"time"
	ip_capnp "zenhack.net/go/sandstorm/capnp/ip"
	util_capnp "zenhack.net/go/sandstorm/capnp/util"
)

type echoPort struct {
}

type echoWriter struct {
	out *util_capnp.ByteStream
}

func (p *echoPort) Connect(args ip_capnp.TcpPort_connect) error {
	downstream := args.Params.Downstream()
	w := util_capnp.ByteStream_ServerToClient(&echoWriter{&downstream})
	args.Results.SetUpstream(w)
	return nil
}

func (w *echoWriter) Write(args util_capnp.ByteStream_write) error {
	data, err := args.Params.Data()
	if err != nil {
		return err
	}
	w.out.Write(
		context.TODO(),
		func(args util_capnp.ByteStream_write_Params) error {

			args.SetData(data)
			return nil
		})
	return nil
}

func (w *echoWriter) Done(args util_capnp.ByteStream_done) error {
	w.out.Done(
		context.TODO(),
		func(args util_capnp.ByteStream_done_Params) error {
			return nil
		})
	return nil
}

func (w *echoWriter) ExpectSize(args util_capnp.ByteStream_expectSize) error {
	size := args.Params.Size()
	w.out.ExpectSize(
		context.TODO(),
		func(args util_capnp.ByteStream_expectSize_Params) error {
			args.SetSize(size)
			return nil
		})
	return nil
}

func TestConnect(t *testing.T) {
	port := ip_capnp.TcpPort_ServerToClient(&echoPort{})
	conn := connect(context.TODO(), port)
	expected := "Hello! How are you, Bob?"
	buf := &bytes.Buffer{}
	go func() {
		conn.Write([]byte(expected))
		// XXX: we need to allow time for the io.Copy below to copy
		// everything we've written, but we need to close conn before
		// it will return. It would be nice to find a better way to
		// deal with this, but right now we're just doing a sleep.
		// Technically there's a race condition here, but half a
		// second seems unlikely to fail frequently
		time.Sleep(time.Second / 2)
		conn.Close()
	}()
	io.Copy(buf, conn)
	actual := buf.String()
	if actual != expected {
		t.Fatalf("Expected %q but got %q.", expected, actual)
	}
}
