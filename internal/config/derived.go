package config

// Constants derived from compile-time config paths:

const (
	TempDir     = Localstatedir + "/tmp/tempest"
	PackagesDir = Localstatedir + "/sandstorm/grains"
)
