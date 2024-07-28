package internal

import (
	"errors"
	"fmt"
	"math/rand"
)

func (p *Pokedex) catch(args ...string) error {
	if len(args) < 1 {
		return errors.New("must provide a pokemon name to catch it")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := p.client.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := rand.Int() % 200
	caught := chance >= pokemon.BaseExperience

	if caught {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the 'inspect' command")
		p.pokemonCaught[name] = pokemon
	} else {
		fmt.Printf("%s excaped!\n", name)
	}

	return nil
}
