package servermain

import (
	"os"

	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/util"
)

// Initializes directories needed at runtime
func initStorage() {
	util.Chkfatal(os.MkdirAll(config.TempDir, 0700))
	util.Chkfatal(os.MkdirAll(config.PackagesDir, 0700))
	util.Chkfatal(os.MkdirAll(config.GrainsDir, 0700))
}
