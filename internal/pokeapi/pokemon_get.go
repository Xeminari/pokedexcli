package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (RespPokemonDetails, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	if data, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemonDetails{}
		err := json.Unmarshal(data, &pokemonResp)
		return pokemonResp, err
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespPokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonDetails{}, err
	}
	pokemonResp := RespPokemonDetails{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemonDetails{}, err
	}
	c.cache.Add(url, dat)
	return pokemonResp, nil
}
