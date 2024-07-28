package pokeapi

import "fmt"

type LocationArea struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

func (client PokeapiClient) GetLocationAreas(page int) ([]LocationArea, error) {
	path := client.prepareUrl("/location-area", page)

	body, found := client.Cache.Get(path)
	if !found {
		res, err := client.get(path)
		if err != nil {
			return nil, err
		}
		body = res
		client.Cache.Add(path, body)
	}

	out, err := parseResponse[ParsedResponse[[]LocationArea]](body)
	if err != nil {
		return nil, err
	}

	return out.Results, nil
}

func (client PokeapiClient) GetLocationAreaPokemons(locationArea string) (LocationArea, error) {
	path := fmt.Sprintf("/location-area/%s", locationArea)
	path = client.prepareUrl(path, 0)

	var la LocationArea
	body, found := client.Cache.Get(path)

	if !found {
		res, err := client.get(path)
		if err != nil {
			return la, err
		}
		body = res
		client.Cache.Add(path, body)
	}

	la, err := parseResponse[LocationArea](body)
	if err != nil {
		return la, err
	}

	return la, nil
}
