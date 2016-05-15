#!/usr/bin/env python
# Copyright (c) 2016 Ian Denhardt <ian@zenhack.net>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy of
# this software and associated documentation files (the "Software"), to deal in
# the Software without restriction, including without limitation the rights to
# use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
# the Software, and to permit persons to whom the Software is furnished to do so,
# subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
# FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
# COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
# IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
# CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
from glob import glob
from os.path import basename, join
from subprocess import check_call
import os
import sys


def get_package_name(src_file):
    package = basename(src_file).replace('.capnp', '')
    if package == 'package':
        package = 'spk'
    package = package\
        .replace('-', '')\
        .replace('+', 'x')
    return package


def go(capnp_file, package):
    os.mkdir(package)
    check_call([
        'capnp', 'compile',
        '-I', join(os.getenv('GOPATH'), 'src/zombiezen.com/go/capnproto2'),
        '-I', '.',
        '-ogo:' + package,
        basename(capnp_file),
    ])


def capnp_1(src_file, dst_file, package):
    with open(src_file) as f:
        src = f.read()
    with open(dst_file, 'w') as f:
        f.write(src)
        f.write('using Go = import "/go.capnp";\n')
        f.write('$Go.package("%s");\n' % package)
        f.write('$Go.import("zenhack.net/go/sandstorm/capnp/%s");\n' % package)


def capnp_many(srcdir):
    src_files = glob(join(srcdir, '*.capnp'))

    for src_file in src_files:
        dst_file = basename(src_file)
        package = get_package_name(src_file)
        capnp_1(src_file, dst_file, package)

    for src_file in src_files:
        dst_file = basename(src_file)
        package = get_package_name(src_file)
        go(dst_file, package)


def main():
    if len(sys.argv) != 2:
        sys.exit("usage: ./gen.py <sandstorm-include>")
    capnp_many(sys.argv[1])


if __name__ == '__main__':
    main()
