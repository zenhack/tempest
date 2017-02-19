#!/bin/bash -e

infer_package_name() {
	# Convert the filename $1 to a package name. We munge the name as follows:
	#
	# 1. strip off the capnp file extension and dirname
	# 2. remove dashes
	# 3. convert '+' to 'x'. This is really just for c++.capnp, but it's not
	#    any easier to special case it.
	# 4. special case reserved words, e.g. 'package'
	ret=$(printf '%s' "$(basename $1)" | sed 's/\.capnp$// ; s/-//g ; s/+/x/g')
	case "$ret" in
		package) printf '%s' spk ;;
		*) printf '%s' "$ret" ;;
	esac
}

gen_annotated_schema() {
	# Copy the schema from file "$1" to the current directory, and add
	# appropriate $Go annotations.
	infile="$1"
	outfile="$(basename $infile)"
	cp "$infile" "$outfile"
	package_name="$(infer_package_name $outfile)"
	cat >> "$outfile" << EOF
using Go = import "/go.capnp";
\$Go.package("$package_name");
\$Go.import("zenhack.net/go/sandstorm/capnp/$package_name");
EOF
}

gen_go_src() {
	# Generate go source code from the schema file $1. Create the package
	# directory if necessary.
	file="$1"
	package_name="$(infer_package_name $file)"
	[ -d $package_name ] || mkdir $package_name
	GOPATH=${GOPATH:-$HOME/go}
	capnp compile \
		-I"$GOPATH/src/zombiezen.com/go/capnproto2/std" \
		-ogo:$package_name $file
}

usage() {
	echo "Usage:"
	echo ""
	echo "    $0 import <path/to/schema/files>"
	echo "    $0 patch      # Patch schema"
	echo "    $0 compile    # Generate go source files"
	echo "    $0 clean-gen  # Remove generated files"
	echo "    $0 clean-all  # Remove generated and imported files"
}

# do_* implements the corresponding subcommand described in usage's output.
do_import() {
	input_dir="$1"
	for file in $input_dir/*.capnp; do
		gen_annotated_schema "$file"
	done
	cp $input_dir/sandstorm-http-bridge-request.html ./
}

do_patch() {
	for patchfile in *.patch; do
		patch -p2 < $patchfile
	done
}

do_compile() {
	for file in *.capnp; do
		gen_go_src "$file"
	done
}

do_clean_gen() {
	find "$(dirname $0)" -name '*.capnp.go' -delete
	find "$(dirname $0)" -name '*.capnp.orig' -delete
	find "$(dirname $0)" -type d -empty -delete
}

do_clean_all() {
	do_clean_gen
	find "$(dirname $0)" -name '*.capnp' -delete
}

eq_or_usage() {
	# If "$1" is not equal to "$2", call usage and exit.
	if [ ! $1 = $2 ] ; then
		usage
		exit 1
	fi
}

case "$1" in
	import)    eq_or_usage $# 2; do_import "$2" ;;
	patch)     eq_or_usage $# 1; do_patch ;;
	compile)   eq_or_usage $# 1; do_compile ;;
	clean-gen)  eq_or_usage $# 1; do_clean_gen ;;
	clean-all) eq_or_usage $# 1; do_clean_all ;;
	*) usage; exit 1 ;;
esac
