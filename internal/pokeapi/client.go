package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
}

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	myStruct := LocationAreasResponse{}
	err = json.Unmarshal(body, &myStruct)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return myStruct, nil
}
