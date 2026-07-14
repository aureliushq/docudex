// Package config resolves docudex's runtime configuration — the store/home
// directory and the module-proxy base URL — from CLI flags, environment
// variables, and built-in defaults.
//
// These two values are the load-bearing seams of docudex's test strategy: the
// proxy URL lets tests point at a file:// fixture proxy instead of the network,
// and the home directory lets each test run against an isolated temporary store.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config is the resolved runtime configuration.
type Config struct {
	// Home is the docudex store/home directory (default ~/.docudex).
	Home string
	// ProxyURL is the base URL of the Go module proxy (default
	// https://proxy.golang.org). May be a file:// URL for fixture proxies.
	ProxyURL string
}

// Options carries per-invocation overrides, typically sourced from CLI flags.
// An empty string means "not set" and defers to the environment or default.
type Options struct {
	Home     string
	ProxyURL string
}

// Environment variables read by Resolve.
const (
	EnvHome     = "DOCUDEX_HOME"
	EnvProxyURL = "DOCUDEX_PROXY_URL"
)

const (
	// DefaultProxyURL is the public Go module proxy used when none is configured.
	DefaultProxyURL = "https://proxy.golang.org"

	// homeDirName is the default store directory name under the user's home.
	homeDirName = ".docudex"
)

// Resolve builds a Config using the precedence flag > environment > default,
// applied independently to each field.
func Resolve(opts Options) (Config, error) {
	home, err := resolveHome(opts.Home)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Home:     home,
		ProxyURL: resolveProxyURL(opts.ProxyURL),
	}, nil
}

func resolveHome(flag string) (string, error) {
	if v := firstNonEmpty(flag, os.Getenv(EnvHome)); v != "" {
		return v, nil
	}

	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving default home directory: %w", err)
	}
	return filepath.Join(userHome, homeDirName), nil
}

func resolveProxyURL(flag string) string {
	v := firstNonEmpty(flag, os.Getenv(EnvProxyURL), DefaultProxyURL)
	return strings.TrimRight(v, "/")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
