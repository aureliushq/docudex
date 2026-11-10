package fetch

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func useFixtureProxy(t *testing.T) {
	t.Helper()

	abs, err := filepath.Abs("testdata/proxy")
	if err != nil {
		t.Fatal(err)
	}

	t.Setenv("GOPROXY", "file://"+filepath.ToSlash(abs))
	t.Setenv("GOSUMDB", "off")
	t.Setenv("GOMODCACHE", t.TempDir())
	t.Setenv("GOFLAGS", "-modcacherw")
}

func mustExist(t *testing.T, path string) {
	t.Helper()

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

func TestDownloadExactVersion(t *testing.T) {
	useFixtureProxy(t)

	mod, err := DownloadModuleWithVersion("example.com/hello", "v1.0.0")
	if err != nil {
		t.Fatal(err)
	}
	if mod.Version != "v1.0.0" {
		t.Errorf("Version = %q, want v1.0.0", mod.Version)
	}
	mustExist(t, filepath.Join(mod.Dir, "hello.go"))
}

func TestDownloadLatest(t *testing.T) {
	useFixtureProxy(t)

	mod, err := DownloadModuleWithVersion("example.com/hello", "")
	if err != nil {
		t.Fatal(err)
	}
	if mod.Version != "v1.1.0" {
		t.Errorf("Version = %q, want v1.1.0 (from @latest)", mod.Version)
	}
}

func TestTwoVersionsSideBySide(t *testing.T) {
	useFixtureProxy(t)

	a, err := DownloadModuleWithVersion("example.com/hello", "v1.0.0")
	if err != nil {
		t.Fatal(err)
	}
	b, err := DownloadModuleWithVersion("example.com/hello", "v1.1.0")
	if err != nil {
		t.Fatal(err)
	}
	if a.Dir == b.Dir {
		t.Fatalf("both versions share Dir %s", a.Dir)
	}
	mustExist(t, a.Dir) // first download survives the second
	mustExist(t, b.Dir)
}

func TestDownloadUnknownModule(t *testing.T) {
	useFixtureProxy(t)

	_, err := DownloadModuleWithVersion("example.com/nope", "v1.0.0")
	if err == nil {
		t.Fatal("expected error for unknown module")
	}
	if !strings.Contains(err.Error(), "example.com/nope") {
		t.Errorf("error lacks go's message: %v", err)
	}
}
