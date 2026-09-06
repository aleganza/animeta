package media

import (
	"os"
	"testing"

	"animeta/lib/core/env"
)

func TestMain(m *testing.M) {
	env.Init()

	code := m.Run()

	os.Exit(code)
}
