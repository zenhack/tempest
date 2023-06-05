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

  ( # URL for the acme directory to use to obtain TLS certs.
    # TODO: default to using Let's Encrypt.
    name = "ACME_DIRECTORY_URL",
    type = (text = void),
  ),
  ( # DNS provider to use for acme; see https://go-acme.github.io/lego/dns/
    # for a list of valid providers.
    name = "ACME_DNS_PROIVDER",
    type = (text = void),
  ),
  ( # Email address to use for acme protocol.
    name = "ACME_EMAIL",
    type = (text = void),
  ),

  ( # the main URL for the Tempest web interface
    name = "BASE_URL",
    type = (text = void),
    default = (text = "http://local.sandstorm.io"),
  ),
  ( # Port to listen on for (non-encrypted) HTTP
    name = "HTTP_PORT",
    type = (text = void),
    default = (text = "80"),
  ),
  ( # Port to listen on for HTTPS
    name = "HTTPS_PORT",
    type = (text = void),
    default = (text = "443"),
  ),
  ( # Path to HTTPS cert file
    name = "HTTPS_CERT_FILE",
    type = (text = void),
  ),
  ( # Path to HTTPS key file
    name = "HTTPS_KEY_FILE",
    type = (text = void),
  ),
  ( # when sending email, SMTP server to connect to
    name = "SMTP_HOST",
    type = (text = void),
  ),
  ( # port on `SMTP_HOST` to connect to. Can be a numeric port or a string
    # like "smtp"
    name = "SMTP_PORT",
    type = (text = void),
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
