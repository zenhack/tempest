#!/usr/bin/env python
# Copyright (c) 2016 Ian Denhardt <ian@zenhack.net>
#
# All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
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
