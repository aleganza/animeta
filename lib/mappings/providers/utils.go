package mapping_providers

func (p MappingProvider) IsValid() bool {
	switch p {
	case
		MappingProviderAniDB,
		MappingProviderAniList,
		MappingProviderAnimePlanet,
		MappingProviderAnimeCountdown,
		MappingProviderAnimeNewsNetwork,
		MappingProviderAniSearch,
		MappingProviderIMDb,
		MappingProviderKitsu,
		MappingProviderLiveChart,
		MappingProviderMAL,
		MappingProviderSimkl,
		MappingProviderTMDB,
		MappingProviderTVDB:
		return true
	default:
		return false
	}
}

// func (p MappingProvider) hasSeasonAsId() bool {

// }
