package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (DetailLocationArea, error) {
	url := baseURL + "/location-area/" + name

	//check the cache to see if an entry matches
	if val, ok := c.pokeCache.Get(url); ok {
		fmt.Println("cache hit:", url)
		detailStruct := DetailLocationArea{}
		err := json.Unmarshal(val, &detailStruct)
		if err != nil {
			return DetailLocationArea{}, err
		}
		return detailStruct, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DetailLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DetailLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return DetailLocationArea{}, err
	}
	//add to cache
	fmt.Println("cache miss, fetching:", url)
	c.pokeCache.Add(url, dat)

	detailStruct := DetailLocationArea{}
	err = json.Unmarshal(dat, &detailStruct)
	if err != nil {
		return DetailLocationArea{}, err
	}

	return detailStruct, nil
}
