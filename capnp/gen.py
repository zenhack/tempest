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
from os.path import basename, dirname, join
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
    destdir = join(dirname(capnp_file), package)
    os.mkdir(destdir)
    os.chdir(dirname(destdir))
    check_call([
        'capnp', 'compile',
        '-I', join(os.getenv('GOPATH'), 'src'),
        '-I', '..',
        '-ogo:' + package,
        basename(capnp_file),
    ])
    os.chdir('..')


def capnp_1(src_file, dst_file, package, import_path):
    with open(src_file) as f:
        src = f.read()
    with open(dst_file, 'w') as f:
        f.write(src)
        f.write('using Go = import "/zombiezen.com/go/capnproto2/go.capnp";\n')
        f.write('$Go.package("%s");\n' % package)
        f.write('$Go.import("zenhack.net/go/sandstorm/capnp/%s");\n' % import_path)


def capnp_many(srcdir, destdir):
    src_files = glob(join(srcdir, '*.capnp'))

    for src_file in src_files:
        dst_file = join(destdir, basename(src_file))
        package = get_package_name(src_file)
        capnp_1(src_file, dst_file, package, join(destdir, package))

    for src_file in src_files:
        dst_file = join(destdir, basename(src_file))
        package = get_package_name(src_file)
        go(dst_file, package)


def main():
    if len(sys.argv) != 3:
        sys.exit("usage: ./gen.py <capnp-include> <sandstorm-include>")
    os.mkdir('capnp')
    os.mkdir('sandstorm')
    capnp_many(sys.argv[1], 'capnp')
    capnp_many(sys.argv[2], 'sandstorm')


if __name__ == '__main__':
    main()
