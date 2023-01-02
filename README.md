This repository contains an experimental rewrite of Sandstorm.

Currently, most of the sandbox setup code is built, and sandstorm-next
is capable of spawning sandstorm apps and plumbing http traffic to them
from the outside, though many http features are not yet implemented.

# Building

To build sandstorm-next, you will need:

- Go 1.19 or later
- [tinygo](https://tinygo.org/)
  - If the build complains about missing `wasm-opt`, you may also need
    to install the `binaryen` package.
- Standard C development tools (make, a C compiler, etc).
- capnp (command line tool) version 0.8 or later.
- capnpc-go code generator plugin

You will also need to separately check out the source for go-capnp and
go.sandstorm:

```
mkdir ../deps
cd ../deps
git clone https://github.com/capnproto/go-capnproto2
git clone https://github.com/zenhack/go.sandstorm
cd -
```

Then, run the configure script and then `make`. The configure script
accepts
most of the same options as typical gnu packages. Additionally you will
need to supply the paths to the repositories checked out above:

```
./configure \
    --with-go-capnp=../deps/go-capnproto2 \
    --with-go-sandstorm=../deps/go.sandstorm
make
```

Then run `make install` to install sandstorm-next system wide.

Tip: you can configure sandstorm-next to share a grain/app storage
directory with a legacy sandstorm system by passing
`--localstatedir=/opt/sandstorm/var` to `./configure`.  In addition to
the files used by legacy sandstorm, `sandstorm-next` will create a
couple extra things underneath that path, namely:

- an extra directory at `sandstorm/mnt`
- a sqlite3 database at `sandstorm/sandstorm.sqlite3`

# Importing data from legacy sandstorm

Sandstorm-next comes with a tool to import some data from a legacy
sandstorm installation's database; after running `make`, there will be
an executable at `_build/sandstorm-legacy-tool`. On a typical sandstorm
server you can export the contents of the database via:

```
./_build/sandstorm-legacy-tool --snapshot-dir /desired/path/to/snapshot export
```

If your sandstorm installation is in a non-standard path or mongoDB is
listening on a different port, you may have to supply additional
options; see `sandstorm-legacy-tool --help` to see the full list.

You can then import the snapshot into sandstorm-next via:

```
./_build/sandstorm-legacy-tool --snapshot-dir /path/to/snapshot import
```

For some development, it can be useful to export & import from legacy
sandstorm frequently. Therefore, we have a Makefile target for this:

```
sudo make export-import
```

...which will automate the above, using the default values to
sandstorm-legacy-tool's flags. It will also destroy the old database
and fix permissions on the new one.

# Running

At present, sandstorm-next has no user interface, and no way to install
apps or create grains; to experiment with it you must separately arrange
a suitable grain storage directory and populate the database. The easiest
way to do this is to point it at an existing sandstorm installation, and
import database info using `sandstsorm-legacy-tool`, per above.

`sandstorm-next` should be run as the user and group chosen by the via
the `--user` and --group` flags to `./configure` (by default both
`sandstorm`).  The easiest way to do this is to run as root:

```
sudo -u sandstorm -g sandstorm /path/to/sandstorm-next
```

For development purposes, the Makefile includes a `dev` target that will
rebuild, reinstall, and then spawn sandstorm-next; simply run:

```
sudo make dev`
```

`sandstorm-next` will start a web server on port 8000; to connect to the
UI, go to `http://local.sandstorm.io:8000`.

To log in with a developer account, click "Log in with dev account", and
enter the dev account's name, e.g. "Alice Dev Admin."

Once you have logged in, a list of grains the user owns will be
displayed; click the links to open the grains.

This will display the grain's UI within an iframe. Things like
offer iframes and anything that uses sandstorm specific APIs will not
work currently.
