package internal

import (
	"example.com/pokedex/cache"
	"example.com/pokedex/pokeapi"
	"fmt"
	"strings"
	"time"
)

type Pokedex struct {
	commandMap     map[string]Command
	currentMapPage int
	client         pokeapi.PokeapiClient
}

func (p *Pokedex) HandleCommand(input string) {
	parts := strings.Fields(input)

	commandName := parts[0]
	command, found := p.commandMap[commandName]
	if !found {
		fmt.Printf("Unknown command '%s'. Try running 'help'\n", commandName)
		return
	}

	err := command.Handler(parts[1:]...)
	if err != nil {
		fmt.Printf("Error handling '%s'. Error: %v\n", commandName, err)
	}
}

func NewPokedex() Pokedex {
	expiringCache := cache.NewExpiringCache[[]byte](time.Second * 10)
	pokedex := Pokedex{
		currentMapPage: -1,
		client: pokeapi.PokeapiClient{
			BaseUrl:  "https://pokeapi.co/api/v2",
			PageSize: 2,
			Cache:    expiringCache,
		},
	}
	pokedex.commandMap = GetPokedexCommands(&pokedex)
	return pokedex
}
