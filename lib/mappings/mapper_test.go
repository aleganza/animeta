package mappings

import (
	mapping_providers "animeta/lib/mappings/providers"
	"testing"
)

func TestFetchMappingsFromProviderId_InvalidProvider(t *testing.T) {
	_, err := FetchMappingsFromProviderId("invalid", "1")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestFetchMappingsFromProviderId(t *testing.T) {
	result, err := FetchMappingsFromProviderId(
		mapping_providers.MappingProviderAniDB,
		"1",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.AniDBID != 1 {
		t.Errorf("expected AniDBID 1, got %d", result.AniDBID)
	}
}