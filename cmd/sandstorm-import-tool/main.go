package main

import (
	"flag"
	"fmt"
	"os"

	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/database/legacy"
	"zenhack.net/go/util"
)

var (
	mongoPort  = flag.Int("mongo-port", 6081, "Port on which mongo is listening (export only)")
	passwdFile = flag.String("passwd-file", "/opt/sandstorm/var/mongo/passwd",
		"File storing the mongo user password (export only)")
	snapshotDir = flag.String("snapshot-dir", "./mongo-snapshot",
		"Directory in which to store a temporary snapshot")
	sqlitePath = flag.String("sqlite-path", database.DBPath, "path to sqlite database (import only)")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: sandstorm-import-tool [ flags ] <export | import>")
		os.Exit(1)
	}
	switch args[0] {
	case "export":
		util.Chkfatal(legacy.Export(*mongoPort, *passwdFile, *snapshotDir))
	case "import":
		util.Chkfatal(legacy.Import(*sqlitePath, *snapshotDir))
	}
}
