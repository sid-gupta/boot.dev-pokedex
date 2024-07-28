package internal

import (
	"fmt"
)

func (p *Pokedex) help(args ...string) error {
	fmt.Println("Available commands: ")
	for k, v := range p.commandMap {
		fmt.Printf("%s: %s\n", k, v.Desc)
	}
	return nil
}
