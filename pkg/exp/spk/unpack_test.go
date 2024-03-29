package spk

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"zenhack.net/go/util/exn"
)

const pkgUrlBase = "https://app-index.sandstorm.io/packages/"

type testPackage struct {
	Name  string // Name for the test
	ID    string // Package ID
	AppID string // Return value of AppID.String()
	Valid bool   // Should the package validate?
}

var marketPackages = []testPackage{
	{
		Name:  "Filedrop 1.0.6",
		ID:    "d13f67230d0a4f9a63fc2c0bba87fc52",
		AppID: "nn7axgy3y8kvd0m1mtk3cwca34t916p5d7m4j1j2e874nuz3t8y0",
		Valid: true,
	},
	{
		Name:  "Yet Another TODO 0.1.2",
		ID:    "346bff4bc867afd54320af244dc64e9c",
		AppID: "6ee9j59rp4mepat8pz2pt1sq97xmname4kvhzh4m3uzwrvf1hu10",
		Valid: true,
	},
	{
		Name:  "Davros 0.31.1",
		ID:    "c4c975c3adbeeb77fd928bb90202c049",
		AppID: "8aspz4sfjnp8u89000mh2v1xrdyx97ytn8hq71mdzv4p4d8n0n3h",
		Valid: true,
	},
}

func (p testPackage) FilePath() string {
	return "testdata/" + p.ID + ".spk"
}

func (p testPackage) fetchFromAppMarket() ([]byte, error) {
	return exn.Try(func(throw exn.Thrower) []byte {
		resp, err := http.Get(pkgUrlBase + p.ID)
		throw(err)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			throw(fmt.Errorf(
				"Failed to fetch package %v; http status = %v",
				p.ID, resp.StatusCode,
			))
		}
		buf := &bytes.Buffer{}
		_, err = io.Copy(buf, resp.Body)
		throw(err)
		return buf.Bytes()
	})
}

func (p testPackage) runTest(t *testing.T, downloadIfNeeded bool) {
	buf, err := os.ReadFile(p.FilePath())
	if err != nil {
		if !downloadIfNeeded {
			t.Skip("Package ", p.ID, "not present; skipping")
		}
		t.Log("Fetching package ", p.ID, " from app market")
		buf, err = p.fetchFromAppMarket()
		require.NoError(t, err, "fetching package from app market")
		require.NoError(t, os.WriteFile(p.FilePath(), buf, 0600), "saving package")
	}
	tmpDir := t.TempDir()
	scratchDir := tmpDir + "/scratch-" + p.ID
	require.NoError(t, os.MkdirAll(scratchDir, 0755))
	meta, err := Unpack(scratchDir, bytes.NewBuffer(buf))
	if !p.Valid {
		require.NotNil(t, err, "invalid package fails")
		return
	}
	require.NoError(t, err, "valid package succeeds")
	require.Equal(t, p.AppID, meta.AppID.String(), "app id is as expected")
	expectedPkgHash := sha256.Sum256(buf)
	require.Equal(t, expectedPkgHash[:], meta.Hash[:], "package hash is correct")
	require.Equal(t, p.ID, meta.Hash.ID(), "package id is correct")

	// "corrupt" the package, and make sure it fails:
	buf[rand.Int()%len(buf)]++
	_, err = Unpack(scratchDir, bytes.NewBuffer(buf))
	require.NotNil(t, err, "modifying the package causes failure")
}

func TestUnpackMarketPackages(t *testing.T) {
	t.Parallel()

	downloadIfNeeded := os.Getenv("NETWORK_TEST_OK") == "1"

	for _, p := range marketPackages {
		pkg := p // Avoid sharing a reference across tests.
		t.Run(pkg.Name, func(t *testing.T) {
			t.Parallel()
			pkg.runTest(t, downloadIfNeeded)
		})
	}
}
