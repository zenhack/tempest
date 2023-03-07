# Runs the test suite, adding any arguments as extra flags to `go test`.
# Takes care of passing paths.
#
# Annoyingly, we can't just do `go test ./...`, because that will pull in
# frontend code that doesn't build if GOOS != js.

cd "$(dirname $0)/.."
go test $@ \
	./internal/server/... \
	./pkg/...
