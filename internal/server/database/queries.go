package database

// This file contains wrappers for SQL queries

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"math"
	"time"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/exc"
	"capnproto.org/go/capnp/v3/packed"
	"zenhack.net/go/tempest/capnp/grain"
	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/tokenutil"
)

// AddPackage adds a package to the database, initally marked as not ready.
// The caller must then move the extracted package to the right location,
// and then call ReadyPackage to complete installation.
func (tx Tx) AddPackage(pkg Package) error {
	manifestBlob, err := encodeCapnp(pkg.Manifest)
	if err != nil {
		return err
	}
	_, err = tx.sqlTx.Exec(
		`INSERT INTO
			packages(id, manifest, ready)
			VALUES (?, ?, ?)
		`,
		pkg.ID,
		manifestBlob,
		false,
	)
	return exc.WrapError("AddPackage", err)
}

func (tx Tx) ReadyPackage(id types.ID[Package]) error {
	_, err := tx.sqlTx.Exec(`UPDATE packages SET ready = true WHERE id = ?`, id)
	return exc.WrapError("ReadyPackage", err)
}

// GetCredentialPackages returns a list of all packages installed for the user
// associated with the credential.
func (tx Tx) GetCredentialPackages(cred types.Credential) ([]Package, error) {
	// Note: we don't yet handle app installation, so we behave as if all
	// packages are installed for all users. When that changes, we will
	// have to actually filter by account.
	rows, err := tx.sqlTx.Query("SELECT id, manifest FROM packages")
	if err != nil {
		return nil, exc.WrapError("GetCredentialPackages", err)
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
	PkgID   types.ID[Package]
	OwnerID string
	Title   string
}

type NewAccount struct {
	ID      string
	Role    types.Role
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
			, role
			, profileDisplayName
			, profilePreferredHandle
			)
			VALUES (?, ?, ?, ?)`,
		a.ID,
		a.Role,
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
	return result, exc.WrapError("GetGrainPackageID", err)
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

func (tx Tx) AccountGrainPermissions(accountID string, grainID types.GrainID) (permissions []bool, err error) {
	row := tx.sqlTx.QueryRow(
		`SELECT
			uiViewSturdyRefs.appPermissions
		FROM
			sturdyRefs, uiViewSturdyRefs
		WHERE
			uiViewSturdyRefs.sha256 = sturdyRefs.sha256
			AND sturdyRefs.grainId = ?
			AND sturdyRefs.objectId = null
			AND sturdyRefs.ownerType = 'userkeyring'
			AND sturdyRefs.owner = ?
			AND sturdyRefs.expires > ?
		`,
		grainID,
		accountID,
		time.Now().Unix(),
	)
	var perm string
	err = row.Scan(&perm)
	if err != nil {
		return nil, err
	}
	return parsePermissions(perm)
}

func (tx Tx) NewSharingToken(
	grainID types.GrainID,
	ownerID string,
	perms []bool,
	note string,
) (string, error) {
	permString := fmtPermissions(perms)
	token := base64.RawURLEncoding.EncodeToString(tokenutil.GenToken())
	hash, err := tx.SaveSturdyRef(
		SturdyRefKey{
			Token:     []byte(token),
			OwnerType: "userkeyring",
			Owner:     ownerID,
		},
		SturdyRefValue{
			Expires: time.Unix(math.MaxInt64, 0), // never
			GrainID: grainID,
		},
	)
	if err != nil {
		return "", err
	}
	_, err = tx.sqlTx.Exec(`
		INSERT INTO
			uiViewSturdyRefs (sha256, appPermissions)
		VALUES (?, ?)`,
		hash, permString,
	)
	if err != nil {
		return "", err
	}
	return token, err
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
	err = exc.WrapError("GetCredentialAccount", row.Scan(&accountID))
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
			AND sturdyRefs.ownerType = 'userkeyring'
			AND sturdyRefs.owner = ?
			AND sturdyRefs.expires > ?
		`,

		accountID,
		time.Now().Unix(),
	)
	if err != nil {
		return nil, exc.WrapError("AccountUIViews", err)
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
	err = exc.WrapError("getGrainOwner", err)
	return
}

