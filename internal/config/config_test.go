package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolve(t *testing.T) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir: %v", err)
	}
	defaultHome := filepath.Join(userHome, ".docudex")

	tests := []struct {
		name         string
		envHome      string // "" means leave unset
		envProxy     string
		opts         Options
		wantHome     string
		wantProxyURL string
	}{
		{
			name:         "defaults when nothing set",
			wantHome:     defaultHome,
			wantProxyURL: DefaultProxyURL,
		},
		{
			name:         "empty env vars are treated as unset",
			envHome:      "",
			envProxy:     "",
			wantHome:     defaultHome,
			wantProxyURL: DefaultProxyURL,
		},
		{
			name:         "env overrides defaults",
			envHome:      "/tmp/docudex-home",
			envProxy:     "https://proxy.example.com",
			wantHome:     "/tmp/docudex-home",
			wantProxyURL: "https://proxy.example.com",
		},
		{
			name:         "file:// proxy passes through unchanged (fixture-proxy seam)",
			envProxy:     "file:///tmp/fixtures/proxy",
			wantHome:     defaultHome,
			wantProxyURL: "file:///tmp/fixtures/proxy",
		},
		{
			name:         "trailing slash trimmed from proxy URL",
			envProxy:     "https://proxy.golang.org/",
			wantHome:     defaultHome,
			wantProxyURL: "https://proxy.golang.org",
		},
		{
			name:         "flag options beat env and defaults",
			envHome:      "/tmp/env-home",
			envProxy:     "https://env.example.com",
			opts:         Options{Home: "/tmp/flag-home", ProxyURL: "https://flag.example.com"},
			wantHome:     "/tmp/flag-home",
			wantProxyURL: "https://flag.example.com",
		},
		{
			name:         "flag home only; proxy falls back to env",
			envProxy:     "https://env.example.com",
			opts:         Options{Home: "/tmp/flag-home"},
			wantHome:     "/tmp/flag-home",
			wantProxyURL: "https://env.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnvHome, tt.envHome)
			t.Setenv(EnvProxyURL, tt.envProxy)

			got, err := Resolve(tt.opts)
			if err != nil {
				t.Fatalf("Resolve: unexpected error: %v", err)
			}
			if got.Home != tt.wantHome {
				t.Errorf("Home = %q, want %q", got.Home, tt.wantHome)
			}
			if got.ProxyURL != tt.wantProxyURL {
				t.Errorf("ProxyURL = %q, want %q", got.ProxyURL, tt.wantProxyURL)
			}
		})
	}
}
