@0xcf769d5676f4466f;

using Go = import "/go.capnp";
$Go.package("http");
$Go.import("zenhack.net/go/tempest/capnp/http");

using Util = import "/util.capnp";

using Header = Util.KeyValue;

struct Request {
  path @2 :Text;
  method @0 :Text;
  headers @1 :List(Header);
}

interface Responder {
  respond @0 (status :UInt16, headers :List(Header)) -> (sink :Util.ByteStream);
}

interface Server {
  request @0 (request :Request, responder :Responder) -> (requestBody :Util.ByteStream);

}
