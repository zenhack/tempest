@0xcf769d5676f4466f;

using Go = import "/go.capnp";
$Go.package("http");
$Go.import("zenhack.net/go/sandstorm-next/capnp/http");

using Util = import "/util.capnp";

using Header = Util.KeyValue;

struct Request {
  method @0 :Text;
  headers @1 :List(Header);
  responder @2 :Responder;
}

interface Responder {
  respond @0 (status :UInt16, headers :List(Header)) -> (sink :Util.ByteStream);
}

interface Server {
  request @0 (request :Request) -> (requestBody :Util.ByteStream);
}
