package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

type areasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (areasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return areasResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areasResponse{}, err
	}
	
	var areas areasResponse
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return areasResponse{}, err
	}

	return areas, nil
}
