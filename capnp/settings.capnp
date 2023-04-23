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

  name @0 :Text;
  # Name of the setting. For admin/system-wide settings,
  # if a variable by this name is set in the environment
  # it overrides other sources.

  type @1 :Schema.Type; # expected type of the setting

  default @2 :Schema.Value;
  # default value of the setting, used if no other sources
  # specify a value.
}

const adminSettings: List(Setting) = [
  # system wide settings, which require admin access to read or modify

  ( # the main URL for the Tempest web interface
    name = "BASE_URL",
    type = (text = void),
    default = (text = "http://local.sandstorm.io:8000"),
  ),
  ( # when sending email, SMTP server to connect to
    name = "SMTP_HOST",
    type = (text = void),
  ),
  ( # port on `smtp.host` to connect to
    name = "SMTP_PORT",
    type = (uint16 = void),
  ),
  ( # email address to send mail from
    name = "SMTP_USERNAME",
    type = (text = void),
  ),
  ( # password to use when authenticating with the SMTP server
    name = "SMTP_PASSWORD",
    type = (text = void),
  ),
];
