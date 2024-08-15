package main

import (
	"time"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/client"
	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/locations"
	"github.com/nelsonmarro/pokedexcli/internal/repl"
)

func main() {
	pokeClient := client.NewClient(5*time.Second, time.Minute*5)
	cmdArgs := repl.CmdArgs{
		PokeapiClient: &pokeClient,
		Pagination:    &locations.ApiPagination{},
	}

	repl.StartRepl(cmdArgs)
}
