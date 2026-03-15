package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + name

	//check the cache to see if an entry matches
	if val, ok := c.pokeCache.Get(url); ok {
		fmt.Println("cache hit:", url)
		pokedetailStruct := PokemonDetails{}
		err := json.Unmarshal(val, &pokedetailStruct)
		if err != nil {
			return PokemonDetails{}, err
		}
		return pokedetailStruct, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}
	//add to cache
	fmt.Println("cache miss, fetching:", url)
	c.pokeCache.Add(url, dat)

	pokedetailStruct := PokemonDetails{}
	err = json.Unmarshal(dat, &pokedetailStruct)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokedetailStruct, nil
}
