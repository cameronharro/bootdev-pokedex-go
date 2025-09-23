package main

import "net/url"

type Config struct {
	Next     url.URL
	Previous url.URL
}

func getOriginalConfig() (*Config, error) {
	originalURL, err := url.Parse("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return &Config{}, err
	}
	config := Config{
		Next:     *originalURL,
		Previous: *originalURL,
	}
	return &config, nil
}
