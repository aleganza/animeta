package anime_mappings

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

// these providers' media require additional information/parsing, e.g. season id for tvdb 
func (p MappingProvider) isMappingsRetrievalHandled() bool {
	switch p {
	case
		MappingProviderIMDb,
		MappingProviderTMDB,
		MappingProviderTVDB:
		return false
	default:
		return true
	}
}

// func (p MappingProvider) hasSeasonAsId() bool {

// }
