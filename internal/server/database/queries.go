package database

// This file contains wrappers for SQL queries

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
	"time"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/packed"
	"zenhack.net/go/tempest/capnp/grain"
	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util"
)

// AddPackage adds a package to the database. The caller must separately ensure
// that the contents of the package have been extracted to
// <localstatedir>/sandstorm/apps/<pkg.ID>
func (tx Tx) AddPackage(pkg Package) error {
	manifestBlob, err := encodeCapnp(pkg.Manifest)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`INSERT INTO
			packages(id, manifest)
			VALUES (?, ?)
		`,
		pkg.ID,
		manifestBlob,
	)
	return err
}

// GetCredentialPackages returns a list of all packages installed for the user
// associated with the credential.
func (tx Tx) GetCredentialPackages(cred types.Credential) ([]Package, error) {
	// Note: we don't yet handle app installation, so we behave as if all
	// packages are installed for all users. When that changes, we will
	// have to actually filter by account.
	rows, err := tx.sqlTx.Query("SELECT id, manifest FROM packages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []Package
	for rows.Next() {
		var (
			pkg           Package
			manifestBytes []byte
		)
		err = rows.Scan(&pkg.ID, &manifestBytes)
		if err != nil {
			return nil, err
		}
		pkg.Manifest, err = decodeCapnp[spk.Manifest](manifestBytes)
		if err != nil {
			return nil, err
		}
		ret = append(ret, pkg)
	}
	return ret, nil
}

type NewGrain struct {
	GrainID types.GrainID
	PkgID   string
	OwnerID string
	Title   string
}

type NewAccount struct {
	ID      string
	IsAdmin bool
	Profile Profile
}

type Profile struct {
	DisplayName     string
	PreferredHandle string
}

type NewCredential struct {
	AccountID  string
	Login      bool
	Credential types.Credential
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
		a.ID,
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
		c.AccountID,
		c.Login,
		c.Credential.Type,
		c.Credential.ScopedID,
	)
	return err
}

func (tx Tx) AddGrain(g NewGrain) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO grains(id, packageId, title, ownerId) VALUES (?, ?, ?, ?)`,
		g.GrainID, g.PkgID, g.Title, g.OwnerID,
	)
	if err != nil {
		return err
	}
	return tx.createOwnerSturdyRef(g.GrainID)
}

// Returns the package id for the specified grain
func (tx Tx) GetGrainPackageID(grainID types.GrainID) (string, error) {
	row := tx.sqlTx.QueryRow("SELECT packageId FROM grains WHERE id = ?", grainID)
	var result string
	err := row.Scan(&result)
	return result, err
}

type GrainInfo struct {
	ID    types.GrainID
	Title string
	Owner string
}

type UiViewInfo struct {
	Grain       GrainInfo
	Permissions []bool
}

func (tx Tx) GetCredentialUiViews(cred types.Credential) ([]UiViewInfo, error) {
	accountID, err := tx.GetCredentialAccount(cred)
	if err != nil {
		return nil, err
	}
	return tx.AccountUiViews(accountID)
}

func (tx Tx) GetCredentialAccount(cred types.Credential) (accountID string, err error) {
	row := tx.sqlTx.QueryRow(
		`SELECT accountId FROM credentials WHERE type = ? AND scopedId = ?`,
		cred.Type, cred.ScopedID,
	)
	err = row.Scan(&accountID)
	return
}

func (tx Tx) AccountUiViews(accountID string) ([]UiViewInfo, error) {
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

		"userkeyring-"+accountID,
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
			&item.Grain.ID,
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

func (tx Tx) getGrainOwner(grainID types.GrainID) (accountID string, err error) {
	err = tx.sqlTx.QueryRow(
		`SELECT ownerId FROM grains WHERE id = ?`,
		grainID,
	).Scan(&accountID)
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
func (tx Tx) createOwnerSturdyRef(grainID types.GrainID) error {
	hash := genLostToken()
	ownerID, err := tx.getGrainOwner(grainID)
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
		"userkeyring-"+ownerID,
		math.MaxInt64,
		grainID,
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

// Inverse of encodeCapnp
func decodeCapnp[T ~capnp.StructKind](buf []byte) (T, error) {
	buf, err := packed.Unpack(nil, buf)
	if err != nil {
		return T{}, err
	}
	msg := &capnp.Message{Arena: capnp.SingleSegment(buf)}
	ptr, err := msg.Root()
	return T(ptr.Struct()), err
}

func (tx Tx) SetGrainViewInfo(grainID string, viewInfo grain.UiView_ViewInfo) error {
	buf, err := encodeCapnp(viewInfo)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`UPDATE grains
		SET cachedViewInfo = ?
		WHERE id = ?`,
		buf,
		grainID,
	)
	return err
}
