#!/usr/bin/env python
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
        .replace('-', '_')\
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
