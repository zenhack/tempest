// Package types defines various data types related to Tempest's business logic
// that are used in multiple subsystems.
package types

// An ID is a string identifier. The type parameter is a phantom type, to
// prevent mixing up identifiers for different types of objects.
type ID[T any] string

// Phantom type for use with ID.
type Grain struct{}

// GrainID is an alias for ID[Grain].
type GrainID = ID[Grain]

// A Credential is something that Tempest can authenticate a user as. Examples
// (not necessarily all implemented) owner of an email address, SSO account.
type Credential struct {
	// Type defines the type of credential this is.
	Type CredentialType

	// ScopedId is an identifier for this credential that is unique for
	// all credentials of the same type. For example, if the credential
	// type is email addresses, then this could be the address itself.
	ScopedId string
}

// A Type of credential
type CredentialType string

const (
	// Special "dev" accounts; useful for testing but not suitable for
	// real-world use.
	DevCredential CredentialType = "dev"
)
