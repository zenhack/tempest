package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoleEncompasses(t *testing.T) {
	cases := []struct {
		Lhs, Rhs Role
		Result   bool
	}{
		{
			Lhs:    RoleVisitor,
			Rhs:    RoleVisitor,
			Result: true,
		},
		{
			Lhs:    RoleVisitor,
			Rhs:    RoleUser,
			Result: false,
		},
	}
	for _, c := range cases {
		testCase := c
		t.Run(string(c.Lhs)+" encompasses "+string(c.Rhs), func(t *testing.T) {
			require.Equal(t, testCase.Result, testCase.Lhs.Encompasses(testCase.Rhs))
		})
	}
}
