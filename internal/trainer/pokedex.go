package trainer

import (
	"github.com/nelsonmarro/pokedexcli/internal/pokeapi/pokemon"
)

var pokedex = make(map[string]pokemon.Pokemon)

func GetPokemonList() []pokemon.Pokemon {
	var pokes []pokemon.Pokemon
	for _, poke := range pokedex {
		pokes = append(pokes, poke)
	}
	return pokes
}

func GetPokemonFromPokedex(name string) (*pokemon.Pokemon, bool) {
	poke, ok := pokedex[name]
	if !ok {
		return nil, false
	}

	return &poke, true
}

func SetPokemonInPokedex(pokemon *pokemon.Pokemon) {
	pokedex[pokemon.Name] = *pokemon
}
