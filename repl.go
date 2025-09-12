package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	words = filterSlice(words, func(s string) bool {
		return s != ""
	})
	return words
}

func filterSlice[T any](s []T, f func(T) bool) []T {
	n := 0
	for _, val := range s {
		if f(val) {
			s[n] = val
			n++
		}
	}
	return s[:n]
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getRegistry()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		firstWord := cleanedText[0]
		command, ok := registry[firstWord]
		if !ok {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}
