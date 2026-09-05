package media

import (
	"strings"
)

func (p Provider) IsValid() bool {
	switch p {
	case
		ProviderAniList,
		ProviderMAL:
		return true
	default:
		return false
	}
}

func GetProvidersPretty() string {
	providers := make([]string, len(Providers))

	for i, provider := range Providers {
		providers[i] = string(provider)
	}

	return strings.Join(providers, ", ")
}
