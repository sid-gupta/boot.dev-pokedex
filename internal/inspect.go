package internal

import (
	"errors"
	"fmt"
)

func (p *Pokedex) inspect(args ...string) error {
	if len(args) < 1 {
		return errors.New("pokemon name not provided")
	}

	name := args[0]

	pokemon, found := p.pokemonCaught[name]
	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
