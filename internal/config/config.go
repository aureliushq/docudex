package config

import (
	"github.com/aureliushq/docudex/internal"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

const DefaultConfigDir = "."
const DefaultConfigFile = "docudex.toml"

type DependencyPath string

type DependencyMetadata struct {
	Default bool   `toml:"default"`
	Version string `toml:"version"`
}

type Dependencies map[DependencyPath]DependencyMetadata

type Config struct {
	Title      string `mapstructure:"title"`
	ModulePath string `mapstructure:"module_path"`
	GoVersion  string `mapstructure:"go_version"`
	GoModCache string `mapstructure:"go_mod_cache"`
	GoProxy    string `mapstructure:"go_proxy"`

	Dependencies Dependencies `mapstructure:"dependencies"`

	Viper *viper.Viper `mapstructure:"-"`
}

func NewConfig() Config {
	goModCache := internal.GetGoModCachePath()
	goProxy := internal.GetGoProxy()

	cfg := Config{
		GoModCache: goModCache,
		GoProxy:    goProxy,

		Viper: viper.New(),
	}

	return cfg
}

func (cfg *Config) SetupConfig() error {
	cfg.Viper.SetConfigName("docudex")
	cfg.Viper.SetConfigType("toml")

	cfg.Viper.AddConfigPath(".")

	err := cfg.Viper.ReadInConfig()

	return err
}

func (cfg *Config) WriteConfig() error {
	cfgMap := map[string]any{}
	err := mapstructure.Decode(cfg, &cfgMap)
	if err != nil {
		return err
	}

	err = cfg.Viper.MergeConfigMap(cfgMap)
	if err != nil {
		return err
	}

	err = cfg.Viper.SafeWriteConfig()
	if err != nil {
		return err
	}

	return nil
}
