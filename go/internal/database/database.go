package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"zenhack.net/go/tempest/go/internal/config"
	"zenhack.net/go/util/exn"
)

const (
	DBPath = config.Localstatedir + "/sandstorm/sandstorm.sqlite3"
)

func Open() (DB, error) {
	sqlDB, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		return DB{}, err
	}
	return InitDB(sqlDB)
}

// Wrapper object around a SQL database.
type DB struct {
	sqlDB *sql.DB
}

// Transaction
type Tx struct {
	sqlTx *sql.Tx
}

// Initializes the database schema if needed, and returns a DB object.
func InitDB(sqlDB *sql.DB) (DB, error) {
	return exn.Try(func(throw func(error)) DB {
		tx, err := sqlDB.Begin()
		throw(err)
		defer tx.Rollback()
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS packages (
				-- 128-bit prefix of the sha256 hash of the spk file, hex encoded:
				id VARCHAR(32) PRIMARY KEY NOT NULL
			)`)
		throw(err)
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS accounts (
				id VARCHAR PRIMARY KEY,

				isAdmin BOOLEAN NOT NULL,

				profileDisplayName VARCHAR NOT NULL,
				profilePreferredHandle VARCHAR NOT NULL
				-- TODO: picture, identicon, pronouns
			)`)
		throw(err)
		_, err = tx.Exec(
			// TODO: research SSO support libraries for Go.
			`CREATE TABLE IF NOT EXISTS credentials (
				accountId VARCHAR NOT NULL REFERENCES accounts(id),
				-- Whether this credential is sufficient for logging
				-- in to the account:
				login BOOLEAN NOT NULL,
				-- The type of the credential. Currently always "dev".
				type VARCHAR NOT NULL,
				-- The name of the credential, within the type's naming system.
				-- e.g. for an email authentication system this would just be
				-- the email address.
				scopedId VARCHAR NOT NULL,
				PRIMARY KEY (type, scopedId)
			)`)
		throw(err)
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS grains (
				-- random base64 url-encoded:
				id VARCHAR(22) PRIMARY KEY NOT NULL,
				-- id of the package for this grain:
				packageId VARCHAR(32) NOT NULL REFERENCES packages(id),
				-- Human readable title chosen by the grain owner:
				title VARCHAR NOT NULL,
				ownerId VARCHAR NOT NULL REFERENCES accounts(id)
			)`)
		throw(err)
		_, err = tx.Exec(
			`-- A SturdyRef is a random token that grants access to a live capability.
			 --
			 -- This table stores information necessary to restore a sturdyRef. We
			 -- do not store the token itself; instead, we store a sha256 hash, so
			 -- that a database leak does not reveal the necessary information to
			 -- actually restore a sturdyref; the original token must be presented,
			 -- and it is not derivable from the contents of the database.
			CREATE TABLE IF NOT EXISTS sturdyRefs (
				-- raw sha256 hash of the token.
				sha256 BLOB PRIMARY KEY NOT NULL,
				-- The "owner" of this token; this determines who is allowed to
				-- restore the sturdyRef, and from where. The meaning of scopes
				-- are:
				--
				-- * grain-FOO must be restored via SandstormApi.restore(), from
				--   the grain with grainId FOO
				-- * userkeyring-FOO where FOO is in accounts.id: not restorable
				--   directly; logically each user has a "keyring" of capabilities
				--   reachable via APIs that require them to be logged in; some of
				--   those APIs look for this scope.
				owner VARCHAR NOT NULL,

				-- Unix timestamp after which this entry is invalid.
				expires INTEGER,

				-- If not null, this is a reference hosted by the grain with
				-- id 'grainId'. Otherwise, this is provided by the platform
				-- itself.
				grainId VARCHAR(22) REFERENCES grains(id) ON DELETE CASCADE,

				-- This is a single packed capnp message describing the object
				-- this sturdyRef refers to. If grainId is not null, then
				-- the root object of the message is the ObjectId returned
				-- by AppPersistent.save() (see grain.capnp). If this is null,
				-- then this sturdyRef refers to the root UiView exported by
				-- the grain.
				--
				-- When we first define a sturdyRef type provided by the platform,
				-- we will have to define a format for "system" object ids.
				objectId BLOB
			)`)
		throw(err)
		_, err = tx.Exec(
			`-- Extended fields for sturdyRefs that refer to uiViews.
			CREATE TABLE IF NOT EXISTS uiViewSturdyRefs (
				-- Hash of the token. The corresponding entry in sturdyRefs
				-- must have a non-null grainId.
				sha256 BLOB PRIMARY KEY NOT NULL REFERENCES sturdyRefs(id) ON DELETE CASCADE

				-- TODO: add permissions data and whatever else we need.
			)`)
		throw(err)
		throw(tx.Commit())
		return DB{sqlDB: sqlDB}
	})
}

func (db DB) Begin() (Tx, error) {
	tx, err := db.sqlDB.Begin()
	return Tx{sqlTx: tx}, err
}

func (tx Tx) Commit() error {
	return tx.sqlTx.Commit()
}

func (tx Tx) Rollback() error {
	return tx.sqlTx.Rollback()
}

// Closes the underlying database
func (db DB) Close() error {
	return db.sqlDB.Close()
}
