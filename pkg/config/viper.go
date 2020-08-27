package config

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type safeConfig struct {
	*viper.Viper
	sync.RWMutex
}

// NewConfig returns a new Config object.
func NewConfig(file string, envPrefix string, envKeyReplacer *strings.Replacer) Config {
	c := safeConfig{
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
	c := safeConfig{
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
func (c *safeConfig) SetConfigFile(in string) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetConfigFile(in)
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
// E.g. if your prefix is "spf", the env registry will look for env
// variables that start with "SPF_".
func (c *safeConfig) SetEnvPrefix(in string) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetEnvPrefix(in)
}

// ConfigFileUsed returns the file used to populate the config registry.
func (c *safeConfig) ConfigFileUsed() string {
	c.Lock()
	defer c.Unlock()
	return c.Viper.ConfigFileUsed()
}

// AddConfigPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func (c *safeConfig) AddConfigPath(in string) {
	c.Lock()
	defer c.Unlock()
	c.Viper.AddConfigPath(in)
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func (c *safeConfig) Get(key string) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.Get(key)
}

// GetString returns the value associated with the key as a string.
func (c *safeConfig) GetString(key string) string {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetString(key)
}

// GetBool returns the value associated with the key as a boolean.
func (c *safeConfig) GetBool(key string) bool {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetBool(key)
}

// GetInt returns the value associated with the key as an integer.
func (c *safeConfig) GetInt(key string) int {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetInt(key)
}

// GetInt32 returns the value associated with the key as an integer.
func (c *safeConfig) GetInt32(key string) int32 {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetInt32(key)
}

// GetInt64 returns the value associated with the key as an integer.
func (c *safeConfig) GetInt64(key string) int64 {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetInt64(key)
}

// GetUint returns the value associated with the key as an unsigned integer.
func (c *safeConfig) GetUint(key string) uint {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetUint(key)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func (c *safeConfig) GetUint32(key string) uint32 {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetUint32(key)
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (c *safeConfig) GetUint64(key string) uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetUint64(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *safeConfig) GetFloat64(key string) float64 {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetFloat64(key)
}

// GetTime returns the value associated with the key as time.
func (c *safeConfig) GetTime(key string) time.Time {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func (c *safeConfig) GetDuration(key string) time.Duration {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetDuration(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *safeConfig) GetStringSlice(key string) []string {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetStringSlice(key)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *safeConfig) GetStringMap(key string) map[string]interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetStringMap(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *safeConfig) GetStringMapString(key string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetStringMapString(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *safeConfig) GetStringMapStringSlice(key string) map[string][]string {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func (c *safeConfig) GetSizeInBytes(key string) uint {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.GetSizeInBytes(key)
}

// UnmarshalKey takes a single key and unmarshals it into a Struct.
func (c *safeConfig) UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.UnmarshalKey(key, rawVal, opts...)
}

// Unmarshal unmarshals the config into a Struct. Make sure that the tags
// on the fields of the structure are properly set.
func (c *safeConfig) Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.Unmarshal(rawVal, opts...)
}

// UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent
// in the destination struct.
func (c *safeConfig) UnmarshalExact(rawVal interface{}) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.UnmarshalExact(rawVal)
}

// BindEnv binds a Viper key to a ENV variable.
// ENV variables are case sensitive.
// If only a key is provided, it will use the env key matching the key, uppercased.
// EnvPrefix will be used when set when env name is not provided.
func (c *safeConfig) BindEnv(input ...string) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.BindEnv(input...)
}

// IsSet checks to see if the key has been set in any of the data locations.
// IsSet is case-insensitive for a key.
func (c *safeConfig) IsSet(key string) bool {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.IsSet(key)
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
// Useful for mapping an environmental variable to a key that does
// not match it.
func (c *safeConfig) SetEnvKeyReplacer(r *strings.Replacer) {
	c.RLock()
	defer c.RUnlock()
	c.Viper.SetEnvKeyReplacer(r)
}

// InConfig checks to see if the given key (or an alias) is in the config file.
func (c *safeConfig) InConfig(key string) bool {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.InConfig(key)
}

// SetDefault sets the default value for this key.
// SetDefault is case-insensitive for a key.
// Default only used when no value is provided by the user via flag, config or ENV.
func (c *safeConfig) SetDefault(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetDefault(key, value)
}

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
func (c *safeConfig) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.Viper.Set(key, value)
}

// ReadInConfig will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func (c *safeConfig) ReadInConfig() error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.ReadInConfig()
}

// MergeInConfig merges a new configuration with an existing config.
func (c *safeConfig) MergeInConfig() error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.MergeInConfig()
}

// ReadConfig will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func (c *safeConfig) ReadConfig(in io.Reader) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.ReadConfig(in)
}

// MergeConfig merges a new configuration with an existing config.
func (c *safeConfig) MergeConfig(in io.Reader) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.MergeConfig(in)
}

// MergeConfigMap merges the configuration from the map given with an existing config.
// Note that the map given may be modified.
func (c *safeConfig) MergeConfigMap(cfg map[string]interface{}) error {
	c.Lock()
	defer c.Unlock()
	return c.Viper.MergeConfigMap(cfg)
}

// AllKeys returns all keys holding a value, regardless of where they are set.
// Nested keys are returned with a v.keyDelim (= ".") separator
func (c *safeConfig) AllKeys() []string {
	c.RLock()
	defer c.RUnlock()
	return c.Viper.AllKeys()
}

// AllSettings merges all settings and returns them as a map[string]interface{}.
func (c *safeConfig) AllSettings() map[string]interface{} {
	c.Lock()
	defer c.Unlock()
	return c.Viper.AllSettings()
}

// SetFs sets the filesystem to use to read configuration.
func (c *safeConfig) SetFs(fs afero.Fs) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetFs(fs)
}

// SetConfigName sets name for the config file.
// Does not include extension.
func (c *safeConfig) SetConfigName(in string) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetConfigName(in)
}

// SetConfigType sets the type of the configuration returned by the
// remote source, e.g. "json".
func (c *safeConfig) SetConfigType(in string) {
	c.Lock()
	defer c.Unlock()
	c.Viper.SetConfigType(in)
}
