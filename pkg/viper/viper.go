package viper

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var sv *SafeViper

type SafeViper struct {
	*viper.Viper
	sync.RWMutex
}

func init() {
	New()
}

// New returns an initialized SafeViper instance.
func New() Config {
	sv = &SafeViper{
		Viper: viper.New(),
	}

	return sv
}

// NewConfig returns a new Config object.
func NewConfig(file string, envPrefix string, envKeyReplacer *strings.Replacer) Config {
	c := SafeViper{
		Viper: viper.New(),
	}

	c.SetConfigFile(file)
	c.SetEnvPrefix(envPrefix)
	c.SetEnvKeyReplacer(envKeyReplacer)
	c.SetTypeByDefaultValue(true)

	return &c
}

// NewConfig returns a new Config object.
func NewConfigLookupPaths(name string, envPrefix string, envKeyReplacer *strings.Replacer, configPaths ...string) Config {
	c := SafeViper{
		Viper: viper.New(),
	}

	c.SetConfigName(name)
	c.SetEnvPrefix(envPrefix)
	c.SetEnvKeyReplacer(envKeyReplacer)
	c.SetTypeByDefaultValue(true)

	for _, path := range configPaths {
		c.AddConfigPath(path)
	}

	if err := c.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Ignore error if Config file not found.
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	return &c
}

// SetConfigFile explicitly defines the path, name and extension of the config file.
// Viper will use this and not check any of the config paths.
func SetConfigFile(in string) { sv.SetConfigFile(in) }
func (sv *SafeViper) SetConfigFile(in string) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetConfigFile(in)
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
// E.g. if your prefix is "spf", the env registry will look for env
// variables that start with "SPF_".
func SetEnvPrefix(in string) { sv.SetEnvPrefix(in) }
func (sv *SafeViper) SetEnvPrefix(in string) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetEnvPrefix(in)
}

// ConfigFileUsed returns the file used to populate the config registry.
func ConfigFileUsed() string { return sv.ConfigFileUsed() }
func (sv *SafeViper) ConfigFileUsed() string {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.ConfigFileUsed()
}

// AddConfigPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func AddConfigPath(in string) { sv.AddConfigPath(in) }
func (sv *SafeViper) AddConfigPath(in string) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.AddConfigPath(in)
}

