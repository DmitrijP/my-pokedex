package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func RequestLocations(overwriteUrl string) LocationsResponse {
	url := baseURL + "location-area"
	if overwriteUrl != "" {
		url = overwriteUrl
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

	locations := LocationsResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}
	return locations
}
