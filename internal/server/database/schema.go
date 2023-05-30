package database

import (
	"database/sql"

	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
)

// A Package records information about a package in the database.
type Package struct {
	ID       types.ID[Package] // The package id.
	Manifest spk.Manifest      // The manifest as encoded in the spk.
}

// Initializes the database schema if needed, and returns a DB object.
func InitDB(sqlDB *sql.DB) (DB, error) {
	return exn.Try(func(throw func(error)) DB {
		tx, err := sqlDB.Begin()
		throw(err)
		defer tx.Rollback()

		// Some general notes about the schema:
		//
		// - Anywhere we store a capnp value in a column, it is stored as a single segment
		//   (no headers) in packed encoding.
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS packages (
				-- 128-bit prefix of the sha256 hash of the spk file, hex encoded:
				id VARCHAR(32) PRIMARY KEY NOT NULL,

				-- capnp-encoded package manifest
				manifest BLOB NOT NULL,

				-- Is the package ready to use? The process of installing a package
				-- works like:
				--
				-- 1. Add an entry to this table for the package with ready = false
				-- 2. Move the extracted package to the right location
				-- 3. Set ready to true
				ready BOOLEAN NOT NULL
			)`)
		throw(err)
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS accounts (
				id VARCHAR PRIMARY KEY,

				-- Either "visitor", "user", or "admin"
				role VARCHAR NOT NULL,

				profileDisplayName VARCHAR NOT NULL,
				profilePreferredHandle VARCHAR NOT NULL
				-- TODO: picture, identicon, pronouns
			)`)
		throw(err)
		_, err = tx.Exec(
			// TODO: research SSO support libraries for Go.
			`CREATE TABLE IF NOT EXISTS credentials (
				accountId VARCHAR NOT NULL REFERENCES accounts(id),
				-- Whether this credential is sufficient for logging
				-- in to the account:
				login BOOLEAN NOT NULL,
				-- The type of the credential. Currently always "dev".
				type VARCHAR NOT NULL,
				-- The name of the credential, within the type's naming system.
				-- e.g. for an email authentication system this would just be
				-- the email address.
				scopedId VARCHAR NOT NULL,
				PRIMARY KEY (type, scopedId)
			)`)
		throw(err)
		_, err = tx.Exec(
			`CREATE TABLE IF NOT EXISTS grains (
				-- random base64 url-encoded:
				id VARCHAR(22) PRIMARY KEY NOT NULL,
				-- id of the package for this grain:
				packageId VARCHAR(32) NOT NULL REFERENCES packages(id),
				-- Human readable title chosen by the grain owner:
				title VARCHAR NOT NULL,
				ownerId VARCHAR NOT NULL REFERENCES accounts(id),
				-- cached results for .getViewInfo() on the grain's main UiView.
				cachedViewInfo BLOB
			)`)
		throw(err)
		_, err = tx.Exec(
			`-- A SturdyRef is a random token that grants access to a live capability.
			 --
			 -- This table stores information necessary to restore a sturdyRef. We
			 -- do not store the token itself; instead, we store a sha256 hash, so
			 -- that a database leak does not reveal the necessary information to
			 -- actually restore a sturdyref; the original token must be presented,
			 -- and it is not derivable from the contents of the database.
			CREATE TABLE IF NOT EXISTS sturdyRefs (
				-- raw sha256 hash of the token.
				sha256 BLOB PRIMARY KEY NOT NULL,

				-- "ownerType" and "owner" determine who is allowed to restore
				-- the sturdyRef, and from where. The meaning of "owner" depends
				-- on the value of "ownerType":
				-- * 'grain': "owner" is a grain ID, and the sturdyRef must be
				--   restored via SandstormApi.restore(), from the grain with
				--   the specified ID. grainId FOO
				-- * 'userkeyring': "owner" is in accounts.id: not restorable
				--   directly; logically each user has a "keyring" of capabilities
				--   reachable via APIs that require them to be logged in; some of
				--   those APIs look for this type.
				-- * 'external-api': "owner" is the empty string, and the sturdyRef
				--   must be restored via ExternalApi.restore().
				ownerType VARCHAR NOT NULL,
				owner VARCHAR NOT NULL,

				-- Unix timestamp after which this entry is invalid.
				expires INTEGER,

				-- If not null, this is a reference hosted by the grain with
				-- id 'grainId'. Otherwise, this is provided by the platform
				-- itself.
				grainId VARCHAR(22) REFERENCES grains(id) ON DELETE CASCADE,

				-- capnp struct describing the object this sturdyRef refers to.
				--
				-- If grainId is not null, then the root object of the message
				-- is the ObjectId returned by AppPersistent.save() (see grain.capnp).
				-- If this is null, then this sturdyRef refers to the root UiView
				-- exported by the grain.
				--
				-- If grainId is null, then this the root object is a struct of type
				-- SystemObjectId, from system.capnp.
				objectId BLOB
			)`)
		throw(err)
		_, err = tx.Exec(
			`-- Extended fields for sturdyRefs that refer to uiViews.
			CREATE TABLE IF NOT EXISTS uiViewSturdyRefs (
				-- Hash of the token. The corresponding entry in sturdyRefs
				-- must have a non-null grainId.
				sha256 BLOB PRIMARY KEY NOT NULL REFERENCES sturdyRefs(id) ON DELETE CASCADE,

				-- The permissions defined by the app this sturdyref grants on the grain.
				-- This is a logically a PermissionSet from identity.capnp, encoded as a string
				-- of the characters 't' and 'f' indicating boolean values.
				--
				-- NOTE: if the user is the owner of a grain, then they have all
				-- possible permissions, regardless of the value of this field.
				appPermissions VARCHAR NOT NULL
			)`)
		throw(err)
		throw(tx.Commit())
		return DB{sqlDB: sqlDB}
	})
}
