find $(dirname $0) -name '*.capnp' -delete
find $(dirname $0) -name '*.capnp.go' -delete
find $(dirname $0) -type d -empty -delete
