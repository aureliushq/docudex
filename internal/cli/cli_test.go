package cli

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// runRoot executes a fresh root command tree with the given args, capturing
// stdout+stderr, and returns the combined output and any execution error.
func runRoot(t *testing.T, args ...string) (string, error) {
	t.Helper()

	root := NewRootCmd()
	var buf bytes.Buffer
	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}

// stubCommands is every subcommand that should exist and behave as a clean stub.
var stubCommands = []string{
	"init", "add", "remove", "sync", "list", "search", "serve", "rebuild-index",
}

func TestStubCommandsPrintUsageAndExitCleanly(t *testing.T) {
	for _, name := range stubCommands {
		t.Run(name, func(t *testing.T) {
			out, err := runRoot(t, name)
			if err != nil {
				t.Fatalf("docudex %s: unexpected error: %v", name, err)
			}
			if !strings.Contains(out, "not yet implemented") {
				t.Errorf("docudex %s: output missing not-implemented notice:\n%s", name, out)
			}
			if !strings.Contains(out, "Usage:") {
				t.Errorf("docudex %s: output missing usage text:\n%s", name, out)
			}
		})
	}
}

func TestRootCommandShowsHelp(t *testing.T) {
	out, err := runRoot(t, "--help")
	if err != nil {
		t.Fatalf("docudex --help: unexpected error: %v", err)
	}
	for _, name := range stubCommands {
		if !strings.Contains(out, name) {
			t.Errorf("docudex --help: expected subcommand %q listed:\n%s", name, out)
		}
	}
}

func TestPersistentConfigFlagsAreRegistered(t *testing.T) {
	root := NewRootCmd()
	for _, flag := range []string{"home", "proxy-url"} {
		if root.PersistentFlags().Lookup(flag) == nil {
			t.Errorf("expected persistent --%s flag to be registered", flag)
		}
	}
}

// TestConfigFlagsThreadThroughToParsing proves the persistent --home/--proxy-url
// flags parse end-to-end through a real Execute (with a subcommand), so the same
// values reach the config resolver. Precedence itself is covered in the config
// package; this closes the CLI-wiring half of the seam.
func TestConfigFlagsThreadThroughToParsing(t *testing.T) {
	root := NewRootCmd()
	root.SetOut(io.Discard)
	root.SetErr(io.Discard)
	root.SetArgs([]string{"--home", "/tmp/flag-home", "--proxy-url", "https://flag.example.com", "list"})

	if err := root.Execute(); err != nil {
		t.Fatalf("Execute with config flags: %v", err)
	}

	if got, _ := root.PersistentFlags().GetString("home"); got != "/tmp/flag-home" {
		t.Errorf("--home parsed as %q, want %q", got, "/tmp/flag-home")
	}
	if got, _ := root.PersistentFlags().GetString("proxy-url"); got != "https://flag.example.com" {
		t.Errorf("--proxy-url parsed as %q, want %q", got, "https://flag.example.com")
	}
}

func TestUnknownCommandErrors(t *testing.T) {
	if _, err := runRoot(t, "no-such-command"); err == nil {
		t.Error("expected error for unknown command, got nil")
	}
}
