package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	data, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	result := Pokemon{}
	if err = json.Unmarshal(data, &result); err != nil {
		return Pokemon{}, err
	}
	return result, nil
}
