package internal

import "fmt"

func (p *Pokedex) pokedex(args ...string) error {
	if len(p.pokemonCaught) == 0 {
		fmt.Println("No pokemon caught by you")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k, _ := range p.pokemonCaught {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
