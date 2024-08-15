package locations

type ApiPagination struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type LocationsResp struct {
	Results []LocationAreas `json:"results"`
	Count   int64           `json:"count"`
}

type LocationAreas struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResp struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
	ID                int64              `json:"id"`
}

type PokemonEncounter struct {
	Pokemon PokemonInLocation `json:"pokemon"`
}

type PokemonInLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
