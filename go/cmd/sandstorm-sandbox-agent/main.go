// `sandstorm-sandbox-agent` runs inside the grain's sandbox, and is the first
// program executed during grain startup. Its file descriptor #3 is a socket
// over which we can speak capnp to the sandstorm server outside the sandbox.
//
// Any APIs available to the grain which don't actually need privileges the grain
// doesn't have should ideally be implemented here; this helps us minimize attack
// surface.
package main

import (
	"fmt"
	"io/ioutil"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/sandstorm-next/go/internal/util"
	"zenhack.net/go/sandstorm/capnp/spk"
)

func main() {
	data, err := ioutil.ReadFile("/sandstorm-manifest")
	util.Chkfatal(err)
	msg, err := capnp.Unmarshal(data)
	util.Chkfatal(err)
	manifest, err := spk.ReadRootManifest(msg)
	util.Chkfatal(err)
	appTitle, err := manifest.AppTitle()
	util.Chkfatal(err)
	text, err := appTitle.DefaultText()
	util.Chkfatal(err)
	fmt.Println("App title: ", text)
}
