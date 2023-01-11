[![Travis CI][travis-img]][travis-ci]
[![Go Report Card][goreport-img]][goreport]

Go Wrappers for Sandstorm's API

The `capnp/`, subtree contains generated packages for the sandstorm API
capnproto schema. Since the schema names are not all legal package
names, the following changes have been made:

* Schema with dashes in their names have had the dashes removed.
* The schema `package` has been mapped to the package `spk` (since
  `package` is a go reserved word).

Note that these schema use the `v3` branch of the go-capnproto2 library,
rather than the stable branch, as its rpc support is more robust.

I try to keep the schema up to date but unfortunately the way the go
capnproto code generator works, that means that new methods on an
interface are breaking changes at the source level, even though
they are compatible at the protocol level. If you upgrade and get
an error about some type not implementing a capnproto `_Server`
interface that it used to, a quick fix is to return an
"unimplemented," error, which will satisfy the type checker and
have the same behavior as before:

```go
func (t *MyType) NewMethod(context.Context, pkg.Interface_newMethod) error {
    return capnp.Unimplemented("TODO: implement")
}
```

...of course, you may alternatively want to implement the new functionality.

The `exp/` subtree contains experimental helper packages; no promises of API
stability are made for these.

Note: we use import path checking; you'll need to import things as e.g:

    import "zenhack.net/go/sandstorm/capnp/grain"

...rather than directly via the URL for this repository.

# Licensing

Apache 2.0, both my stuff and the bits pulled from upstream:

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

A copy of the license is in the file `license.txt`. Individual source
files contain relevant copyright notices.

[travis-ci]: https://travis-ci.org/zenhack/go.sandstorm
[travis-img]: https://travis-ci.org/zenhack/go.sandstorm.svg?branch=master
[goreport-img]: https://goreportcard.com/badge/github.com/zenhack/go.sandstorm
[goreport]: https://goreportcard.com/report/github.com/zenhack/go.sandstorm
