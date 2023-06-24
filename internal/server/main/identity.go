package servermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/identity"
)

// Implementation of identity.Identity. getProfile() just returns
// the stored profile.
type identityImpl struct {
	profile identity.Profile
}

func (id identityImpl) GetProfile(ctx context.Context, p identity.Identity_getProfile) error {
	res, err := p.AllocResults()
	if err != nil {
		return err
	}
	return res.SetProfile(id.profile)
}
