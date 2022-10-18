This repository contains an experimental rewrite of Sandstorm.

Currently, most of the sandbox setup code is built, and sandstorm-next
is capable of spawning sandstorm-http-bridge based apps and plumbing
http traffic to them from the outside.

# Building

To build sandstorm-next, you will need:

- Go 1.19 or later
- Standard C development tools (make, a C compiler, etc).

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
`--localstatedir=/opt/sandstorm/var` to `./configure`. `sandstorm-next`
will create an extra directory `sandstorm/mnt` underneath that path, but
this is harmless.

# Running

At present, sandstorm-next has no user interface, and no way to install
apps or create grains; to experiment with it you must separately arrange
a suitable grain storage directory. The easiest way to do this is to
point it at an existing sandstorm installation, per above.

You will need to be a member of the `sandstorm` group to run
`sandstorm-next` (or if you specified a different group name via the
`--group` configure flag, that group instead).

On startup, the `sandstorm-next` executable will attempt to launch the
grain specified in the environment variable `DUMMY_GRAIN_ID`, using the
package `DUMMY_PACKAGE_ID`. You are responsible for making sure these
are present and agree.

`sandstorm-next` will start a web server on port 8000; to connect to the
grain's UI, go to `http://grain.local.sandstorm.io:8000`.

This will display the grain's UI with no surrounding iframe. Things like
offer iframes and anything that uses sandstorm specific APIs will not
work currently.
