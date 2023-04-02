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
$Go.import("zenhack.net/go/tempest/capnp/external");

using Util = import "util.capnp";
using Collection = import "collection.capnp";
using Spk = import "package.capnp";

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

  authenticator @2 () -> (authenticator :Authenticator);
  # Return an authenticator that can be used to log in.
}

interface Authenticator {
  # An authenticator provides functionality for authenticating a user with
  # the Tempest server.

  sendEmailAuthToken @0 (address :Text);
  # Send an email authentication token to the specified email address.
  # The token can be passed to ExternalApi.restore to get a LoginSession.
  # Alternatively, making an http request to /login/email/<base64url-encoded token> will
  # return a response that sets a login cookie.
}

interface LoginSession {
  userInfo @0 () -> (info :UserInfo);
  # Return info about the user.

  listGrains @1 (into :Collection.Pusher(Text, Grain));

  listPackages @2 (into :Collection.Pusher(Text, Package));
}

struct Package {
  manifest @0 :Spk.Manifest;
  controller @1 :Controller;

  interface Controller {
    create @0 (title :Text, actionIndex :UInt32) -> (id :Text, grain :Grain);
    # Create a new grain using this package, with the given title
    # and using the action at the specified index in manifest.Actions
    # to spawn. Right now the action must have input == none.
  }
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

struct Grain {
  title @1 :Text;
  # The title of the grain

  sessionToken @2 :Text;
  # A session token which can be exchanged for a cookie on a ui-* subdomain. To
  # open a UI session in the browser, navigate to:
  #
  #   http(s)://ui-${random}.sandstorm.example.net/_sandstorm-init?sandstorm-sid=${sessionToken}&path=${path}
  #
  # This will set an authentication cookie using the Set-Cookie header, and
  # return a redirect to /${path}
  #
  # The token is valid until `handle` is dropped.

  handle @0 :Util.Handle;
  # When handle is dropped, sessionToken is invalidated.
}
