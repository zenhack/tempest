package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the GetCredentialAccount method
func TestGetCredentialAccount(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	tx, err := db.Begin()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, tx.Commit())
	}()
	addTestUsers(t, tx)

	id, err := tx.GetCredentialAccount("dev", "Alice Dev Admin")
	assert.NoError(t, err)
	assert.Equal(t, id, "id_alice")

	id, err = tx.GetCredentialAccount("dev", "Bob Dev User")
	assert.NoError(t, err)
	assert.Equal(t, id, "id_bob")
}

// addTestUsers adds some user accounts to the database.
func addTestUsers(t *testing.T, tx Tx) {
	accounts := []NewAccount{
		{
			Id:      "id_alice",
			IsAdmin: true,
			Profile: Profile{
				DisplayName:     "Alice",
				PreferredHandle: "alice",
			},
		},
		{
			Id:      "id_bob",
			IsAdmin: false,
			Profile: Profile{
				DisplayName:     "Bob",
				PreferredHandle: "bob",
			},
		},
	}

	credentials := []NewCredential{
		{
			AccountId: "id_alice",
			Login:     true,
			Type:      "dev",
			ScopedId:  "Alice Dev Admin",
		},
		{
			AccountId: "id_bob",
			Login:     true,
			Type:      "dev",
			ScopedId:  "Bob Dev User",
		},
	}

	for _, a := range accounts {
		assert.NoError(t, tx.AddAccount(a), "Adding account: ", a)
	}
	for _, c := range credentials {
		assert.NoError(t, tx.AddCredential(c), "Adding credential: ", c)
	}
}

func openTestDB(t *testing.T) DB {
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	db, err := InitDB(sqlDB)
	assert.NoError(t, err)
	return db
}
