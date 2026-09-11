package anime_mappings

import (
	"encoding/json"
	"testing"
)

func TestGetMappingsFromProviderId_InvalidProvider(t *testing.T) {
	_, err := getMappingsFromProviderId("invalid", "1")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetTVMappingsFromAniListId(t *testing.T) {
	result, err := GetMappingsFromAniListId("114745")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

	if result.AniListID != 114745 {
		t.Errorf("expected AniListId 114745, got %d", result.AniListID)
	}
}

func TestGetMovieMappingsFromAniListId(t *testing.T) {
	result, err := GetMappingsFromAniListId("100643")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

	if result.AniListID != 100643 {
		t.Errorf("expected AniListId 100643, got %d", result.AniListID)
	}
}

func TestGetMappingsFromMALId(t *testing.T) {
	result, err := GetMappingsFromMALId("34599")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	t.Logf("Response:\n%s", jsonData)

	if result.MALID != 34599 {
		t.Errorf("expected MALId 34599, got %d", result.MALID)
	}
}
