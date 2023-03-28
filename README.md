This repository contains an experimental replacement for [Sandstorm][1].
See the [blog post][2].

Currently, most of the sandbox setup code is built, and tempest
is capable of spawning sandstorm apps and plumbing http traffic to them
from the outside, though many http features are not yet implemented.

# Building

To build tempest, you will need:

- Go 1.19 or later
- [tinygo](https://tinygo.org/)
  - If the build complains about missing `wasm-opt`, you may also need
    to install the `binaryen` package.
- Standard C development tools (make, a C compiler, etc).
- The `bpf_asm` command, included in the linux kernel source tree.
- capnp (command line tool) version 0.8 or later.
- capnpc-go code generator plugin

You will also need to separately check out the source for go-capnp:

```
mkdir ../deps
cd ../deps
git clone https://github.com/capnproto/go-capnp
cd -
```

Then, run the configure script and then `make`. The configure script
accepts
most of the same options as typical gnu packages. Additionally you will
need to supply the paths to the repository checked out above:

```
./configure \
    --with-go-capnp=../deps/go-capnp \
    --localstatedir=/opt/sandstorm/var
make
```

Note: this will also configure tempest to share a grain/app storage
directory with a sandstorm system.  In addition to
the files used by sandstorm, `tempest` will create a couple extra things
underneath that path, namely:

- an extra directory at `sandstorm/mnt`
- a sqlite3 database at `sandstorm/sandstorm.sqlite3`

Then run `make install` to install tempest system wide.

# Importing data from sandstorm

Tempest comes with a tool to import some data from a sandstorm
installation's database; after running `make`, there will be
an executable at `_build/sandstorm-import-tool`. On a typical sandstorm
server you can export the contents of the database via:

```
mkdir ../sandstormexport
./_build/sandstorm-import-tool --snapshot-dir ../sandstormexport export
```

If your sandstorm installation is in a non-standard path or mongoDB is
listening on a different port, you may have to supply additional
options; see `sandstorm-import-tool --help` to see the full list.

You can then import the snapshot into tempest via:

```
./_build/sandstorm-import-tool --snapshot-dir ../sandstormexport import
```

For some development, it can be useful to export & import from sandstorm
frequently. Therefore, we have a Makefile target for this:

```
sudo make export-import
```

...which will automate the above, using the default values to
sandstorm-import-tool's flags. It will also destroy the old database
and fix permissions on the new one.

# Running

At present, tempest has no way to install apps or create grains; to
experiment with it you must separately arrange a suitable grain storage
directory and populate the database. The easiest way to do this is to
point it at an existing sandstorm installation, and import database info
using `sandstorm-import-tool`, per above.

`tempest` should be run as the user and group chosen by the via
the `--user` and `--group` flags to `./configure` (by default both
`sandstorm`).  The easiest way to do this is to run as root:

```
sudo -u sandstorm -g sandstorm ./_build/tempest
```

For development purposes, the Makefile includes a `dev` target that will
rebuild, reinstall, and then spawn tempest; simply run:

```
sudo make dev
```

The `tempest` will start tempest. The following environment variables
are influential:

- `BASE_URL`: the main URL for the tempest web interface. Defaults to
  `http://local.sandstorm.io:8000`.
- `SMTP_HOST`: When sending email, SMTP server to connect to.
- `SMTP_PORT`: Port on `SMTP_HOST` to connect to.
- `SMTP_USERNAME`: Email address to send mail from.
- `SMTP_PASSWORD`: Password to use when authenticating with
  the SMTP server.

To log in with a developer account, visit the web interface (as defined
by `BASE_URL`), click "Log in with dev account", and enter the dev
account's name, e.g. "Alice Dev Admin."

Once you have logged in, a list of grains the user owns will be
displayed; click the links to open the grains.

This will display the grain's UI within an iframe. Things like
offer iframes and anything that uses sandstorm specific APIs will not
work currently.

[1]: https://sandstorm.io
[2]: https://zenhack.net/2023/01/06/introducing-tempest.html
