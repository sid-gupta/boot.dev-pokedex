package internal

import (
	"errors"
	"fmt"
)

func (p *Pokedex) mapBack(args ...string) error {
	nextPage := p.currentMapPage - 1

	if nextPage < 0 {
		return errors.New("Already at first page")
	}

	fmt.Println("Mapping over page ", nextPage)

	locationAreas, err := p.client.GetLocationAreas(nextPage)
	if err != nil {
		return err
	}

	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}

	p.currentMapPage = nextPage
	return nil
}
