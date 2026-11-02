package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetGoModCachePath() string {
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

func GetGoProxy() string {
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
