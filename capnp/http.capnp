@0xcf769d5676f4466f;

using Io = import "./io.capnp";

struct Request {
  method @0 :Text;
  headers @1 :List(Header);
  responder @2 :Responder
}

struct Header {
  key @0 :Text;
  value @1 :Text;
}

interface Responder {
  respond @0 (status :UInt16, headers :ListHeader) -> (sink :Io.ByteSink);
}

interface Server {
  makeRequest @0 (request :Request) -> (requestBody :Io.ByteSink);
}
