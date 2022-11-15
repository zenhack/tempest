package legacy

import (
	"database/sql"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"go.mongodb.org/mongo-driver/bson"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/util/exn"
)

type iter struct {
	r io.Reader
}

func (it iter) readValue() (bson.Raw, error) {
	var lenBuf [4]byte
	_, err := io.ReadFull(it.r, lenBuf[:])
	if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint32(lenBuf[:])
	ret := make([]byte, int(length))
	_, err = io.ReadFull(it.r, ret)
	return bson.Raw(ret), err
}

func Import(sqlitePath, snapshotDir string) error {
	return exn.Try0(func(throw func(error)) {
		sqliteDB, err := sql.Open("sqlite3", sqlitePath)
		throw(err)
		db, err := database.InitDB(sqliteDB)
		throw(err)
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()
		throw(importPackages(snapshotDir, tx))
		throw(importGrains(snapshotDir, tx))
		throw(tx.Commit())
	})
}

func eachEntry(snapshotDir, collection string, tx database.Tx, fn func(bson.Raw) error) error {
	return exn.Try0(func(throw func(error)) {
		f, err := os.Open(filepath.Join(snapshotDir, "packages"))
		throw(err)
		defer f.Close()
		it := iter{r: f}
		for {
			raw, err := it.readValue()
			if err == io.EOF {
				return
			}
			throw(err)
			throw(fn(raw))
		}
	})
}

func importPackages(snapshotDir string, tx database.Tx) error {
	return eachEntry(snapshotDir, "packages", tx, func(raw bson.Raw) error {
		return exn.Try0(func(throw func(error)) {
			elts, err := raw.Elements()
			throw(err)

			for _, e := range elts {
				if e.Key() == "_id" {
					id := e.Value().StringValue()
					throw(tx.AddPackage(id))
					break
				}
			}
		})
	})
}

func importGrains(snapshotDir string, tx database.Tx) error {
	return eachEntry(snapshotDir, "grains", tx, func(raw bson.Raw) error {
		return exn.Try0(func(throw func(error)) {
			elts, err := raw.Elements()
			throw(err)

			var grainId, pkgId, title string

			for _, e := range elts {
				switch e.Key() {
				case "_id":
					grainId = e.Value().StringValue()
				case "packageId":
					pkgId = e.Value().StringValue()
				case "title":
					title = e.Value().StringValue()
				}
			}

			throw(tx.AddGrain(grainId, pkgId, title))
		})
	})
}
