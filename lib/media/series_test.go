package media

import (
	"encoding/json"
	"testing"
)

// func TestGetSeriesFromTvdbId_InvalidId(t *testing.T) {

// }

func TestGetSeriesFromTvdbId(t *testing.T) {
	// season 1 of made in abyss
	result, err := FetchSeries(326109, 1)

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

	if testTitle != "Made in Abyss" {
		t.Errorf(`italian name should be "Made in Abyss" but is "%s"`, testTitle)
	}

	// test episodes
	testEpisodeTitle := ""

	for _, episode := range result.Episodes {
		if episode.TvdbId == 6180539 {
			for _, title := range episode.Titles {
				if title.Language == "eng" {
					testEpisodeTitle = title.Name
				}
			}
		}
	}

	if testEpisodeTitle != "Departure" {
		t.Errorf(`episode title for episode id 6180539 should be "Departure" but is "%s"`, testEpisodeTitle)
	}
}

func TestGetWholeSeriesFromTvdbId(t *testing.T) {
	// hunter x hunter (2011): a single anilist season split across tvdb seasons
	result, err := FetchSeries(252322, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.EpisodeCount != 148 {
		t.Errorf(`whole series should have 148 episodes but has %d`, result.EpisodeCount)
	}

	// numbering must be global (1..148) to match anidb/tenrai
	seen := make(map[int]bool)
	for _, episode := range result.Episodes {
		if seen[episode.Number] {
			t.Errorf(`duplicate episode number %d`, episode.Number)
		}
		seen[episode.Number] = true
	}

	if result.Episodes[0].Number != 1 || result.Episodes[147].Number != 148 {
		t.Errorf(
			`global numbering expected 1..148 but first=%d last=%d`,
			result.Episodes[0].Number,
			result.Episodes[147].Number,
		)
	}
}
