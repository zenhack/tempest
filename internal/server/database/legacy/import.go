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
	ID      string
	Type    string
	Role    database.Role
	Profile struct {
		DisplayName     string
		PreferredHandle string
	}

	Services struct {
		Dev struct {
			Name    string
			IsAdmin bool
		}
		Email struct {
			Email string
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

		ret.Role = database.RoleVisitor

		for _, e := range elts {
			switch e.Key() {
			case "_id":
				ret.ID = e.Value().StringValue()
			case "type":
				ret.Type = e.Value().StringValue()
			case "isAdmin":
				ret.Role = database.RoleAdmin
			case "signupKey":
				// The presence of this key means the user is not a visitor.
				// isAdmin takes precedence though.
				if ret.Role == database.RoleVisitor {
					ret.Role = database.RoleUser
				}
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
					case "email":
						eelts, err := se.Value().Document().Elements()
						throw(err)
						for _, ee := range eelts {
							switch ee.Key() {
							case "email":
								ret.Services.Email.Email = ee.Value().StringValue()
							}
						}
					default:
						// TODO: handle github, google, etc.
					}
				}
			}
		}
		return
	})
}

type credentialOwner struct {
	accountID string
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
				ID:      u.ID,
				Role:    u.Role,
				Profile: u.Profile,
			}))
			// Store these for lookup during the second pass:
			for _, credID := range u.LoginCredentials {
				credentialOwners[credID] = credentialOwner{
					accountID: u.ID,
					login:     true,
				}
			}
			for _, credID := range u.NonloginCredentials {
				credentialOwners[credID] = credentialOwner{
					accountID: u.ID,
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
			owner := credentialOwners[u.ID]
			var entry database.NewCredential
			entry.AccountID = owner.accountID
			entry.Login = owner.login
			if u.Services.Dev.Name != "" {
				entry.Credential = types.Credential{
					Type:     types.DevCredential,
					ScopedID: u.Services.Dev.Name,
				}
			} else if u.Services.Email.Email != "" {
				entry.Credential = types.Credential{
					Type:     types.EmailCredential,
					ScopedID: u.Services.Email.Email,
				}
			} else {
				fmt.Println("TODO: add support for other credential types (skipping)")
				return
			}
			throw(tx.AddCredential(entry))
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
					id := types.ID[database.Package](e.Value().StringValue())
					path := config.Localstatedir +
						"/sandstorm/apps/" +
						string(id) +
						"/sandstorm-manifest"
					buf, err := os.ReadFile(path)
					throw(err)
					msg, err := capnp.Unmarshal(buf)
					throw(err)
					manifest, err := spk.ReadRootManifest(msg)
					throw(err)
					throw(tx.AddPackage(database.Package{
						ID:       id,
						Manifest: manifest,
					}))
					throw(tx.ReadyPackage(id))
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
					grain.GrainID = types.GrainID(e.Value().StringValue())
				case "packageId":
					grain.PkgID = types.ID[database.Package](e.Value().StringValue())
				case "title":
					grain.Title = e.Value().StringValue()
				case "userId":
					grain.OwnerID = e.Value().StringValue()
				}
			}

			throw(tx.AddGrain(grain))
		}))
	})
}
