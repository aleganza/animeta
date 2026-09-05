package anime_mappings

import (
	"testing"
)

func TestGetMappingsFromProviderId_InvalidProvider(t *testing.T) {
	_, err := GetMappingsFromProviderId("invalid", "1")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetMappingsFromAniListId(t *testing.T) {
	result, err := GetMappingsFromProviderId(
		MappingProviderAniList,
		"114745",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.AniListID != 114745 {
		t.Errorf("expected AniListID 114745, got %d", result.AniDBID)
	}
}
