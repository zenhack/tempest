package database

import (
	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/packed"
	"zenhack.net/go/tempest/capnp/grain"
)

func (tx Tx) AddPackage(pkgId string) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO packages(id) VALUES (?)`,
		pkgId,
	)
	return err
}

type NewGrain struct {
	GrainId string
	PkgId   string
	OwnerId string
	Title   string
}

type NewAccount struct {
	Id      string
	IsAdmin bool
	Profile struct {
		DisplayName     string
		PreferredHandle string
	}
}

type NewCredential struct {
	AccountId string
	Login     bool
	Type      string
	ScopedId  string
}

func (tx Tx) AddAccount(a NewAccount) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO accounts
			( id
			, isAdmin
			, profileDisplayName
			, profilePreferredHandle
			)
			VALUES (?, ?, ?, ?)`,
		a.Id,
		a.IsAdmin,
		a.Profile.DisplayName,
		a.Profile.PreferredHandle,
	)
	return err
}

func (tx Tx) AddCredential(c NewCredential) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO credentials
			( accountId
			, login
			, type
			, scopedId
			)
			VALUES (?, ?, ?, ?)`,
		c.AccountId,
		c.Login,
		c.Type,
		c.ScopedId,
	)
	return err
}

func (tx Tx) AddGrain(g NewGrain) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO grains(id, packageId, title, ownerId) VALUES (?, ?, ?, ?)`,
		g.GrainId, g.PkgId, g.Title, g.OwnerId,
	)
	return err
}

// Returns the package id for the specified grain
func (tx Tx) GetGrainPackageId(grainId string) (string, error) {
	row := tx.sqlTx.QueryRow("SELECT packageId FROM grains WHERE id = ?", grainId)
	var result string
	err := row.Scan(&result)
	return result, err
}

type GrainInfo struct {
	Id    string
	Title string
}

func (tx Tx) GetCredentialGrains(typ, scopedId string) ([]GrainInfo, error) {
	rows, err := tx.sqlTx.Query(
		`SELECT
			grains.id, grains.title
		FROM
			grains, credentials
		WHERE
			credentials.type = ?
			AND credentials.scopedId = ?
			AND credentials.accountId = grains.ownerId`,
		typ, scopedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []GrainInfo
	for rows.Next() {
		var item GrainInfo
		err := rows.Scan(&item.Id, &item.Title)
		if err != nil {
			return nil, err
		}
		ret = append(ret, item)
	}
	return ret, rows.Err()
}

// encodeCapnp encodes a capnp struct in the format we store in the database,
// i.e. a single packed segment.
func encodeCapnp[T ~capnp.StructKind](v T) ([]byte, error) {
	// We don't strictly need to canonicalize, but it's expedient.
	buf, err := capnp.Canonicalize(capnp.Struct(v))
	if err != nil {
		return nil, err
	}
	return packed.Pack(nil, buf), nil
}

func (tx Tx) SetGrainViewInfo(grainId string, viewInfo grain.UiView_ViewInfo) error {
	buf, err := encodeCapnp(viewInfo)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`UPDATE grains
		SET cachedViewInfo = ?
		WHERE id = ?`,
		buf,
		grainId,
	)
	return err
}
