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

// An account identifies a user account. Right now this is just a phantom
// type for use with ID.
type Account struct {
}

// AcccountID is an alias for ID[Account]
type AccountID = ID[Account]

// A Credential is something that Tempest can authenticate a user as. Examples
// (not necessarily all implemented) owner of an email address, SSO account.
type Credential struct {
	// Type defines the type of credential this is.
	Type CredentialType

	// ScopedID is an identifier for this credential that is unique for
	// all credentials of the same type. For example, if the credential
	// type is email addresses, then this could be the address itself.
	ScopedID string `capnp:"scopedId"`
}

// A Type of credential
type CredentialType string

const (
	// Special "dev" accounts; useful for testing but not suitable for
	// real-world use.
	DevCredential CredentialType = "dev"

	// Email login.
	EmailCredential CredentialType = "email"
)

type Role string

const (
	RoleVisitor Role = "visitor"
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
)

func (r Role) IsValid() bool {
	return r == RoleVisitor ||
		r == RoleUser ||
		r == RoleAdmin
}

func (r Role) assertValid() {
	if !r.IsValid() {
		panic("invalid role: " + r)
	}
}

func (r Role) Encompasses(other Role) bool {
	r.assertValid()
	other.assertValid()
	if r == other {
		return true
	}
	switch r {
	case RoleVisitor:
		return false
	case RoleAdmin:
		return true
	case RoleUser:
		return other == RoleVisitor
	default:
		panic("impossible")
	}
}
