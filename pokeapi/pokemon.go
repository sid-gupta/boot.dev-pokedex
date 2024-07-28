package pokeapi

import "fmt"

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Stat struct {
	Name string `json:"name"`
}

type Type struct {
	Name string `json:"name"`
}

func (client PokeapiClient) GetPokemon(name string) (Pokemon, error) {
	path := fmt.Sprintf("/pokemon/%s", name)

	var p Pokemon
	body, found := client.Cache.Get(path)

	if !found {
		res, err := client.get(path)
		if err != nil {
			return p, err
		}

		body = res
		client.Cache.Add(path, body)
	}

	p, err := parseResponse[Pokemon](body)
	if err != nil {
		return p, err
	}

	return p, nil
}
