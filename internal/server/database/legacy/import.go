package legacy

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"capnproto.org/go/capnp/v3"
	_ "github.com/mattn/go-sqlite3"

	"go.mongodb.org/mongo-driver/bson"
	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/database"
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
		throw(importUsers(snapshotDir, tx))
		throw(importPackages(snapshotDir, tx))
		throw(importGrains(snapshotDir, tx))
		throw(tx.Commit())
	})
}

func eachEntry(snapshotDir, collection string, tx database.Tx, fn func(bson.Raw)) error {
	return exn.Try0(func(throw func(error)) {
		f, err := os.Open(filepath.Join(snapshotDir, collection))
		throw(err)
		defer f.Close()
		it := iter{r: f}
		for {
			raw, err := it.readValue()
			if err == io.EOF {
				return
			}
			throw(err)
			fn(raw)
		}
	})
}

type user struct {
	Id      string
	Type    string
	IsAdmin bool
	Profile struct {
		DisplayName     string
		PreferredHandle string
	}

	Services struct {
		Dev struct {
			Name    string
			IsAdmin bool
		}
	}

	LoginCredentials    []string
	NonloginCredentials []string
}

// Decodes a list of credentials in the form of loginCredentials or nonloginCredentials.
func decodeCredentialList(e bson.RawElement) ([]string, error) {
	return exn.Try(func(throw func(error)) []string {
		var ret []string
		vals, err := e.Value().Array().Values()
		throw(err)
		for _, v := range vals {
			elts, err := v.Document().Elements()
			throw(err)
			for _, e := range elts {
				if e.Key() == "id" {
					ret = append(ret, e.Value().StringValue())
				}
			}
		}
		return ret
	})
}

func decodeUser(raw bson.Raw) (user, error) {
	return exn.Try(func(throw func(error)) (ret user) {
		elts, err := raw.Elements()
		throw(err)

		for _, e := range elts {
			switch e.Key() {
			case "_id":
				ret.Id = e.Value().StringValue()
			case "type":
				ret.Type = e.Value().StringValue()
			case "isAdmin":
				ret.IsAdmin = e.Value().Boolean()
			case "profile":
				pelts, err := e.Value().Document().Elements()
				throw(err)
				for _, pe := range pelts {
					switch pe.Key() {
					case "name":
						ret.Profile.DisplayName = pe.Value().StringValue()
					case "handle":
						ret.Profile.PreferredHandle = pe.Value().StringValue()
					}
				}
			case "loginCredentials":
				ret.LoginCredentials, err = decodeCredentialList(e)
				throw(err)
			case "nonloginCredentials":
				ret.NonloginCredentials, err = decodeCredentialList(e)
				throw(err)
			case "services":
				selts, err := e.Value().Document().Elements()
				throw(err)
				for _, se := range selts {
					switch se.Key() {
					case "dev":
						delts, err := se.Value().Document().Elements()
						throw(err)
						for _, de := range delts {
							switch de.Key() {
							case "name":
								ret.Services.Dev.Name = de.Value().StringValue()
							case "isAdmin":
								ret.Services.Dev.IsAdmin = de.Value().Boolean()
							}
						}
					default:
						// TODO: handle github, google, email, etc.
					}
				}
			}
		}
		return
	})
}

type credentialOwner struct {
	accountId string
	login     bool
}

func importUsers(snapshotDir string, tx database.Tx) error {
	// Mapping from credential ids to info about their owner:
	credentialOwners := make(map[string]credentialOwner)

	return exn.Try0(func(throw func(error)) {
		// We do this in two passes, accounts first, then credentials.
		throw(eachEntry(snapshotDir, "users", tx, func(raw bson.Raw) {
			u, err := decodeUser(raw)
			throw(err)

			if u.Type != "account" {
				return
			}
			throw(tx.AddAccount(database.NewAccount{
				Id:      u.Id,
				IsAdmin: u.IsAdmin,
				Profile: u.Profile,
			}))
			// Store these for lookup during the second pass:
			for _, credId := range u.LoginCredentials {
				credentialOwners[credId] = credentialOwner{
					accountId: u.Id,
					login:     true,
				}
			}
			for _, credId := range u.NonloginCredentials {
				credentialOwners[credId] = credentialOwner{
					accountId: u.Id,
					login:     false,
				}
			}
		}))
		throw(eachEntry(snapshotDir, "users", tx, func(raw bson.Raw) {
			u, err := decodeUser(raw)
			throw(err)

			if u.Type != "credential" {
				return
			}
			owner := credentialOwners[u.Id]
			var entry database.NewCredential
			entry.AccountId = owner.accountId
			entry.Login = owner.login
			if u.Services.Dev.Name != "" {
				entry.Credential = types.Credential{
					Type:     "dev",
					ScopedId: u.Services.Dev.Name,
				}
				throw(tx.AddCredential(entry))
			} else {
				fmt.Printf("TODO: add support for other credential types (skipping)")
			}
		}))
	})
}

func importPackages(snapshotDir string, tx database.Tx) error {
	return exn.Try0(func(throw func(error)) {
		throw(eachEntry(snapshotDir, "packages", tx, func(raw bson.Raw) {
			elts, err := raw.Elements()
			throw(err)

			for _, e := range elts {
				if e.Key() == "_id" {
					id := e.Value().StringValue()
					path := config.Localstatedir +
						"/sandstorm/apps/" +
						id +
						"/sandstorm-manifest"
					buf, err := os.ReadFile(path)
					throw(err)
					msg, err := capnp.Unmarshal(buf)
					throw(err)
					manifest, err := spk.ReadRootManifest(msg)
					throw(err)
					throw(tx.AddPackage(database.Package{
						Id:       id,
						Manifest: manifest,
					}))
					break
				}
			}
		}))
	})
}

func importGrains(snapshotDir string, tx database.Tx) error {
	return exn.Try0(func(throw func(error)) {
		throw(eachEntry(snapshotDir, "grains", tx, func(raw bson.Raw) {
			elts, err := raw.Elements()
			throw(err)

			var grain database.NewGrain

			for _, e := range elts {
				switch e.Key() {
				case "_id":
					grain.GrainId = e.Value().StringValue()
				case "packageId":
					grain.PkgId = e.Value().StringValue()
				case "title":
					grain.Title = e.Value().StringValue()
				case "userId":
					grain.OwnerId = e.Value().StringValue()
				}
			}

			throw(tx.AddGrain(grain))
		}))
	})
}
