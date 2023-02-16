package database

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
	"time"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/packed"
	"zenhack.net/go/tempest/capnp/grain"
	"zenhack.net/go/util"
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
	Profile Profile
}

type Profile struct {
	DisplayName     string
	PreferredHandle string
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
	if err != nil {
		return err
	}
	return tx.createOwnerSturdyRef(g.GrainId)
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
	Owner string
}

type UiViewInfo struct {
	Grain       GrainInfo
	Permissions []bool
}

func (tx Tx) GetCredentialUiViews(typ, scopedId string) ([]UiViewInfo, error) {
	accountId, err := tx.GetCredentialAccount(typ, scopedId)
	if err != nil {
		return nil, err
	}
	return tx.AccountUiViews(accountId)
}

func (tx Tx) GetCredentialAccount(typ, scopedId string) (accountId string, err error) {
	row := tx.sqlTx.QueryRow(
		`SELECT accountId FROM credentials WHERE type = ? AND scopedId = ?`,
		typ, scopedId,
	)
	err = row.Scan(&accountId)
	return
}

func (tx Tx) AccountUiViews(accountId string) ([]UiViewInfo, error) {
	rows, err := tx.sqlTx.Query(
		`SELECT
			grains.id,
			grains.title,
			grains.ownerId,
			uiViewSturdyRefs.appPermissions
		FROM
			grains, sturdyRefs, uiViewSturdyRefs
		WHERE
			uiViewSturdyRefs.sha256 = sturdyRefs.sha256
			AND sturdyRefs.grainId = grains.id
			AND sturdyRefs.owner = ?
			AND sturdyRefs.expires > ?
		`,

		"userkeyring-"+accountId,
		time.Now().Unix(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []UiViewInfo
	for rows.Next() {
		var item UiViewInfo
		var perm string
		err := rows.Scan(
			&item.Grain.Id,
			&item.Grain.Title,
			&item.Grain.Owner,
			&perm,
		)
		if err != nil {
			return nil, err
		}
		item.Permissions, err = parsePermissions(perm)
		if err != nil {
			return nil, err
		}
		ret = append(ret, item)
	}
	return ret, rows.Err()
}

func (tx Tx) getGrainOwner(grainId string) (accountId string, err error) {
	err = tx.sqlTx.QueryRow(
		`SELECT ownerId FROM grains WHERE id = ?`,
		grainId,
	).Scan(&accountId)
	return
}

// Generate a random sturdyRef and return the hash of the token (losting the
// original token). Useful for stuff that goes in user keyrings, and thus
// doesn't need to have an actually-usable token.
func genLostToken() [sha256.Size]byte {
	var token [32]byte
	_, err := rand.Read(token[:])
	util.Chkfatal(err)
	return sha256.Sum256(token[:])
}

// createOwnerSturdyRef adds a uiViewSturdyRef for the grain's root
// uiView, belonging to its owner. Should be called once when the grain
// is created.
func (tx Tx) createOwnerSturdyRef(grainId string) error {
	hash := genLostToken()
	ownerId, err := tx.getGrainOwner(grainId)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`INSERT INTO sturdyRefs (
			sha256,
			owner,
			expires,
			grainId
			-- objectId is NULL
		) VALUES (?, ?, ?, ?)`,
		hash[:],
		"userkeyring-"+ownerId,
		math.MaxInt64,
		grainId,
	)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`INSERT INTO uiViewSturdyRefs (sha256, appPermissions)
		-- Don't need to actually supply permissions, because
		-- we own the grain:
		VALUES (?, '')
	`, hash[:])
	return err
}

// parsePermissions parses a string like "ttfftf" from the database into a
// slice of booleans. Returns an error if any characters other than 't'
// and 'f' are in the string.
func parsePermissions(s string) ([]bool, error) {
	ret := make([]bool, len(s))
	for i := range s {
		if s[i] == 't' {
			ret[i] = true
		} else if s[i] == 'f' {
			ret[i] = false
		} else {
			return nil, fmt.Errorf(
				"Error: invalid permissions string in database: %q",
				s,
			)
		}
	}
	return ret, nil
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
