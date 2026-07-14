// Package cli defines docudex's command tree: the root command, its
// stub subcommands, and the persistent flags that feed configuration resolution.
package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/aureliushq/docudex/internal/config"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// NewRootCmd builds the docudex root command with all subcommands registered.
// It is a factory (rather than a package-level singleton) so tests can drive a
// fresh command tree with isolated I/O and without calling os.Exit.
func NewRootCmd() *cobra.Command {
	var opts config.Options

	root := &cobra.Command{
		Use:   "docudex",
		Short: "Per-project, version-exact Go API reference docs",
		Long: "docudex aggregates version-exact API reference documentation for a\n" +
			"Go project's dependencies into a local store you can search from the\n" +
			"terminal or browse in a local web UI — scoped to exactly what your\n" +
			"project uses, and fully available offline.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// Persistent flags feed the two configuration seams. They override the
	// DOCUDEX_HOME / DOCUDEX_PROXY_URL environment variables when set.
	root.PersistentFlags().StringVar(&opts.Home, "home", "",
		"docudex store/home directory (overrides $"+config.EnvHome+", default ~/.docudex)")
	root.PersistentFlags().StringVar(&opts.ProxyURL, "proxy-url", "",
		"Go module proxy base URL (overrides $"+config.EnvProxyURL+", default "+config.DefaultProxyURL+")")

	// resolveConfig is closed over the flag-backed opts so every subcommand
	// resolves configuration the same way once real behaviour lands.
	resolveConfig := func() (config.Config, error) { return config.Resolve(opts) }

	root.AddCommand(
		newInitCmd(resolveConfig),
		newAddCmd(resolveConfig),
		newRemoveCmd(resolveConfig),
		newSyncCmd(resolveConfig),
		newListCmd(resolveConfig),
		newSearchCmd(resolveConfig),
		newServeCmd(resolveConfig),
		newRebuildIndexCmd(resolveConfig),
	)

	return root
}

// Execute builds and runs the root command, exiting non-zero on error. It is
// the single entrypoint called by main.
func Execute() {
	cmd := NewRootCmd()
	if err := fang.Execute(context.Background(), cmd); err != nil {
		fmt.Fprintln(os.Stderr, "docudex:", err)
		os.Exit(1)
	}
}

// configResolver resolves the effective configuration for a command run.
type configResolver func() (config.Config, error)

// stubRunE returns a RunE that announces the command is not yet implemented and
// prints its usage, exiting cleanly. Subcommands share it until real behaviour
// lands in later tickets.
func stubRunE(name string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.Printf("docudex %s: not yet implemented\n\n", name)
		return cmd.Usage()
	}
}
