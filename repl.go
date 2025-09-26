package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func cleanUserInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	return filterSlice(words, func(s string) bool {
		return s != ""
	})
}

func promptAndGetInput(scanner *bufio.Scanner) []string {
	fmt.Print("Pokedex > ")
	scanner.Scan()
	text := scanner.Text()
	cleanedText := cleanUserInput(text)
	return cleanedText
}

func getFirstWord(words []string) string {
	if len(words) == 0 {
		fmt.Printf("No command provided\n\n")
		return "help"
	}
	return words[0]
}

func runCommandCallback(config *Config, command CliCommand, args []string) {
	err := command.callback(config, args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func runREPL(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		words := promptAndGetInput(scanner)

		firstWord := getFirstWord(words)

		command, ok := getCommandRegistry()[firstWord]
		if !ok {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}

		runCommandCallback(config, command, words[1:])
	}
}
