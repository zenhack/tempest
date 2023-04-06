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
	db database.DB
	util.ByteStream_Server
	cancel context.CancelFunc
	ready  chan struct{}
	pkg    external.Package
}

func (s *installStream) Shutdown() {
	s.cancel()
	if shutdowner, ok := s.ByteStream_Server.(capnpServer.Shutdowner); ok {
		shutdowner.Shutdown()
	}
}

func newInstallStream(db database.DB) *installStream {
	r, w := bytestream.PipeServer()
	ctx, cancel := context.WithCancel(context.Background())
	s := &installStream{
		db:                db,
		ByteStream_Server: w,
		cancel:            cancel,
		ready:             make(chan struct{}),
	}
	go s.install(ctx, r)
	return s
}

func (s *installStream) install(ctx context.Context, r *io.PipeReader) {
	err := exn.Try0(func(throw exn.Thrower) {
		meta, err := spk.Unpack(config.TempDir, r)
		throw(err)
		tx, err := s.db.Begin()
		throw(err)
		defer tx.Rollback()
		pkg := database.Package{
			ID:       types.ID[database.Package](meta.Hash.ID()),
			Manifest: meta.Manifest,
		}
		throw(tx.AddPackage(pkg))
		throw(tx.Commit())
		throw(os.Rename(meta.Dir, filepath.Join(config.PackagesDir, string(pkg.ID))))
		tx, err = s.db.Begin()
		throw(err)
		defer tx.Rollback()
		throw(tx.ReadyPackage(pkg.ID))
		throw(tx.Commit())
	})
	if err != nil {
		r.CloseWithError(err)
		// TODO: delete temporary files & package directory.
	}
}
