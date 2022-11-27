@0xbbb10ce386c6624a;
# This file defines the schema for data stored in cookies;
# we store these as a single segment using gorilla's session
# package.

using Go = import "/go.capnp";
$Go.package("cookie");
$Go.import("zenhack.net/go/sandstorm-next/capnp/private/cookie");

struct UserSession {
  # Session cookie for a user's login. Stored at Values["data"]
  # in the "user-session" cookie.

  sessionId @2 :Data;
  # Randomized session id

  credential :group {
    # Details about the credential that was used to log in.
    type @0 :Text;
    scopedId @1 :Text;
  }
}
