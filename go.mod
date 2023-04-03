module zenhack.net/go/tempest

go 1.19

// NOTE: this *must not* be updated to v1.11; that version drops support for
// version 2 of the wire protocol, which is the most recent supported by the
// (ancient) version of mongo bundled with sandstorm.
require go.mongodb.org/mongo-driver v1.10.0

require (
	capnproto.org/go/capnp/v3 v3.0.0-alpha.26
	github.com/gobwas/ws v1.1.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/stretchr/testify v1.8.2
	github.com/tj/assert v0.0.3
	github.com/ulikunitz/xz v0.5.10
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
	golang.org/x/sys v0.1.0
	zenhack.net/go/util v0.0.0-20230402020905-c5270b07d5f2
	zenhack.net/go/vdom v0.0.0-20230403002732-075ae8862f69
	zenhack.net/go/websocket-capnp v0.0.0-20230212023810-f179b8b2c72b
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
