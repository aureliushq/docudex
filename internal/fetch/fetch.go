package fetch

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"os/exec"
)

type Origin struct {
	VCS  string `json:"VCS"`
	URL  string `json:"URL"`
	Hash string `json:"Hash"`
	Ref  string `json:"Ref"`
}

type Module struct {
	Path string `json:"Path"` // module path
	// Query    string `json:"Query"`    // version query corresponding to this version
	Version string `json:"Version"` // module version
	Error   string `json:"Error"`   // error loading module
	// Info     string `json:"Info"`     // absolute path to cached .info file
	// GoMod    string `json:"GoMod"`    // absolute path to cached .mod file
	Zip string `json:"Zip"` // absolute path to cached .zip file
	Dir string `json:"Dir"` // absolute path to cached source root directory
	// Sum      string `json:"Sum"`      // checksum for path, version (as in go.sum)
	// GoModSum string `json:"GoModSum"` // checksum for go.mod (as in go.sum)
	Origin Origin `json:"Origin"` // provenance of module
	// Reuse    bool   `json:"Reuse"`    // reuse of old module info is safe
}

func DownloadModuleWithVersion(module, version string) (Module, error) {
	if version == "" {
		version = "latest"
	}
	modulePath := fmt.Sprintf("%s@%s", module, version)
	resp, runErr := exec.Command("go", "mod", "download", "-json", modulePath).Output()
	var mod Module
	if err := json.Unmarshal(resp, &mod); err != nil {
		return Module{}, errors.Join(runErr, err)
	}
	if mod.Error != "" {
		return Module{}, errors.New(mod.Error)
	}
	return mod, runErr
}
