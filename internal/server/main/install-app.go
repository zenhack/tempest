package servermain

import (
	"context"
	"io"
	"os"
	"path/filepath"

	capnpServer "capnproto.org/go/capnp/v3/server"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/util"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/pkg/exp/spk"
	"zenhack.net/go/tempest/pkg/exp/util/bytestream"
	"zenhack.net/go/util/exn"
)

type installStream struct {
	util.ByteStream_Server
	cancel      context.CancelFunc
	userSession userSessionImpl
	ready       chan struct{}
	pkg         external.Package
}

func (s *installStream) Shutdown() {
	s.cancel()
	if shutdowner, ok := s.ByteStream_Server.(capnpServer.Shutdowner); ok {
		shutdowner.Shutdown()
	}
}

func newInstallStream(userSession userSessionImpl) *installStream {
	r, w := bytestream.PipeServer()
	ctx, cancel := context.WithCancel(context.Background())
	s := &installStream{
		ByteStream_Server: w,
		cancel:            cancel,
		userSession:       userSession,
		ready:             make(chan struct{}),
	}
	go s.install(ctx, r)
	return s
}

func (s *installStream) install(ctx context.Context, r *io.PipeReader) {
	err := exn.Try0(func(throw exn.Thrower) {
		db := s.userSession.login.server.db
		meta, err := spk.Unpack(config.TempDir, r)
		throw(err)
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()
		dbPkg := database.Package{
			ID:       types.ID[database.Package](meta.Hash.ID()),
			Manifest: meta.Manifest,
		}
		throw(tx.AddPackage(dbPkg))
		throw(tx.Commit())
		throw(os.Rename(meta.Dir, filepath.Join(config.PackagesDir, string(dbPkg.ID))))
		tx, err = db.Begin()
		throw(err)
		defer tx.Rollback()
		throw(tx.ReadyPackage(dbPkg.ID))
		throw(tx.Commit())

		pkg, err := external.NewPackage(meta.Manifest.Segment())
		throw(err)
		throw(pkg.SetManifest(meta.Manifest))

		pkg.SetController(external.Package_Controller_ServerToClient(pkgController{
			loginSessionImpl: s.userSession.login,
			pkg:              dbPkg,
		}))
		s.pkg = pkg
		close(s.ready)
	})
	if err != nil {
		r.CloseWithError(err)
		// TODO: delete temporary files & package directory.
		return
	}
}

func (s *installStream) GetPackage(ctx context.Context, p external.Package_InstallStream_getPackage) error {
	p.Go()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.ready:
		return exn.Try0(func(throw exn.Thrower) {
			results, err := p.AllocResults()
			throw(err)
			throw(results.SetPackage(s.pkg))
		})
	}
}
