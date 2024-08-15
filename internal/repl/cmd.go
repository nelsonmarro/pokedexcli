package repl

import (
	"errors"
	"strings"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/client"
	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/locations"
)

type CmdArgs struct {
	Pagination    *locations.ApiPagination
	PokeapiClient *client.Client
	Arguments     []string
}

type cmd struct {
	Action      func(CmdArgs) error
	Name        string
	Description string
}

func GetCmds() map[string]cmd {
	return map[string]cmd{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Action:      execHelpCmd,
		},
		"map": {
			Name:        "map",
			Description: "Get the next 20 location areas",
			Action:      execMapCmd,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous 20 location areas",
			Action:      execMapbCmd,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area. Usage: explore <location-area-name>",
			Action:      execExploreCmd,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon by it's name. Usage: catch <pokemon-name>",
			Action:      execCatchCmd,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Print the details of a pokemon. Usage: inspect <pokemon-name>",
			Action:      execInspectCmd,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Print the list of pokemons in the pokedex",
			Action:      execPokedexCmd,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Action:      execExitCmd,
		},
	}
}

func GetCmdAction(name string) (func(CmdArgs) error, error) {
	for _, cmd := range GetCmds() {
		if strings.ToLower(cmd.Name) == name {
			return cmd.Action, nil
		}
	}

	return nil, errors.New("command not found!. Please try again")
}
