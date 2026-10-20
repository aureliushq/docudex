/*
Copyright © 2026 Ilango Rajagopal <hey@i4o.dev>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aureliushq/docudex/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

type RequiredDependency struct {
	Path         string
	Version      string
	ModCachePath string
}

func NewRequiredDependency() RequiredDependency {
	return RequiredDependency{
		Path:         "",
		Version:      "",
		ModCachePath: "",
	}
}

type ModInfo struct {
	Path                 string
	GoVersion            string
	RequiredDependencies []RequiredDependency
}

func NewModInfo() ModInfo {
	return ModInfo{
		Path:                 "",
		GoVersion:            "",
		RequiredDependencies: []RequiredDependency{},
	}
}

func newInitCommand() *cobra.Command {
	modInfo := ModInfo{}
	// initCmd represents the init command
	initCmd := &cobra.Command{
		Use:   "init",
		Args:  cobra.NoArgs,
		Short: "Create docudex.toml from the project's go.mod",
		Long: "Read the project's go.mod, materialize a docudex.toml listing the\n" +
			"direct dependencies (plus the Go standard library) at their exact\n" +
			"versions, and fetch their docs into the store.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.NewConfig()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runInit(&modInfo)
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

func runInit(mi *ModInfo) error {
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
	mi.Path = goMod.Module.Mod.Path
	mi.GoVersion = goMod.Go.Version
	for _, req := range goMod.Require {
		if !req.Indirect {
			rd := NewRequiredDependency()
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
			rd.Path = escapedPath
			rd.Version = escapedVersion
			if err = module.Check(escapedPath, escapedVersion); err != nil {
				fmt.Printf("Path is invalid for module: %s, version: %s, error: %s", escapedPath, escapedVersion, err)
				return err
			}
			fmt.Printf("%s@%s\n", escapedPath, escapedVersion)
			mi.RequiredDependencies = append(mi.RequiredDependencies, rd)
		}
	}

	fmt.Printf("ModInfo: %v\n", mi)

	return nil
}
