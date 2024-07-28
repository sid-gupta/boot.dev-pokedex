package pokeapi

import (
	"encoding/json"
	"example.com/pokedex/cache"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PokeapiClient struct {
	PageSize int
	BaseUrl  string
	Cache    cache.Cache[[]byte]
}

type ParsedResponse[T any] struct {
	Count    int
	Next     string
	Previous string
	Results  T
}

func (client PokeapiClient) prepareUrl(path string, page int) string {
	offset := client.PageSize * page
	return fmt.Sprintf("%s?offset=%d&limit=%d", path, offset, client.PageSize)
}

func (client PokeapiClient) get(path string) ([]byte, error) {
	time.Sleep(1 * time.Second)
	url := client.BaseUrl + path
	res, err := http.Get(url)

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%s: Status: %d. Response: %s", path, res.StatusCode, body)
	}

	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseResponse[T any](raw []byte) (T, error) {
	var resp T
	err := json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
