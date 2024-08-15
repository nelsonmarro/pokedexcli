package locations

import (
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/jsonutils"
)

func getLocationsFromJson(
	data []byte,
) (*ApiPagination, *LocationsResp, error) {
	// Deserializar el json
	apiResponse, err := jsonutils.DeserializeJson[LocationsResp](data)
	if err != nil {
		return nil, nil, fmt.Errorf("error deserializing json: %w", err)
	}
	nextPagination, err := jsonutils.DeserializeJson[ApiPagination](data)
	if err != nil {
		return nil, nil, fmt.Errorf("error deserializing json: %w", err)
	}

	return nextPagination, apiResponse, nil
}

func getLocationFromJson(data []byte) (*LocationResp, error) {
	locationRes, err := jsonutils.DeserializeJson[LocationResp](data)
	if err != nil {
		return nil, fmt.Errorf("error deserializing json: %w", err)
	}

	return locationRes, nil
}
