package database

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func TestPermissions(t *testing.T) {
	// Check that parsePermissions is the inverse of fmtPermissions
	assert.NoError(t, quick.Check(func(permissions []bool) bool {
		str := fmtPermissions(permissions)
		parsed, err := parsePermissions(str)
		if err != nil {
			return false
		}
		if len(parsed) != len(str) {
			return false
		}
		ret := true
		for i := range permissions {
			ret = ret && parsed[i] == permissions[i]
		}
		return ret
	}, nil))
}
