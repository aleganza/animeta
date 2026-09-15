package media

import "testing"

func TestEnrichWithTenrai(t *testing.T) {
	if client.tenrai == nil {
		t.Skip("tenrai client not configured")
	}

	mediaData := Media{
		Episodes: []Episode{
			{TvdbId: 1, Number: 9},
			{Number: 19},
			{Number: 25},
		},
	}

	// fairy tail (mal id 6702) has filler episodes, e.g. 9 and 19
	enrichWithTenrai(&mediaData, 6702)

	if !mediaData.Episodes[0].Filler {
		t.Errorf(`episode 9 should be marked as filler`)
	}

	if !mediaData.Episodes[1].Filler {
		t.Errorf(`episode 19 should be marked as filler`)
	}

	if mediaData.Episodes[2].Filler {
		t.Errorf(`episode 25 should not be marked as filler`)
	}

	if mediaData.Episodes[0].Recap {
		t.Errorf(`episode 9 should not be marked as recap`)
	}
}
