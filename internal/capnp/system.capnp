# TODO(cleanup): this file doesn't feel like a very natural divison to me;
# figure out somewhere conceptually cleaner to put the stuff in defined
# here.

@0xa9980bd0b9075eb0;

using Go = import "/go.capnp";
$Go.package("system");
$Go.import("zenhack.net/go/tempest/internal/capnp/system");

struct SystemObjectId {
  # A SystemObjectId describes a capability that is implemented by the
  # platform; this contains information necessary to restored it to a
  # live capability, but it is not a sturdyRef, because it is not
  # self-authenticating.
  union {
    emailLoginToken @0 :Text;
    # An email login token. For now, this isn't used at the capnp level,
    # but is used to support email login. Perhaps conceptually the right
    # thing is for this to restore to a LoginSession (from external.capnp)?

    unknown @1 :Void;
    # placeholder so this can be a union.
  }
}
