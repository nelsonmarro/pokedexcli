package repl

import (
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/locations"
)

func execExploreCmd(args CmdArgs) error {
	if len(args.Arguments) == 0 {
		return fmt.Errorf(
			"explore command requires a location area name\nUsage: explore <location-area-name>",
		)
	}

	locationResp, err := locations.ExploreLocation(
		args.Arguments[0],
		args.PokeapiClient,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemons:")
	for _, pokeEncounters := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokeEncounters.Pokemon.Name)
	}

	return nil
}
