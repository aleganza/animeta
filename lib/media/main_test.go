package media

import (
	"log"
	"os"
	"testing"

	"animeta/lib/core/env"
	"animeta/lib/data_sources/anime_meta/providers/anidb"
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
)

func TestMain(m *testing.M) {
	env.Init()

	tvdbClient, err := tvdb.Authorize()
	if err != nil {
		log.Fatalf("TVDB authorize failed: %v", err)
	}

	anidbClient, err := anidb.NewClient()
	if err != nil {
		log.Printf("AniDB client not configured: %v", err)
		anidbClient = anidb.Client{}
	}

	NewClient(&tvdbClient, &anidbClient)

	code := m.Run()

	os.Exit(code)
}
