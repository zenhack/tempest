package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"zenhack.net/go/tempest/internal/common/types"
)

// Test the GetCredentialAccount method
func TestGetCredentialAccount(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		addTestData(t, tx)

		id, err := tx.GetCredentialAccount(types.Credential{
			Type:     "dev",
			ScopedID: "Alice Dev Admin",
		})
		assert.NoError(t, err)
		assert.Equal(t, id, "id_alice")

		id, err = tx.GetCredentialAccount(types.Credential{
			"dev",
			"Bob Dev User",
		})
		assert.NoError(t, err)
		assert.Equal(t, id, "id_bob")
	})
}

func TestGetUiViews(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		addTestData(t, tx)

		views, err := tx.GetCredentialUiViews(types.Credential{
			Type:     "dev",
			ScopedID: "Alice Dev Admin",
		})
		assert.NoError(t, err)

		assert.Equal(t, 1, len(views))
		assert.Equal(t, GrainInfo{
			ID:    "grain123",
			Title: "Example Grain",
			Owner: "id_alice",
		}, views[0].Grain)
		assert.Equal(t, 0, len(views[0].Permissions))
	})
}

// addTestData populates the database with some initial data.
func addTestData(t *testing.T, tx Tx) {
	accounts := []NewAccount{
		{
			ID:   "id_alice",
			Role: types.RoleAdmin,
			Profile: Profile{
				DisplayName:     "Alice",
				PreferredHandle: "alice",
			},
		},
		{
			ID:   "id_bob",
			Role: types.RoleUser,
			Profile: Profile{
				DisplayName:     "Bob",
				PreferredHandle: "bob",
			},
		},
	}

	credentials := []NewCredential{
		{
			AccountID: "id_alice",
			Login:     true,
			Credential: types.Credential{
				Type:     "dev",
				ScopedID: "Alice Dev Admin",
			},
		},
		{
			AccountID: "id_bob",
			Login:     true,
			Credential: types.Credential{
				Type:     "dev",
				ScopedID: "Bob Dev User",
			},
		},
	}

	packages := []Package{
		{
			ID: "abcdef",
			// FIXME: Manifest
		},
	}

	grains := []NewGrain{
		{
			GrainID: "grain123",
			PkgID:   "abcdef",
			OwnerID: "id_alice",
			Title:   "Example Grain",
		},
	}

	for _, a := range accounts {
		assert.NoError(t, tx.AddAccount(a), "Adding account: ", a)
	}
	for _, c := range credentials {
		assert.NoError(t, tx.AddCredential(c), "Adding credential: ", c)
	}
	for _, p := range packages {
		assert.NoError(t, tx.AddPackage(p), "Adding package: ", p)
		assert.NoError(t, tx.ReadyPackage(p.ID), "Readying package: ", p)
	}
	for _, g := range grains {
		assert.NoError(t, tx.AddGrain(g), "Adding grain: ", g)
	}
}

// testWithTx runs f in a fresh transaction, then commits. It fails
// the test if the commit fails.
func testWithTx(t *testing.T, f func(tx Tx)) {
	db := openTestDB(t)
	defer db.Close()
	tx, err := db.Begin()
	assert.NoError(t, err)
	defer mustCommit(t, tx)
	f(tx)
}

func mustCommit(t *testing.T, tx Tx) {
	assert.NoError(t, tx.Commit())
}

func openTestDB(t *testing.T) DB {
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	db, err := InitDB(sqlDB)
	assert.NoError(t, err)
	return db
}
