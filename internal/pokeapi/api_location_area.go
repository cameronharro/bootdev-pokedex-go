package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocationAreas(pageURL *string) (BulkLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.Get(url)
	if err != nil {
		return BulkLocationArea{}, err
	}

	result := BulkLocationArea{}
	if err = json.Unmarshal(data, &result); err != nil {
		return BulkLocationArea{}, err
	}

	return result, nil
}

func (c *Client) ExploreLocationArea(locationArea string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationArea

	data, err := c.Get(url)
	if err != nil {
		return LocationArea{}, err
	}

	result := LocationArea{}
	if err = json.Unmarshal(data, &result); err != nil {
		return LocationArea{}, err
	}
	return result, nil
}
