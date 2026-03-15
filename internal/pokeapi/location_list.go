package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	//check the cache to see if an entry matches
	if val, ok := c.pokeCache.Get(url); ok {
		fmt.Println("cache hit:", url)
		myStruct := LocationAreasResponse{}
		err := json.Unmarshal(val, &myStruct)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return myStruct, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	//add to cache
	fmt.Println("cache miss, fetching:", url)
	c.pokeCache.Add(url, dat)

	myStruct := LocationAreasResponse{}
	err = json.Unmarshal(dat, &myStruct)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return myStruct, nil
}
