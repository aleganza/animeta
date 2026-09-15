package media

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGetMovieFromTvdbId(t *testing.T) {
	// made in abyss: wandering twilight
	result, err := FetchMovie(118959)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

	// test episode count

	if result.EpisodeCount != len(result.Episodes) {
		t.Errorf(`episode count should be %d but is %d`, len(result.Episodes), result.EpisodeCount)
	}

	// test titles

	testTitle := ""

	for _, title := range result.Titles {
		if title.Language == "ita" {
			testTitle = title.Name
		}
	}

	if testTitle != "Made in Abyss: Crepuscolo errante" {
		t.Errorf(`italian name should be "Made in Abyss: Crepuscolo errante" but is "%s"`, testTitle)
	}

	// test single episode

	if len(result.Episodes) != 1 {
		t.Fatalf("expected exactly 1 episode, got %d", len(result.Episodes))
	}

	episode := result.Episodes[0]

	if episode.Number != 1 {
		t.Errorf(`number should be 1 but is %d`, episode.Number)
	}

	episodeTitle := ""

	for _, title := range episode.Titles {
		if title.Language == "eng" {
			episodeTitle = title.Name
		}
	}

	if episodeTitle != "Made in Abyss: Wandering Twilight" {
		t.Errorf(`title should be "Made in Abyss: Wandering Twilight" but is "%s"`, episodeTitle)
	}

	if episode.Year != "2019" {
		t.Errorf(`year should be "2019" but is "%s"`, episode.Year)
	}

	if episode.Runtime != 105 {
		t.Errorf(`runtime should be 105 but is %d`, episode.Runtime)
	}

	if episode.Aired != "2019-01-18" {
		t.Errorf(`aired should be "2019-01-18" but is "%s"`, episode.Aired)
	}

	if !strings.Contains(episode.Overview, "Second compilation movie") {
		t.Errorf(`overview should contain "Second compilation movie" but is "%s"`, episode.Overview)
	}
}
