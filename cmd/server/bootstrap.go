package main

import (
	"animeta/lib/core/env"
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
	"animeta/lib/media"
	"log"
)

func bootstrap() {
	env.Init()

	tvdbClient, err := tvdb.Authorize()
	if err != nil {
		log.Fatalf("TVDB authorize failed: %v", err)
	}

	media.NewClient(&tvdbClient)
}
