package config

import (
	"io"
	"strings"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var config Config

func InitConfig(configFile string) {
	envPrefix := "OP"
	envKeyReplacer := strings.NewReplacer(".", "_")

	if configFile != "" {
		config = NewConfig(configFile, envPrefix, envKeyReplacer)
	} else {
		// TODO: Add these config paths: $XDG_CONFIG_HOME/ and $HOME/.config/.
		config = NewConfigLookupPaths("config", envPrefix, envKeyReplacer, "/etc/onepaas", ".")
	}

	config.SetDefault("debug", true)
}

func GetConfig() Config {
	return config
}

// SetConfigFile explicitly defines the path, name and extension of the config file.
// Viper will use this and not check any of the config paths.
func SetConfigFile(in string) {
	config.SetConfigFile(in)
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
// E.g. if your prefix is "spf", the env registry will look for env
// variables that start with "SPF_".
func SetEnvPrefix(in string) {
	config.SetEnvPrefix(in)
}

// ConfigFileUsed returns the file used to populate the config registry.
func ConfigFileUsed() string {
	return config.ConfigFileUsed()
}

// AddConfigPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func AddConfigPath(in string) {
	config.AddConfigPath(in)
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} {
	return config.Get(key)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return config.GetString(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return config.GetBool(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return config.GetInt(key)
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 {
	return config.GetInt32(key)
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 {
	return config.GetInt64(key)
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint {
	return config.GetUint(key)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 {
	return config.GetUint32(key)
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 {
	return config.GetUint64(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 {
	return config.GetFloat64(key)
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time {
	return config.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration {
	return config.GetDuration(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string {
	return config.GetStringSlice(key)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} {
	return config.GetStringMap(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string {
	return config.GetStringMapString(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
	return config.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) uint {
	return config.GetSizeInBytes(key)
}

// UnmarshalKey takes a single key and unmarshals it into a Struct.
func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return config.UnmarshalKey(key, rawVal, opts...)
}

// Unmarshal unmarshals the config into a Struct. Make sure that the tags
// on the fields of the structure are properly set.
func Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return config.Unmarshal(rawVal, opts...)
}

// UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent
// in the destination struct.
func UnmarshalExact(rawVal interface{}) error {
	return config.UnmarshalExact(rawVal)
}

// BindEnv binds a Viper key to a ENV variable.
// ENV variables are case sensitive.
// If only a key is provided, it will use the env key matching the key, uppercased.
// EnvPrefix will be used when set when env name is not provided.
func BindEnv(input ...string) error {
	return config.BindEnv(input...)
}

// IsSet checks to see if the key has been set in any of the data locations.
// IsSet is case-insensitive for a key.
func IsSet(key string) bool {
	return config.IsSet(key)
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
// Useful for mapping an environmental variable to a key that does
// not match it.
func SetEnvKeyReplacer(r *strings.Replacer) {
	config.SetEnvKeyReplacer(r)
}

// InConfig checks to see if the given key (or an alias) is in the config file.
func InConfig(key string) bool {
	return config.InConfig(key)
}

// SetDefault sets the default value for this key.
// SetDefault is case-insensitive for a key.
// Default only used when no value is provided by the user via flag, config or ENV.
func SetDefault(key string, value interface{}) {
	config.SetDefault(key, value)
}

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
func Set(key string, value interface{}) {
	config.Set(key, value)
}

// ReadInConfig will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func ReadInConfig() error {
	return config.ReadInConfig()
}

// MergeInConfig merges a new configuration with an existing config.
func MergeInConfig() error {
	return config.MergeInConfig()
}

// ReadConfig will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func ReadConfig(in io.Reader) error {
	return config.ReadConfig(in)
}

// MergeConfig merges a new configuration with an existing config.
func MergeConfig(in io.Reader) error {
	return config.MergeConfig(in)
}

// MergeConfigMap merges the configuration from the map given with an existing config.
// Note that the map given may be modified.
func MergeConfigMap(cfg map[string]interface{}) error {
	return config.MergeConfigMap(cfg)
}

// AllKeys returns all keys holding a value, regardless of where they are set.
// Nested keys are returned with a v.keyDelim (= ".") separator
func AllKeys() []string {
	return config.AllKeys()
}

// AllSettings merges all settings and returns them as a map[string]interface{}.
func AllSettings() map[string]interface{} {
	return config.AllSettings()
}

// SetFs sets the filesystem to use to read configuration.
func SetFs(fs afero.Fs) {
	config.SetFs(fs)
}

// SetConfigName sets name for the config file.
// Does not include extension.
func SetConfigName(in string) {
	config.SetConfigName(in)
}

// SetConfigType sets the type of the configuration returned by the
// remote source, e.g. "json".
func SetConfigType(in string) {
	config.SetConfigType(in)
}
