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
	ID    string
	Valid bool
}

var marketPackages = []testPackage{
	{
		// Filedrop 1.0.6
		ID:    "d13f67230d0a4f9a63fc2c0bba87fc52",
		Valid: true,
	},
	{
		// Yet Another TODO 0.1.2
		ID:    "346bff4bc867afd54320af244dc64e9c",
		Valid: true,
	},
	{
		// Davros 0.31.1
		ID:    "c4c975c3adbeeb77fd928bb90202c049",
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
		buf, err = p.fetchFromAppMarket()
		require.NoError(t, err, "fetching package from app market")
		require.NoError(t, os.WriteFile(p.FilePath(), buf, 0600), "saving package")
	}
	tmpDir := t.TempDir()
	outputPath := tmpDir + "/output-" + p.ID
	scratchDir := tmpDir + "/scratch-" + p.ID
	require.NoError(t, os.MkdirAll(scratchDir, 755))
	_, pkgHash, err := UnpackSpk(outputPath, scratchDir, bytes.NewBuffer(buf))
	if !p.Valid {
		require.NotNil(t, err)
		return
	}
	require.NoError(t, err)
	expectedPkgHash := sha256.Sum256(buf)
	require.Equal(t, expectedPkgHash[:], pkgHash[:])
}

func TestMarketPackages(t *testing.T) {
	t.Parallel()

	downloadIfNeeded := os.Getenv("NETWORK_TEST_OK") == "1"

	for _, p := range marketPackages {
		t.Run("package "+p.ID, func(t *testing.T) {
			t.Parallel()
			p.runTest(t, downloadIfNeeded)
		})
	}
}
