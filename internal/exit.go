package internal

import (
	"fmt"
	"os"
)

func (p *Pokedex) exit(args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(1)
	return nil
}
