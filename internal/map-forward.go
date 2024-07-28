package internal

import "fmt"

func (p *Pokedex) mapForward(args ...string) error {
	nextPage := p.currentMapPage + 1

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
