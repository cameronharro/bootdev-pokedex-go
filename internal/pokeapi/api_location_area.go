package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetLocationAreas(pageURL *string) (BulkLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return BulkLocationArea{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return BulkLocationArea{}, err
	}

	result := BulkLocationArea{}
	if err = json.Unmarshal(data, &result); err != nil {
		return BulkLocationArea{}, err
	}

	return result, nil
}
