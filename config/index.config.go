package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

func Config(key string) string {
	return os.Getenv(key)
}
