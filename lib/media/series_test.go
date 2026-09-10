package media

import (
	"encoding/json"
	"testing"
)

func TestGetSeriesFromTvdbId_InvalidId(t *testing.T) {

}

func TestGetSeriesFromTvdbId(t *testing.T) {
	result, err := FetchSeries("326109", "s")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

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
		if episode.TvdbId == 9220735 {
			testEpisodeTitle = episode.Title
		}
	}

	if testEpisodeTitle != "Gold" {
		t.Errorf(`episode title for episode id 9220735 should be "Gold" but is "%s"`, testEpisodeTitle)
	}
}
