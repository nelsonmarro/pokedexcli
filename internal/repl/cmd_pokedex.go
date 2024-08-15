package repl

import (
	"fmt"

	"github.com/nelsonmarro/pokedexcli/internal/trainer"
)

func execPokedexCmd(args CmdArgs) error {
	pokemons := trainer.GetPokemonList()
	if len(pokemons) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some pokemons!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, poke := range pokemons {
		fmt.Printf("  - %s\n", poke.Name)
	}

	return nil
}