// SetTypeByDefaultValue enables or disables the inference of a key value's
// type when the Get function is used based upon a key's default value as
// opposed to the value returned based on the normal fetch logic.
//
// For example, if a key has a default value of []string{} and the same key
// is set via an environment variable to "a b c", a call to the Get function
// would return a string slice for the key if the key's type is inferred by
// the default value and the Get function would return:
//
//   []string {"a", "b", "c"}
//
// Otherwise the Get function would return:
//
//   "a b c"
func SetTypeByDefaultValue(enable bool) { sv.SetTypeByDefaultValue(enable) }
func (sv *SafeViper) SetTypeByDefaultValue(enable bool) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetTypeByDefaultValue(enable)
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} { return sv.Get(key) }
func (sv *SafeViper) Get(key string) interface{} {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.Get(key)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string { return sv.GetString(key) }
func (sv *SafeViper) GetString(key string) string {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetString(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool { return sv.GetBool(key) }
func (sv *SafeViper) GetBool(key string) bool {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetBool(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int { return sv.GetInt(key) }
func (sv *SafeViper) GetInt(key string) int {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetInt(key)
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return sv.GetInt32(key) }
func (sv *SafeViper) GetInt32(key string) int32 {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetInt32(key)
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return sv.GetInt64(key) }
func (sv *SafeViper) GetInt64(key string) int64 {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetInt64(key)
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint { return sv.GetUint(key) }
func (sv *SafeViper) GetUint(key string) uint {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetUint(key)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 { return sv.GetUint32(key) }
func (sv *SafeViper) GetUint32(key string) uint32 {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetUint32(key)
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 { return sv.GetUint64(key) }
func (sv *SafeViper) GetUint64(key string) uint64 {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetUint64(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 { return sv.GetFloat64(key) }
func (sv *SafeViper) GetFloat64(key string) float64 {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetFloat64(key)
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time { return sv.GetTime(key) }
func (sv *SafeViper) GetTime(key string) time.Time {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration { return sv.GetDuration(key) }
func (sv *SafeViper) GetDuration(key string) time.Duration {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetDuration(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string { return sv.GetStringSlice(key) }
func (sv *SafeViper) GetStringSlice(key string) []string {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetStringSlice(key)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} { return sv.GetStringMap(key) }
func (sv *SafeViper) GetStringMap(key string) map[string]interface{} {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetStringMap(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string { return sv.GetStringMapString(key) }
func (sv *SafeViper) GetStringMapString(key string) map[string]string {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetStringMapString(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string { return sv.GetStringMapStringSlice(key) }
func (sv *SafeViper) GetStringMapStringSlice(key string) map[string][]string {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) uint { return sv.GetSizeInBytes(key) }
func (sv *SafeViper) GetSizeInBytes(key string) uint {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.GetSizeInBytes(key)
}

// UnmarshalKey takes a single key and unmarshals it into a Struct.
func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error { return sv.UnmarshalKey(key, rawVal, opts...) }
func (sv *SafeViper) UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.UnmarshalKey(key, rawVal, opts...)
}

// Unmarshal unmarshals the config into a Struct. Make sure that the tags
// on the fields of the structure are properly set.
func Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error { return sv.Unmarshal(rawVal, opts...) }
func (sv *SafeViper) Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.Unmarshal(rawVal, opts...)
}

// UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent
// in the destination struct.
func UnmarshalExact(rawVal interface{}) error { return sv.UnmarshalExact(rawVal) }
func (sv *SafeViper) UnmarshalExact(rawVal interface{}) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.UnmarshalExact(rawVal)
}

// BindEnv binds a Viper key to a ENV variable.
// ENV variables are case sensitive.
// If only a key is provided, it will use the env key matching the key, uppercased.
// EnvPrefix will be used when set when env name is not provided.
func BindEnv(input ...string) error { return sv.BindEnv(input...) }
func (sv *SafeViper) BindEnv(input ...string) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.BindEnv(input...)
}

// IsSet checks to see if the key has been set in any of the data locations.
// IsSet is case-insensitive for a key.
func IsSet(key string) bool { return sv.IsSet(key) }
func (sv *SafeViper) IsSet(key string) bool {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.IsSet(key)
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
// Useful for mapping an environmental variable to a key that does
// not match it.
func SetEnvKeyReplacer(r *strings.Replacer) { sv.SetEnvKeyReplacer(r) }
func (sv *SafeViper) SetEnvKeyReplacer(r *strings.Replacer) {
	sv.RLock()
	defer sv.RUnlock()
	sv.Viper.SetEnvKeyReplacer(r)
}

// InConfig checks to see if the given key (or an alias) is in the config file.
func InConfig(key string) bool { return sv.InConfig(key) }
func (sv *SafeViper) InConfig(key string) bool {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.InConfig(key)
}

// SetDefault sets the default value for this key.
// SetDefault is case-insensitive for a key.
// Default only used when no value is provided by the user via flag, config or ENV.
func SetDefault(key string, value interface{}) { sv.SetDefault(key, value) }
func (sv *SafeViper) SetDefault(key string, value interface{}) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetDefault(key, value)
}

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
func Set(key string, value interface{}) { sv.Set(key, value) }
func (sv *SafeViper) Set(key string, value interface{}) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.Set(key, value)
}

// ReadInConfig will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func ReadInConfig() error { return sv.ReadInConfig() }
func (sv *SafeViper) ReadInConfig() error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.ReadInConfig()
}

// MergeInConfig merges a new configuration with an existing config.
func MergeInConfig() error { return sv.MergeInConfig() }
func (sv *SafeViper) MergeInConfig() error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.MergeInConfig()
}

// ReadConfig will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func ReadConfig(in io.Reader) error { return sv.ReadConfig(in) }
func (sv *SafeViper) ReadConfig(in io.Reader) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.ReadConfig(in)
}

// MergeConfig merges a new configuration with an existing config.
func MergeConfig(in io.Reader) error { return sv.MergeConfig(in) }
func (sv *SafeViper) MergeConfig(in io.Reader) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.MergeConfig(in)
}

// MergeConfigMap merges the configuration from the map given with an existing config.
// Note that the map given may be modified.
func MergeConfigMap(cfg map[string]interface{}) error { return sv.MergeConfigMap(cfg) }
func (sv *SafeViper) MergeConfigMap(cfg map[string]interface{}) error {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.MergeConfigMap(cfg)
}

// AllKeys returns all keys holding a value, regardless of where they are set.
// Nested keys are returned with a v.keyDelim (= ".") separator
func AllKeys() []string { return sv.AllKeys() }
func (sv *SafeViper) AllKeys() []string {
	sv.RLock()
	defer sv.RUnlock()
	return sv.Viper.AllKeys()
}

// AllSettings merges all settings and returns them as a map[string]interface{}.
func AllSettings() map[string]interface{} { return sv.AllSettings() }
func (sv *SafeViper) AllSettings() map[string]interface{} {
	sv.Lock()
	defer sv.Unlock()
	return sv.Viper.AllSettings()
}

// SetFs sets the filesystem to use to read configuration.
func SetFs(fs afero.Fs) { sv.SetFs(fs) }
func (sv *SafeViper) SetFs(fs afero.Fs) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetFs(fs)
}

// SetConfigName sets name for the config file.
// Does not include extension.
func SetConfigName(in string) { sv.SetConfigName(in) }
func (sv *SafeViper) SetConfigName(in string) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetConfigName(in)
}

// SetConfigType sets the type of the configuration returned by the
// remote source, e.g. "json".
func SetConfigType(in string) { sv.SetConfigType(in) }
func (sv *SafeViper) SetConfigType(in string) {
	sv.Lock()
	defer sv.Unlock()
	sv.Viper.SetConfigType(in)
}
