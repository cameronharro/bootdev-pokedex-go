package main

import (
	"fmt"
	"os"
)

func exitCallback() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
