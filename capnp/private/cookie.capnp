@0xbbb10ce386c6624a;
# This file defines the schema for data stored in cookies;
# the code for managing these is in the package
# `go/internal/server/session`.

using Go = import "/go.capnp";
$Go.package("cookie");
$Go.import("zenhack.net/go/sandstorm-next/capnp/private/cookie");

struct UserSession {
  # Session cookie for a user's login. Stored in the
  # "sandstorm-user-session" cookie.

  sessionId @2 :Data;
  # Randomized session id

  credential :group {
    # Details about the credential that was used to log in.
    type @0 :Text;
    scopedId @1 :Text;
  }
}

struct GrainSession {
  # Session cookie for a user's session on a ui-* subdomain for
  # a given grain. Stored in the "sandstorm-grain-session"
  # cookie.

  grainId @0 :Text;
  # The grain id

  sessionId @1 :Data;
  # The session id (from UserSession)
}
