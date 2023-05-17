package servermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/session"
)

type uiViewControllerImpl struct {
	GrainID types.GrainID
	Session session.UserSession
	DB      database.DB
}

func (c uiViewControllerImpl) MakeSharingToken(ctx context.Context, p external.UiView_Controller_makeSharingToken) error {
	wantPerms, err := p.Args().Permissions()
	if err != nil {
		return err
	}

	note, err := p.Args().Note()
	if err != nil {
		return err
	}

	results, err := p.AllocResults()
	if err != nil {
		return err
	}

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	accountID, err := tx.GetCredentialAccount(c.Session.Credential)
	if err != nil {
		return err
	}
	perms, err := tx.AccountGrainPermissions(accountID, c.GrainID)
	if err != nil {
		return err
	}
	if len(perms) < wantPerms.Len() {
		perms = perms[:wantPerms.Len()]
	}
	for i := range perms {
		perms[i] = perms[i] && wantPerms.At(i)
	}
	token, err := tx.NewSharingToken(c.GrainID, accountID, perms, note)
	if err != nil {
		return err
	}
	return results.SetToken(token)
}
