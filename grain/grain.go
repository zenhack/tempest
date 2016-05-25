// package grain provides helpers for acting as a sandstorm grain.
package grain // import "zenhack.net/go/sandstorm/grain"

// Copyright (c) 2016 Ian Denhardt <ian@zenhack.net>
//
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"golang.org/x/net/context"
	"net"
	"os"
	capnp "zenhack.net/go/sandstorm/capnp/grain"
	"zombiezen.com/go/capnproto2/rpc"
)

// Connect to the sandstorm API, give it a capability to uiview and return the api object.
//
// This assumes that file descriptor #3 is a socket with the API on the other end. This
// will be the case if the program was launched directly as a sandstorm app.
//
// TODO: explain error cases.
func ConnectAPI(ctx context.Context, uiview capnp.UiView_Server) (capnp.SandstormApi, error) {
	file := os.NewFile(3, "<sandstorm-api-socket>")
	conn, err := net.FileConn(file)
	if err != nil {
		return capnp.SandstormApi{}, err
	}
	uiviewClient := capnp.UiView_ServerToClient(uiview)
	transport := rpc.StreamTransport(conn)
	client := rpc.NewConn(transport, rpc.MainInterface(uiviewClient.Client)).Bootstrap(ctx)
	return capnp.SandstormApi{Client: client}, nil
}
