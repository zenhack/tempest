package servermain

import (
	"os"

	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/util"
)

// Initializes directories needed at runtime
func initStorage() {
	util.Chkfatal(os.MkdirAll(config.Localstatedir+"/tmp/tempest", 0700))
	util.Chkfatal(os.MkdirAll(config.Localstatedir+"/sandstorm/apps", 0700))
	util.Chkfatal(os.MkdirAll(config.Localstatedir+"/sandstorm/grains", 0700))
}
