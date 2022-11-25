package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"zenhack.net/go/sandstorm-next/go/internal/config"
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
			// TODO: add unique/primary key constraint for (type, scopedId).
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
				scopedId VARCHAR NOT NULL
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
