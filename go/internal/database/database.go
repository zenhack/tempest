package database

import (
	"database/sql"

	"zenhack.net/go/util/exn"
)

// Wrapper object around a SQL database.
type DB struct {
	sqlDB *sql.DB
}

// Initializes the database schema if needed, and returns a DB object.
func InitDB(sqlDB *sql.DB) (DB, error) {
	return exn.Try(func(throw func(error)) DB {
		_, err := sqlDB.Exec(
			`CREATE TABLE IF NOT EXISTS packages (
				-- 128-bit prefix of the sha256 hash of the spk file, hex encoded:
				id VARCHAR(32) PRIMARY KEY NOT NULL
			)`)
		throw(err)
		_, err = sqlDB.Exec(
			`CREATE TABLE IF NOT EXISTS grains (
				-- random base64 url-encoded:
				id VARCHAR(22) PRIMARY KEY NOT NULL,
				-- id of the package for this grain:
				packageId VARCHAR(32) NOT NULL REFERENCES packages(id),
				-- Human readable title chosen by the grain owner:
				title VARCHAR NOT NULL
			)`)
		throw(err)
		return DB{sqlDB: sqlDB}
	})
}

// Returns the package id for the specified grain
func (db DB) GetGrainPackageId(grainId string) (string, error) {
	row := db.sqlDB.QueryRow("SELECT packageId FROM grains WHERE id = ?", grainId)
	var result string
	err := row.Scan(&result)
	return result, err
}

// Closes the underlying database
func (db DB) Close() error {
	return db.sqlDB.Close()
}
