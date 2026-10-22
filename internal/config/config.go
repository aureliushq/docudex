package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	Title          string `mapstructure:"title"`
	ModulePath     string `mapstructure:"module_path"`
	GoVersion      string `mapstructure:"go_version"`
	GoModCachePath string `mapstructure:"go_mod_cache_path"`
	GoProxy        string `mapstructure:"go_proxy"`

	Dependencies Dependencies `mapstructure:"dependencies"`

	Viper *viper.Viper `mapstructure:"-"`
}

func NewConfig() Config {
	goModCachePath := getGoModCachePath()
	goProxy := getGoProxy()

	cfg := Config{
		GoModCachePath: goModCachePath,
		GoProxy:        goProxy,

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
		fmt.Printf("Error decoding struct to map: %s\n", err)
		return err
	}

	err = cfg.Viper.MergeConfigMap(cfgMap)
	if err != nil {
		fmt.Printf("Error merging config map: %s\n", err)
		return err
	}

	err = cfg.Viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("Error writing configuration file: %s\n", err)
		return err
	}

	return nil
}

func getGoModCachePath() string {
	envCachePath := os.Getenv("GOMODCACHE")
	if envCachePath != "" {
		return envCachePath
	} else {
		envCachePath, err := exec.Command("go", "env", "GOMODCACHE").Output()
		if err != nil {
			fmt.Printf("Error getting GOMODCACHE: %s", err)
		}
		return strings.TrimSpace(string(envCachePath))
	}
}

func getGoProxy() string {
	goProxy := os.Getenv("GOPROXY")
	if goProxy != "" {
		return goProxy
	} else {
		goProxy, err := exec.Command("go", "env", "GOPROXY").Output()
		if err != nil {
			fmt.Printf("Error getting GOPROXY: %s", err)
		}
		return strings.TrimSpace(string(goProxy))
	}
}
