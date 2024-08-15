package pokemon

import (
	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/client"
	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/jsonutils"
)

func CatchPokemon(pokeName string, c *client.Client) (*Pokemon, error) {
	url := client.Base_url + "pokemon/" + pokeName

	json, ok := c.TryGetFromCache(url)
	if ok {
		pokemon, err := jsonutils.DeserializeJson[Pokemon](json)
		if err != nil {
			return nil, err
		}
		return pokemon, nil
	}

	json, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	pokemon, err := jsonutils.DeserializeJson[Pokemon](json)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
