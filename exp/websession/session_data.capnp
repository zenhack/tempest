@0x9cd7bc7194780111;

using Go = import "/go.capnp";
$Go.package("websession");
$Go.import("zenhack.net/go/sandstorm/exp/websession");

using Grain = import "/grain.capnp";
using Identity = import "/identity.capnp";
using Powerbox = import "/powerbox.capnp";
using Util = import "/util.capnp";

struct SessionData {
  # SessionData contains all of the information that is passed to Grain.UiView's
  # new*Session methods, but which does not map to part of an HTTP request. We
  # package up the data in one struct, and make it available via the
  # GetSessionData Go function defined in this package.

  context @0 :Grain.SessionContext;
  userInfo @1 :Identity.UserInfo;
  sessionParams @2 :AnyPointer;
  tabId @3 :Data;

  union {
    normal @4 :Void; # A call to newSession

    request :group {
      # A call to newRequestSession
      requestInfo @5 :List(Powerbox.PowerboxDescriptor);
    }

    offer :group {
      # A call to newOfferSession
      offer @6 :Capability;
      descriptor @7 :Powerbox.PowerboxDescriptor;
    }
  }

  # relevant for both `request` and `offer`, but not `normal`:
  sessionType @8 :UInt64;
}
