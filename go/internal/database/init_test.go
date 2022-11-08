package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Try initializing the database, make sure it succeeds
func TestInit(t *testing.T) {
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)
	db, err := InitDB(sqlDB)
	assert.Nil(t, err)
	assert.Nil(t, db.Close())
}
