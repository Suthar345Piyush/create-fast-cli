// viper based read/write configs
// some based on user preferences
// config file is  - fastcli.yaml

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// setting up viper configs for read and write ~/.fastcli.yaml

func InitViper() error {

	// config path

	home, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("could not determine the home directory: %w", err)
	}

	viper.SetConfigName(".fastcli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	viper.SetEnvPrefix("FASTCLI")
	viper.AutomaticEnv()

	// some defaults

	viper.SetDefault("preferred_ide", string(IDEVscode))
	viper.SetDefault("defualt_framework", string(FrameworkCobra))
	viper.SetDefault("output_dir", filepath.Join(home, "projects"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config: %w", err)
		}
	}

	return nil

}

// saving some user preferences

func SavePreferences(cfg *ProjectConfig) error {

	viper.Set("preferred_ide", string(cfg.IDE))
	viper.Set("default_framework", string(cfg.Framework))

	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	return viper.WriteConfigAs(filepath.Join(home, ".fastcli.yaml"))

}

// these will return the values associated with this keys

// ide

func PreferredIDE() IDE {
	return IDE(viper.GetString("preferred_ide"))
}

// framework

func DefaultFramework() Framework {
	return Framework(viper.GetString("default_framework"))
}

// default output directory

func DefaultOutputDir() string {
	return viper.GetString("output_dir")
}
