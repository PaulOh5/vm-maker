package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type BinPathBook struct {
	Mkdosfs         string
	Mcopy           string
	CloudHyperVisor string
}

type Settings struct {
	DiskPath         string
	CloudInitDirPath string
	TempDirPath      string
	BinPathBook
}

func SetupSettings() Settings {
	_, currentPath, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(currentPath)
	envPath := filepath.Join(basePath, "../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	settings := Settings{
		TempDirPath: isAccessibel(getEnv("TEMP_DIR_PATH", "/tmp")),
		BinPathBook: BinPathBook{
			Mkdosfs:         isAccessibel(getEnv("MKDOSFS_BIN_PATH", "/usr/sbin/mkdosfs")),
			Mcopy:           isAccessibel(getEnv("MCOPY_BIN_PATH", "/usr/bin/mcopy")),
			CloudHyperVisor: isAccessibel(getEnv("CLOUD_HYPERVISOR_BIN_PATH", "/usr/local/bin/cloud-hypervisor")),
		},
	}

	return settings
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func isAccessibel(path string) string {
	if _, err := os.Stat(path); err != nil {
		switch {
		case os.IsNotExist(err):
			log.Fatalf("Path %s does not exist", path)
		case os.IsPermission(err):
			log.Fatalf("Permission denied to access path %s", path)
		default:
			log.Fatalf("Error accessing path %s: %v", path, err)
		}
	}

	return path
}
