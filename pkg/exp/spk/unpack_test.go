package spk

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"zenhack.net/go/util/exn"
)

const pkgUrlBase = "https://app-index.sandstorm.io/packages/"

type testPackage struct {
	ID    string // Package ID
	AppID string // Return value of AppID.String()
	Valid bool   // Should the package validate?
}

var marketPackages = []testPackage{
	{
		// Filedrop 1.0.6
		ID:    "d13f67230d0a4f9a63fc2c0bba87fc52",
		AppID: "nn7axgy3y8kvd0m1mtk3cwca34t916p5d7m4j1j2e874nuz3t8y0",
		Valid: true,
	},
	{
		// Yet Another TODO 0.1.2
		ID:    "346bff4bc867afd54320af244dc64e9c",
		AppID: "6ee9j59rp4mepat8pz2pt1sq97xmname4kvhzh4m3uzwrvf1hu10",
		Valid: true,
	},
	{
		// Davros 0.31.1
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
	outputPath := tmpDir + "/output-" + p.ID
	scratchDir := tmpDir + "/scratch-" + p.ID
	require.NoError(t, os.MkdirAll(scratchDir, 755))
	appID, pkgHash, err := UnpackSpk(outputPath, scratchDir, bytes.NewBuffer(buf))
	if !p.Valid {
		require.NotNil(t, err)
		return
	}
	require.NoError(t, err)
	require.Equal(t, p.AppID, appID.String())
	expectedPkgHash := sha256.Sum256(buf)
	require.Equal(t, expectedPkgHash[:], pkgHash[:])
}

func TestMarketPackages(t *testing.T) {
	t.Parallel()

	downloadIfNeeded := os.Getenv("NETWORK_TEST_OK") == "1"

	for _, p := range marketPackages {
		pkg := p // Avoid sharing a reference across tests.
		t.Run("package "+pkg.ID, func(t *testing.T) {
			t.Parallel()
			pkg.runTest(t, downloadIfNeeded)
		})
	}
}
