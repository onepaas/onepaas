package config

import (
	"io"
	"strings"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Config interface {
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	SetConfigFile(in string)

	// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
	// E.g. if your prefix is "spf", the env registry will look for env
	// variables that start with "SPF_".
	SetEnvPrefix(in string)

	// ConfigFileUsed returns the file used to populate the config registry.
	ConfigFileUsed() string

	// AddConfigPath adds a path for Viper to search for the config file in.
	// Can be called multiple times to define multiple search paths.
	AddConfigPath(in string)

	// Get can retrieve any value given the key to use.
	// Get is case-insensitive for a key.
	// Get has the behavior of returning the value associated with the first
	// place from where it is set. Viper will check in the following order:
	// override, flag, env, config file, key/value store, default
	//
	// Get returns an interface. For a specific value use one of the Get____ methods.
	Get(key string) interface{}

	// GetString returns the value associated with the key as a string.
	GetString(key string) string

	// GetBool returns the value associated with the key as a boolean.
	GetBool(key string) bool

	// GetInt returns the value associated with the key as an integer.
	GetInt(key string) int

	// GetInt32 returns the value associated with the key as an integer.
	GetInt32(key string) int32

	// GetInt64 returns the value associated with the key as an integer.
	GetInt64(key string) int64

	// GetUint returns the value associated with the key as an unsigned integer.
	GetUint(key string) uint

	// GetUint32 returns the value associated with the key as an unsigned integer.
	GetUint32(key string) uint32

	// GetUint64 returns the value associated with the key as an unsigned integer.
	GetUint64(key string) uint64

	// GetFloat64 returns the value associated with the key as a float64.
	GetFloat64(key string) float64

	// GetTime returns the value associated with the key as time.
	GetTime(key string) time.Time

	// GetDuration returns the value associated with the key as a duration.
	GetDuration(key string) time.Duration

	// GetStringSlice returns the value associated with the key as a slice of strings.
	GetStringSlice(key string) []string

	// GetStringMap returns the value associated with the key as a map of interfaces.
	GetStringMap(key string) map[string]interface{}

	// GetStringMapString returns the value associated with the key as a map of strings.
	GetStringMapString(key string) map[string]string

	// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
	GetStringMapStringSlice(key string) map[string][]string

	// GetSizeInBytes returns the size of the value associated with the given key
	// in bytes.
	GetSizeInBytes(key string) uint

	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error

	// Unmarshal unmarshals the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error

	// UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent
	// in the destination struct.
	UnmarshalExact(rawVal interface{}) error

	// BindEnv binds a Viper key to a ENV variable.
	// ENV variables are case sensitive.
	// If only a key is provided, it will use the env key matching the key, uppercased.
	// EnvPrefix will be used when set when env name is not provided.
	BindEnv(input ...string) error

	// IsSet checks to see if the key has been set in any of the data locations.
	// IsSet is case-insensitive for a key.
	IsSet(key string) bool

	// SetEnvKeyReplacer sets the strings.Replacer on the viper object
	// Useful for mapping an environmental variable to a key that does
	// not match it.
	SetEnvKeyReplacer(r *strings.Replacer)

	// InConfig checks to see if the given key (or an alias) is in the config file.
	InConfig(key string) bool

	// SetDefault sets the default value for this key.
	// SetDefault is case-insensitive for a key.
	// Default only used when no value is provided by the user via flag, config or ENV.
	SetDefault(key string, value interface{})

	// Set sets the value for the key in the override register.
	// Set is case-insensitive for a key.
	// Will be used instead of values obtained via
	// flags, config file, ENV, default, or key/value store.
	Set(key string, value interface{})

	// ReadInConfig will discover and load the configuration file from disk
	// and key/value stores, searching in one of the defined paths.
	ReadInConfig() error

	// MergeInConfig merges a new configuration with an existing config.
	MergeInConfig() error

	// ReadConfig will read a configuration file, setting existing keys to nil if the
	// key does not exist in the file.
	ReadConfig(in io.Reader) error

	// MergeConfig merges a new configuration with an existing config.
	MergeConfig(in io.Reader) error

	// MergeConfigMap merges the configuration from the map given with an existing config.
	// Note that the map given may be modified.
	MergeConfigMap(cfg map[string]interface{}) error

	// AllKeys returns all keys holding a value, regardless of where they are set.
	// Nested keys are returned with a v.keyDelim (= ".") separator
	AllKeys() []string

	// AllSettings merges all settings and returns them as a map[string]interface{}.
	AllSettings() map[string]interface{}

	// SetFs sets the filesystem to use to read configuration.
	SetFs(fs afero.Fs)

	// SetConfigName sets name for the config file.
	// Does not include extension.
	SetConfigName(in string)

	// SetConfigType sets the type of the configuration returned by the
	// remote source, e.g. "json".
	SetConfigType(in string)
}
