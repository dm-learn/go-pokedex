package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokemon(pokemon string) (Pokemon, error) {
	fullURL := fmt.Sprintf("%spokemon/%s/", baseURL, pokemon)
	res, err := http.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemonInfo Pokemon
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return Pokemon{}, nil
	}
	return pokemonInfo, nil
}
