package assistance

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rootDir, err := findRootDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func findRootDir(dir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("go.mod not found")
		}
		dir = parent
	}
}
