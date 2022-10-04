@0xd03ac8d96fd2fb71;
# APIs for managing running containers.

using Go = import "/go.capnp";

$Go.package("container");
$Go.import("zenhack.net/go/sandstorm-next/capnp/container");

using Util = import "/util.capnp";

interface Spawner {
  # A spawner can be used to spawn containers.

  spawn @0 (bootstrap :Capability, packageId :Text, grainId :Text)
    -> (bootstrap :Capability, handle :Util.Handle);
  # Spawn a new container, executing with the filesystem for the given package,
  # and the given grain's mutable storage, supplying the given capability as the
  # bootstrap interface exposed to the grain. Returns the bootstrap interface
  # exposed by the container, and a handle. If the handle is dropped, the
  # container will be killed.
}
