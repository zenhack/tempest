package main

import (
	"encoding/base64"
	"flag"

	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/tokenutil"
	"zenhack.net/go/util"
)

var (
	typ      = flag.String("type", "", "credential type to use")
	scopedID = flag.String("id", "", "type-specific credential id")
	roleStr  = flag.String("role", string(types.RoleUser), "role the user should have")
)

func main() {
	accountID := base64.RawStdEncoding.EncodeToString(tokenutil.GenToken()[:16])
	flag.Parse()
	role := types.Role(*roleStr)
	if !role.IsValid() {
		panic("Invalid role: " + role)
	}

	db, err := database.Open()
	util.Chkfatal(err)
	defer db.Close()
	tx, err := db.Begin()
	util.Chkfatal(err)
	defer tx.Rollback()
	util.Chkfatal(tx.AddAccount(database.NewAccount{
		ID:   accountID,
		Role: role,
	}))
	util.Chkfatal(tx.AddCredential(database.NewCredential{
		AccountID: accountID,
		Login:     true,
		Credential: types.Credential{
			Type:     types.CredentialType(*typ),
			ScopedID: *scopedID,
		},
	}))
	util.Chkfatal(tx.Commit())
}
