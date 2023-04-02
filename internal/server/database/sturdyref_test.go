package database

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Save and restore a sturdyRef, and verify that the value is the same.
func TestSturdyRefSaveRestore(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		key := SturdyRefKey{
			Token:     GenToken(),
			OwnerType: "grain",
			Owner:     "grain123",
		}
		value := SturdyRefValue{
			Expires: time.Unix(math.MaxInt64, 0), // Effectively never.
			GrainID: "grain456",
			// TODO: fill something in for the struct.
		}

		_, err := tx.SaveSturdyRef(key, value)
		require.NoError(t, err)

		restored, err := tx.RestoreSturdyRef(key)
		require.NoError(t, err)
		require.Equal(t, value, restored)
	})
}

// Try to restore an expired sturdyRef; it should fail.
func TestSturdyRefRestoreExpired(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		key := SturdyRefKey{
			Token:     GenToken(),
			OwnerType: "grain",
			Owner:     "grain123",
		}
		value := SturdyRefValue{
			Expires: time.Now().Add(-1 * time.Second),
			GrainID: "grain456",
		}

		_, err := tx.SaveSturdyRef(key, value)
		require.NoError(t, err)

		_, err = tx.RestoreSturdyRef(key)
		require.NotNil(t, err, "restoring expired sturdyRef should fail")
	})
}
