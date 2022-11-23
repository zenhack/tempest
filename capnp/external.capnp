@0x9498f3818bafa387;
# This schema describes the APIs available from Sandstorm's externally
# facing websockets (as opposed to those available to grains).
#
# If your sandstorm server is sandstorm.example.net and has HTTPS enabled,
# you can connect to this API at the url:
#
#   wss://sandstorm.example.net/_capnp-api
#
# The bootstrap interface will be of type ExternalApi

using Go = import "/go.capnp";
$Go.package("external");
$Go.import("zenhack.net/go/sandstorm-next/capnp/external");

interface ExternalApi {
  # The bootstrap interface when connecting to the externally facing
  # websocket

  getLoginSession @0 () -> (session :LoginSession);
  # Return a handle to a LoginSession object for the current user. If
  # the user is not logged in, this will fail. logged-in status is
  # determined by the HTTP headers used in the websocket connection
  # request (A session cookie for the web UI. TODO: we may want to
  # support bearer tokens or something for programmatic access?

  restore @1 (sturdyRef :Data) -> (cap :Capability);
  # Restore a sturdyRef as a live capability.
}

interface LoginSession {
  userInfo @0 () -> (info :UserInfo);
  # Return info about the user.
}

struct UserInfo {
  # Information about a user.

  role :union {
    # What level of permissions does the user have on this server?

    visitor @0 :Void;
    # A visitor has no privileges; they can log in and set display name etc,
    # but can't create grains or install apps.

    user @1 :Void;
    # A user can create grains & install apps.

    admin @2 :Void;
    # An admin is an administrator for the server, and has full access to
    # everything.
  }
}
