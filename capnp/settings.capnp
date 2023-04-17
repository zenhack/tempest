@0xdf25727aa20cb09f;
# This schema deals with settings, both system wide (adminSettings)
# and per user/per grain and such (TODO).
#
# TODO: we will want to support storing settings in the database;
# describe the semantics of that. Currently it is not implemented.

using Go = import "/go.capnp";
using Schema = import "/capnp/schema.capnp";

$Go.package("settings");
$Go.import("zenhack.net/go/tempest/capnp/settings");

struct Setting {
  # Specification of a setting

  name @0 :Text; # name of the setting
  type @1 :Schema.Type; # expected type of the setting

  default @2 :Schema.Value;
  # default value of the setting, used if no other sources
  # specify a value.

  envVar @3 :Text;
  # environment variable from which the setting may be read.
  # only applicable to admin/system wide settings. If this
  # variable is set in the environment it overrides other
  # sources of
}

const adminSettings: List(Setting) = [
  # system wide settings, which require admin access to read or modify

  ( # the main URL for the Tempest web interface
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
