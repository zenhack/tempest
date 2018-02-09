#!/usr/bin/env sh

GOPATH=${GOPATH:-${HOME}/go}

capnpc -ogo \
	session_data.capnp \
	-I ${GOPATH}/src/zombiezen.com/go/capnproto2/std/ \
	-I ${GOPATH}/src/zenhack.net/go/sandstorm/capnp/
