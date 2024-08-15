package repl

import (
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/trainer"
)

func execInspectCmd(args CmdArgs) error {
	if len(args.Arguments) == 0 {
		return fmt.Errorf(
			"inspect command requires a pokemon name\nUsage: inspect <pokemon-name>",
		)
	}

	// consultar pokedex
	pokemon, ok := trainer.GetPokemonFromPokedex(args.Arguments[0])
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
