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

				profile_display_name VARCHAR NOT NULL,
				profile_preferred_handle VARCHAR NOT NULL
				-- TODO: picture, identicon, pronouns
			)`)
		throw(err)
		_, err = tx.Exec(
			// TODO: research SSO support libraries for Go.
			// TODO: add unique/primary key constraint for (type, scoped_id).
			`CREATE TABLE IF NOT EXISTS credentials (
				account_id VARCHAR NOT NULL REFERENCES accounts(id),
				-- Whether this credential is sufficient for logging
				-- in to the account:
				login BOOLEAN NOT NULL,
				-- The type of the credential. Currently always "dev".
				type VARCHAR NOT NULL,
				-- The name of the credential, within the type's naming system.
				-- e.g. for an email authentication system this would just be
				-- the email address.
				scoped_id VARCHAR NOT NULL
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

func (tx Tx) AddPackage(pkgId string) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO packages(id) VALUES (?)`,
		pkgId,
	)
	return err
}

type NewGrain struct {
	GrainId string
	PkgId   string
	OwnerId string
	Title   string
}

type NewAccount struct {
	Id      string
	IsAdmin bool
	Profile struct {
		DisplayName     string
		PreferredHandle string
	}
}

type NewCredential struct {
	AccountId string
	Login     bool
	Type      string
	ScopedId  string
}

func (tx Tx) AddAccount(a NewAccount) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO accounts
			( id
			, isAdmin
			, profile_display_name
			, profile_preferred_handle
			)
			VALUES (?, ?, ?, ?)`,
		a.Id,
		a.IsAdmin,
		a.Profile.DisplayName,
		a.Profile.PreferredHandle,
	)
	return err
}

func (tx Tx) AddCredential(c NewCredential) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO credentials
			( account_id
			, login
			, type
			, scoped_id
			)
			VALUES (?, ?, ?, ?)`,
		c.AccountId,
		c.Login,
		c.Type,
		c.ScopedId,
	)
	return err
}

func (tx Tx) AddGrain(g NewGrain) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO grains(id, packageId, title, ownerId) VALUES (?, ?, ?, ?)`,
		g.GrainId, g.PkgId, g.Title, g.OwnerId,
	)
	return err
}

// Returns the package id for the specified grain
func (tx Tx) GetGrainPackageId(grainId string) (string, error) {
	row := tx.sqlTx.QueryRow("SELECT packageId FROM grains WHERE id = ?", grainId)
	var result string
	err := row.Scan(&result)
	return result, err
}

// Closes the underlying database
func (db DB) Close() error {
	return db.sqlDB.Close()
}
