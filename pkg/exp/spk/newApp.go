package spk

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// Parameters for initializing an application's pkgdef.
type PkgDefParams struct {
	AppID AppID
	Key   Key

	// The capnp schema id for the app's sandstorm-pkgdef.capnp
	SchemaId uint64

	// The keyring in which to save the gerated key.
	// Defaults to ~/.sandstorm/sandstorm-keyring
	KeyringPath string

	// Path to save the package definition. Defaults to
	// `.sandstorm/sandstorm-pkgdef.capnp`.
	PkgDefPath string
}

// Emit a pkgdef with the given parameters, and save the app's key to the keyring.
func (p *PkgDefParams) Emit() error {
	keyring := p.KeyringPath
	if keyring == "" {
		keyring = os.Getenv("HOME") + "/.sandstorm/sandstorm-keyring"
	}
	err := p.Key.AddToFile(keyring)
	if err != nil {
		return fmt.Errorf("error saving key to keyring: %v", err)
	}

	pkgDefPath := p.PkgDefPath
	if pkgDefPath == "" {
		pkgDefPath = ".sandstorm/sandstorm-pkgdef.capnp"
	}
	dir := filepath.Dir(pkgDefPath)
	err = os.MkdirAll(dir, 0644)
	if err != nil {
		return err
	}
	pkgDefFile, err := os.Create(pkgDefPath)
	if err != nil {
		return err
	}
	defer pkgDefFile.Close()
	return pkgdefTemplate.Execute(pkgDefFile, p)
}

// Generate the information needed to initialize a new app.
func NewApp() (*PkgDefParams, error) {
	key, err := GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	appID, err := key.AppID()
	if err != nil {
		return nil, err
	}
	return &PkgDefParams{
		AppID:    appID,
		Key:      key,
		SchemaId: rand.Uint64() | (1 << 63),
	}, nil
}
