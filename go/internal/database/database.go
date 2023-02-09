package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"zenhack.net/go/tempest/go/internal/config"
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
