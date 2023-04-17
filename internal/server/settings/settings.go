// Package settings provides support for working with the data in settings.capnp;
// see the comments in that file for an overview.
package settings

import (
	"math"
	"os"
	"strconv"

	"capnproto.org/go/capnp/v3/std/capnp/schema"
	"zenhack.net/go/tempest/capnp/settings"
	"zenhack.net/go/util"
)

var settingsInfo map[string]settings.Setting

func init() {
	settings.AdminSettings.Message().ResetReadLimit(math.MaxUint64)
	size := settings.AdminSettings.Len()
	settingsInfo = make(map[string]settings.Setting, size)
	for i := 0; i < size; i++ {
		setting := settings.AdminSettings.At(i)
		name, err := setting.Name()
		util.Chkfatal(err)
		settingsInfo[name] = setting
	}
}

// Read the environment variable specified in s.envVar
func getFromEnv(s settings.Setting) string {
	varName, err := s.EnvVar()
	util.Chkfatal(err)
	return os.Getenv(varName)
}

// Get the setting struct for the named setting, panicking if it does not have the
// specified type.
func getSettingInfo(name string, expectedType schema.Type_Which) settings.Setting {
	setting, ok := settingsInfo[name]
	if !ok {
		panic("No such setting: " + name)
	}
	typ, err := setting.Type()
	util.Chkfatal(err)
	if typ.Which() != expectedType {
		panic("setting " + name + " is not of type " + expectedType.String())
	}
	return setting
}

// Get the value of a setting by name, which must have type string (panics otherwise).
func GetString(name string) string {
	s := getSettingInfo(name, schema.Type_Which_text)
	val := getFromEnv(s)
	if val == "" && s.HasDefault() {
		def, err := s.Default()
		util.Chkfatal(err)
		val, err = def.Text()
		util.Chkfatal(err)
	}
	return val
}

// Get the value of a setting by name, which must have type uint16 (panics otherwise).
func GetUint16(name string) uint16 {
	s := getSettingInfo(name, schema.Type_Which_uint16)
	str := getFromEnv(s)
	if str == "" && s.HasDefault() {
		var err error
		val, err := s.Default()
		util.Chkfatal(err)
		return val.Uint16()
	}
	u64, err := strconv.ParseUint(str, 10, 16)
	util.Chkfatal(err)
	return uint16(u64)
}
