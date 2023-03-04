@0xd7e3ca8e87d116b7;
# The grain-agent (cmd/tempest-grain-agent) is the first process launched
# when a sandbox is created. This schema defines (private) data structures
# for interacting with it.

using Go = import "/go.capnp";
$Go.package("grainagent");
$Go.import("zenhack.net/go/tempest/internal/capnp/grain-agent");

struct LaunchCommand {
  # A LaunchCommand is passed as the first argument to the grain agent, as
  # a single (standard-)base64 encoded segment. This tells the grain agent what to
  # do on startup.

  union {
    continueGrain @0 :Void;
    # Invoke the package's Mainfest.continueCommand to start up the
    # (pre-existing) grain.

    initGrain @1 :UInt32;
    # Invoke the ith Action in the packages Manifest.actions, to start
    # up a new grain for the first time.
  }
}
