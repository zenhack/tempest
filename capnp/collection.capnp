@0xf85f95aca41e787e;
# This file defines schema for working with collections. A collection is
# logically a map with a key & value type.

using Go = import "/go.capnp";
$Go.package("collection");
$Go.import("zenhack.net/go/tempest/capnp/collection");

using Util = import "util.capnp";

interface Pusher(K, V) {
  # A pusher is a write-only refers to a collection with keys of type K and
  # values of type V. Callers make a series of calls to modify the colleciton,
  # followed a call to ready to mark the end of a batch of updates. The
  # collection on the other side does not need to wait until the end of a
  # batch to start acting on the changes, but it may be useful to know, esp.
  # for initially populating the collection.

  upsert @0 (key :K, value :V);
  # Insert a new key, value pair into the collection, or update the value if
  # the key is already present.

  remove @1 (key :K);
  # remove a key from the collection.

  clear @3 ();
  # Remove all keys from the collection.

  ready @2 ();
  # Report that the caller is done sending updates for the moment.
}

interface Puller(K, V) {
  # A puller is a read-only reference to a collection.

  sync @0 (into :Pusher(K, V)) -> (subscription :Util.Handle);
  # Synchronize the Puller with `into` by pushing updates into it.
  # Calls .ready() after the initial synchronization is complete.
  # Afterwards, any changes to the collection will be sent in subsequent
  # batches of operations.
  #
  # When the the returned `subsription` handle is dropped, the Puller
  # will drop its reference to `into` and cease sending updates.

  key @1 (name :K) -> (getter :Util.Getter(V));
  # Return a getter for a single key.
}
