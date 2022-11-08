package legacy

import (
	"context"
	"encoding/binary"
	"os"
	"path/filepath"
	"strconv"

	"zenhack.net/go/util/exn"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Take a snapshot of a legacy sandstorm installation's mongo database.
// The format of the snapshot will be a directory contianing one file per collection,
// each of which is a sequence of bson values, prefixed with their length as a
// little-endian uint32.
//
// Parameters:
//
// mongoPort - Port on which mongo is listening
// passwdFile - File storing the mongo user password
// snapshotDir - Directory in which to store the snapshot
func Export(mongoPort int, passwdFile string, snapshotDir string) error {
	return exn.Try0(func(throw func(error)) {
		passwd, err := os.ReadFile(passwdFile)
		throw(err)

		throw(os.MkdirAll(snapshotDir, 0700))

		ctx := context.Background()
		client, err := mongo.Connect(
			ctx,
			options.Client().ApplyURI(
				"mongodb://sandstorm:"+string(passwd)+"@127.0.0.1:"+strconv.Itoa(mongoPort),
			),
		)
		throw(err)
		defer client.Disconnect(ctx)
		db := client.Database("meteor")
		collectionNames, err := db.ListCollectionNames(ctx, bson.D{})
		throw(err)
		for _, cname := range collectionNames {
			f, err := os.Create(filepath.Join(snapshotDir, cname))
			throw(err)

			c := db.Collection(cname)
			cur, err := c.Find(ctx, bson.D{})
			throw(err)
			for cur.Next(ctx) {
				throw(binary.Write(f, binary.LittleEndian, uint32(len(cur.Current))))
				_, err = f.Write(cur.Current[:])
				throw(err)
			}
		}
	})
}
