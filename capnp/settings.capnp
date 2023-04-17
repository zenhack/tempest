@0xdf25727aa20cb09f;

using Go = import "/go.capnp";
using Schema = import "/capnp/schema.capnp";

$Go.package("settings");
$Go.import("zenhack.net/go/tempest/capnp/settings");

struct Setting {
  name @0 :Text;
  type @1 :Schema.Type;
  default @2 :Schema.Value;
  envVar @3 :Text;
}

const adminSettings: List(Setting) = [
  ( # the main URL for the tempest web interface
    name = "base_url",
    type = (text = void),
    default = (text = "http://local.sandstorm.io:8000"),
    envVar = "BASE_URL",
  ),
  ( # when sending email, SMTP server to connect to
    name = "smtp.host",
    type = (text = void),
    envVar = "SMTP_HOST",
  ),
  ( # port on `smtp.host` to connect to
    name = "smtp.port",
    type = (uint16 = void),
    envVar = "SMTP_PORT",
  ),
  ( # email address to send mail from
    name = "smtp.username",
    type = (text = void),
    envVar = "SMTP_USERNAME",
  ),
  ( # password to use when authenticating with the SMTP server
    name = "smtp.password",
    type = (text = void),
    envVar = "SMTP_PASSWORD",
  ),
];
