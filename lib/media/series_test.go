package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
	"encoding/json"
	"testing"
)

func TestGetSeriesFromTvdbId(t *testing.T) {
	tvdbClient, err := tvdb.Authorize()
	if err != nil {
		t.Fatalf("TVDB authorize failed: %v", err)
		return
	}

	NewClient(&tvdbClient)

	result, err := FetchSeries("326109", "s")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)
}
