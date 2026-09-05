package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
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
