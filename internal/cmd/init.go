/*
Copyright © 2026 Ilango Rajagopal <hey@i4o.dev>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aureliushq/docudex/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

var (
	cfg          config.Config
	dependencies config.Dependencies
)

func newInitCommand() *cobra.Command {
	// initCmd represents the init command
	initCmd := &cobra.Command{
		Use:   "init",
		Args:  cobra.NoArgs,
		Short: "Create docudex.toml from the project's go.mod",
		Long: "Read the project's go.mod, materialize a docudex.toml listing the\n" +
			"direct dependencies (plus the Go standard library) at their exact\n" +
			"versions, and fetch their docs into the store.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cfg = config.NewConfig()
			dependencies = make(map[config.DependencyPath]config.DependencyMetadata)
			if err := cfg.SetupConfig(); err == nil {
				// TODO: ask user if they want to overwrite and ask them to use a flag (TBD)
				// TODO: suggest running sync
				// return errors.New("Found existing configuration file: docudex.toml")
				return nil
			} else {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					return nil
				}
				return err
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runInit()
			return err
		},
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return initCmd
}

func runInit() error {
	file, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Printf("Error reading go.mod: %s", err)
		return err
	}
	goMod, err := modfile.Parse("go.mod", file, nil)
	if err != nil {
		fmt.Printf("Error parsing go.mod: %s", err)
		return err
	}
	for _, req := range goMod.Require {
		if !req.Indirect {
			if err = module.Check(req.Mod.Path, req.Mod.Version); err != nil {
				fmt.Printf("Path is invalid for module: %s, version: %s, error: %s", req.Mod.Path, req.Mod.Version, err)
				return err
			}
			escapedPath, err := module.EscapePath(req.Mod.Path)
			if err != nil {
				fmt.Printf("Error escaping module path for: %s, error: %s", req.Mod.Path, err)
				return err
			}
			escapedVersion, err := module.EscapeVersion(req.Mod.Version)
			if err != nil {
				fmt.Printf("Error escaping module: %s, version: %s, error: %s", req.Mod.Path, req.Mod.Version, err)
				return err
			}
			dependencies[config.DependencyPath(escapedPath)] = config.DependencyMetadata{Default: false, Version: escapedVersion}
		}
	}

	cfg.ModulePath = goMod.Module.Mod.Path
	cfg.GoVersion = goMod.Go.Version
	cfg.Dependencies = dependencies

	err = cfg.WriteConfig()
	if err != nil {
		fmt.Printf("Error writing configuration file: %s", err)
		return err
	}

	return nil
}
