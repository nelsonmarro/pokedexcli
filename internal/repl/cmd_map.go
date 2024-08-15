package repl

import (
	"errors"
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/locations"
)

func execMapCmd(
	args CmdArgs,
) error {
	if args.Pagination == nil {
		return errors.New("can't execute map command without pagination")
	}
	// make request to get next 20 reqResponse
	reqResponse, err := locations.GetNextLocationAreas(
		args.Pagination,
		args.PokeapiClient,
	)
	if err != nil {
		return err
	}

	for _, location := range reqResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