// createOwnerSturdyRef adds a uiViewSturdyRef for the grain's root
// uiView, belonging to its owner. Should be called once when the grain
// is created.
func (tx Tx) createOwnerSturdyRef(grainID types.GrainID) error {
	ownerID, err := tx.getGrainOwner(grainID)
	if err != nil {
		return err
	}
	hash, err := tx.SaveSturdyRef(
		SturdyRefKey{
			Token:     tokenutil.GenToken(),
			OwnerType: "userkeyring",
			Owner:     ownerID,
		},
		SturdyRefValue{
			Expires: time.Unix(math.MaxInt64, 0), // never
			GrainID: grainID,
		},
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
				"error: invalid permissions string in database: %q",
				s,
			)
		}
	}
	return ret, nil
}

// Inverse of parse permissions
func fmtPermissions(perm []bool) string {
	buf := make([]byte, len(perm))
	for i := range perm {
		if perm[i] {
			buf[i] = 't'
		} else {
			buf[i] = 'f'
		}
	}
	return string(buf)
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

// A SturdyRefKey is the data by which a sturdyRef may be fetched from the database (using
// RestoreSturdyRef).
type SturdyRefKey struct {
	Token            []byte
	OwnerType, Owner string
}

// A SturdyRefValue is a persistent value stored in the database, which may be fetched
// via RestoreSturdyRef.
type SturdyRefValue struct {
	Expires  time.Time
	GrainID  types.GrainID
	ObjectID capnp.Struct
}

// Save a SturdyRef in the database. k's token must not be nil. Returns the sha256
// hash of the token, which serves as a key in the database.
func (tx Tx) SaveSturdyRef(k SturdyRefKey, v SturdyRefValue) ([sha256.Size]byte, error) {
	if k.Token == nil {
		panic("Called SaveSturdyRef with nil token")
	}
	hash := sha256.Sum256(k.Token)
	var grainID *types.GrainID
	if v.GrainID != "" {
		grainID = &v.GrainID
	}
	var (
		objectID []byte
		err      error
	)
	if v.ObjectID.IsValid() {
		objectID, err = encodeCapnp(v.ObjectID)
	}
	if err != nil {
		return hash, err
	}
	_, err = tx.sqlTx.Exec(
		`INSERT INTO sturdyRefs
			( sha256
			, ownerType
			, owner
			, expires
			, grainId
			, objectId
			)
			VALUES (?, ?, ?, ?, ?, ?)
		`,
		hash[:],
		k.OwnerType,
		k.Owner,
		v.Expires.Unix(),
		grainID,
		objectID,
	)
	return hash, err
}

// Restore a SturdyRef from the database.
func (tx Tx) RestoreSturdyRef(k SturdyRefKey) (SturdyRefValue, error) {
	hash := sha256.Sum256(k.Token)
	row := tx.sqlTx.QueryRow(
		`SELECT expires, grainId, objectId
		FROM sturdyRefs
		WHERE
			ownerType = ?
			AND owner = ?
			AND sha256 = ?
			AND expires > ?
		`,
		k.OwnerType,
		k.Owner,
		hash[:],
		time.Now().Unix(),
	)
	var (
		expires  int64
		objectID []byte
		grainID  *types.GrainID

		ret SturdyRefValue
	)
	err := row.Scan(&expires, &grainID, &objectID)
	err = exc.WrapError("RestoreSturdyRef", err)
	if err != nil {
		return ret, err
	}
	ret.Expires = time.Unix(expires, 0)
	if len(objectID) > 0 {
		ret.ObjectID, err = decodeCapnp[capnp.Struct](objectID)
	}
	if grainID != nil {
		ret.GrainID = *grainID
	}
	return ret, err
}

// Delete a sturdyref from the database.
func (tx Tx) DeleteSturdyRef(k SturdyRefKey) error {
	hash := sha256.Sum256(k.Token)
	_, err := tx.sqlTx.Exec(`DELETE FROM sturdyRefs WHERE sha256 = ?`, hash[:])
	return err
}

// Get the role corresponding to the credential. Returns RoleVisitor for unknown
// credentials.
func (tx Tx) CredentialRole(cred types.Credential) (role types.Role, err error) {
	row := tx.sqlTx.QueryRow(`
		SELECT role
		FROM accounts, credentials
		WHERE
			accounts.id = credentials.accountId
			AND credentials.type = ?
			AND credentials.scopedId = ?`,
		cred.Type,
		cred.ScopedID)
	err = row.Scan(&role)
	if err == sql.ErrNoRows {
		return types.RoleVisitor, nil
	}
	return role, exc.WrapError("CredentialRole", err)
}
