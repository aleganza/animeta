package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	root, err := findRoot()
	if err != nil {
		panic(err)
	}

	if err := godotenv.Load(filepath.Join(root, ".env")); err != nil {
		panic(err)
	}
}

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root (go.mod) not found")
		}

		dir = parent
	}
}

func GetVar(name string) (string, error) {
	value, exists := os.LookupEnv(name)

	if !exists {
		return "", fmt.Errorf(`environment variable "%s" not found`, name)
	}

	return value, nil
}

func GetIntVar(name string) (int, error) {
	value, err := GetVar(name)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(`environment variable "%s" is not a valid integer: %w`, name, err)
	}

	return result, nil
}