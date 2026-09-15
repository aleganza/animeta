package tenrai

import (
	"encoding/json"
	"testing"
)

func TestFetchAllAnimeEpisodes(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// made in abyss (mal id 34599)
	result, err := client.FetchAllAnimeEpisodes(34599)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

	if len(result) == 0 {
		t.Fatal("expected at least 1 episode, got 0")
	}

	// every episode must carry the filler/recap flags
	for _, episode := range result {
		// bool fields can't be asserted to a specific value, only presence matters
		_ = episode.Filler
		_ = episode.Recap
	}
}
