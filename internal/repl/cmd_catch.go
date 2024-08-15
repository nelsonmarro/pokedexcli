package repl

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/pokemon"
	"github.com/nelsonmarro/pokedexcli/internal/trainer"
)

func execCatchCmd(args CmdArgs) error {
	if len(args.Arguments) == 0 {
		return fmt.Errorf(
			"catch command requiers a pokemon name\nUsage: catch <pokemon-name>",
		)
	}

	pokemon, err := pokemon.CatchPokemon(args.Arguments[0], args.PokeapiClient)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			fmt.Println(
				"Pokemon not found. Ensure you are using the correct name.",
			)
			return nil
		}
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catched := canCatch(pokemon.BaseExperience)
	if !catched {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	trainer.SetPokemonInPokedex(pokemon)
	return nil
}

func canCatch(baseExp int64) bool {
	// Crear una fuente basada en el tiempo actual
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)

	// Crear un nuevo generador de números aleatorios
	rng := rand.New(source)

	max := int(baseExp)
	// Generar números aleatorios
	catchResult := rng.Intn(max)

	return catchResult < 45
}
