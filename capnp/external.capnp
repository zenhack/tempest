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
using Grain = import "grain.capnp";
using Identity = import "identity.capnp";

interface ExternalApi {
  # The bootstrap interface when connecting to the externally facing
  # websocket

  getSessions @0 () -> Sessions;
  # Return handles to sessions for the current user. Fields representing
  # roles the user does not have will be left null. If the user is not logged
  # in, this will throw an exception. logged-in status is determined by
  # the HTTP headers used in the websocket connection request (A session
  # cookie for the web UI. TODO: we may want to support bearer tokens
  # or something for programmatic access?

  restore @1 (sturdyRef :Data) -> (cap :Capability);
  # Restore a sturdyRef as a live capability.

  authenticator @2 () -> (authenticator :Authenticator);
  # Return an authenticator that can be used to log in.
}

struct Sessions {
  visitor @0 :VisitorSession;
  user @1 :UserSession;
}

interface Authenticator {
  # An authenticator provides functionality for authenticating a user with
  # the Tempest server.

  sendEmailAuthToken @0 (address :Text);
  # Send an email authentication token to the specified email address.
  # Making an http request to /login/email/<base64url-encoded token> will
  # return a response that sets a login cookie. In the future, it will
  # be possible to also redeem the token via ExternalApi.restore(), returning
  # some appropriate capability.
}

interface VisitorSession {
  # A VisitorSession provides operations that only require the 'visitor' role.

  views @0 () -> (views :Collection.Puller(Text, UiView));
  # Get the ui views into grains that the caller has access to.
  # XXX: this currently keys on grain IDs, but we eventually want to
  # be able to have more than one view per grain. This interface will
  # need to be reworked.
}

struct Package {
  manifest @0 :Spk.Manifest;
  controller @1 :Controller;

  interface Controller {
    create @0 (title :Text, actionIndex :UInt32) -> (id :Text, view :UiView);
    # Create a new grain using this package, with the given title
    # and using the action at the specified index in manifest.Actions
    # to spawn. Right now the action must have input == none.
    # Returns the id of the new grain and it's primary UiView.
  }

  interface InstallStream extends (Util.ByteStream) {
    # An InstallStream is a Util.ByteStream into which an .spk file can
    # be written; when done() is called, the package will be installed
    # and become available to the user.

    getPackage @0 () -> (id :Text, package :Package);
    # getPackage returns the package once it is installed. Generally,
    # users will want to call this immediately, and then write the
    # spk using write. This will not return until after done() is called.
  }
}

interface UserSession {
  # A UserSession provides operations that require the 'user' role.

  installPackage @0 () -> (stream :Package.InstallStream);
  # Install a package. If the stream is dropped before calling done(), installation
  # is cancelled.

  listPackages @1 (into :Collection.Pusher(Text, Package));
  # List the packages that the caller has installed.
}

struct UiView {
  # A UiView includes information and access to a UiView. For now, this maps 1-to-1
  # onto grains, but in the future Tempest will support exporting multiple Ui views
  # from the same grain.

  title @1 :Text;
  # The title of the view

  sessionToken @2 :Text;
  # A session token which can be exchanged for a cookie on a ui-* subdomain. To
  # open a UI session in the browser, navigate to:
  #
  #   http(s)://ui-${subdomain}.sandstorm.example.net/_sandstorm-init?sandstorm-sid=${sessionToken}&path=${path}
  #
  # This will set an authentication cookie using the Set-Cookie header, and
  # return a redirect to /${path}
  #
  # The token is valid until `handle` is dropped.

  subdomain @3 :Text;
  # The the subdomain, sans ui- prefix, that a session for this grain view
  # should be loaded from.

  viewInfo @4 :Grain.UiView.ViewInfo;
  # View info for the UiView.

  controller @0 :Controller;
  # Controller for manipulating the grain. When controller is dropped, sessionToken
  # is invalidated.

  interface Controller {
    makeSharingToken @0 (permissions :Identity.PermissionSet, note :Text) -> (token :Text);
    # Make a sharing token, which can be used to construct sharing URLs. The url should be:
    # http(s)://sandstorm.example.net/#/shared/${token}
    #
    # The permissions argument specifies which permissions should be shared. Each index
    # corresponds to a permission from the permissions field of the ViewInfo.
    #
    # Note is a human-readable note regaring the token; this may be displayed in various
    # places in the UI.
  }
}
