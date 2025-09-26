package pokeapi

import (
	"bootdev-pokedex-go/internal/pokecache"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(requestTimeout, reapInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: requestTimeout,
		},
		cache: *pokecache.NewCache(reapInterval),
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	if data, ok := c.cache.Get(url); ok {
		return data, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Error: response status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	c.cache.Add(url, data)
	return data, nil
}
