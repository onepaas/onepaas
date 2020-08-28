package config

import (
	"fmt"
	"strings"

	"github.com/onepaas/onepaas/pkg/viper"
	v "github.com/spf13/viper"
)

func InitConfig(configFile string) {
	viper.SetEnvPrefix("OP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		// TODO: Add these config paths: $XDG_CONFIG_HOME/ and $HOME/.config/.
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/onepaas")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(v.ConfigFileNotFoundError); ok {
				// Ignore error if Config file not found.
			} else {
				// Config file was found but another error was produced
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		}
	}

	viper.SetDefault("debug", true)
}
