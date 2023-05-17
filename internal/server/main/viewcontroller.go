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

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	perms, err := tx.CredentialGrainPermissions(c.Session.Credential, c.GrainID)
	if err != nil {
		return err
	}
	if len(perms) < wantPerms.Len() {
		perms = perms[:wantPerms.Len()]
	}
	for i := range perms {
		perms[i] = perms[i] && wantPerms.At(i)
	}
	panic("TODO: finish")
}
