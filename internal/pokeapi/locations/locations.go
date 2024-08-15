package locations

import (
	"errors"
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/client"
)

func GetNextLocationAreas(
	pagination *ApiPagination,
	c *client.Client,
) (*LocationsResp, error) {
	var url string
	if pagination.Next == nil {
		url = client.Base_url + "location-area?offset=0&limit=20"
	} else {
		url = *pagination.Next
	}

	nextPagination, apiResponse, err := getLocations(url, c)
	if err != nil {
		return nil, err
	}

	*pagination = *nextPagination
	return apiResponse, nil
}

func GetPreviousLocationAreas(
	pagination *ApiPagination,
	c *client.Client,
) (*LocationsResp, error) {
	if pagination.Previous == nil {
		return nil, errors.New("there are no previous locations")
	}

	nextPagination, apiResponse, err := getLocations(*pagination.Previous, c)
	if err != nil {
		return nil, err
	}

	*pagination = *nextPagination
	return apiResponse, nil
}

func ExploreLocation(
	locationName string,
	c *client.Client,
) (*LocationResp, error) {
	url := fmt.Sprintf("%slocation-area/%s", client.Base_url, locationName)

	json, ok := c.TryGetFromCache(url)
	if ok {
		locationResp, err := getLocationFromJson(json)
		if err != nil {
			return nil, err
		}
		return locationResp, nil
	}

	json, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	locationResp, err := getLocationFromJson(json)
	if err != nil {
		return nil, err
	}

	return locationResp, nil
}

func getLocations(
	reqUrl string,
	c *client.Client,
) (*ApiPagination, *LocationsResp, error) {
	json, ok := c.TryGetFromCache(reqUrl)
	if ok {
		nextPagination, apiResp, err := getLocationsFromJson(json)
		if err != nil {
			return nil, nil, err
		}
		return nextPagination, apiResp, nil
	}

	json, err := c.Get(reqUrl)
	if err != nil {
		return nil, nil, err
	}

	nextPagination, apiResp, err := getLocationsFromJson(json)
	if err != nil {
		return nil, nil, err
	}
	return nextPagination, apiResp, nil
}
