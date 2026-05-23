package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func GetAreaPokemon(area string) (areaPokemon, error) {
	fullURL := fmt.Sprintf("%slocation-area/%s/", baseURL, area)
	res, err := http.Get(fullURL)
	if err != nil {
		return areaPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areaPokemon{}, err
	}

	var pokemon areaPokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return areaPokemon{}, err
	}

	return pokemon, nil
}
