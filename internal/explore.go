package internal

import (
	"errors"
	"fmt"
)

func (p *Pokedex) explore(args ...string) error {
	if len(args) < 1 {
		return errors.New("provide a location to explore command")
	}
	//nextPage := p.currentMapPage + 1

	//fmt.Println("Mapping over page ", nextPage)

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)
	locationArea, err := p.client.GetLocationAreaPokemons(location)
	if err != nil {
		return err
	}

	if len(locationArea.PokemonEncounters) > 0 {
		fmt.Println("Found pokemon:")
		for _, pe := range locationArea.PokemonEncounters {
			fmt.Printf(" - %s\n", pe.Pokemon.Name)
		}
	} else {
		fmt.Println("No pokemon found")
	}

	//p.currentMapPage = nextPage
	return nil
}
