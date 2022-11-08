package main

import (
	"flag"

	"zenhack.net/go/sandstorm-next/go/internal/database/legacy"
	"zenhack.net/go/util"
)

var (
	mongoPort  = flag.Int("mongo-port", 6081, "Port on which mongo is listening")
	passwdFile = flag.String("passwd-file", "/opt/sandstorm/var/mongo/passwd",
		"File storing the mongo user password")
	snapshotDir = flag.String("snapshot-dir", "./mongo-snapshot",
		"Directory in which to store a temporary snapshot")
)

func main() {
	flag.Parse()
	util.Chkfatal(legacy.Export(*mongoPort, *passwdFile, *snapshotDir))
}
