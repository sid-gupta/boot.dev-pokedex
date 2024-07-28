package main

import (
	"bufio"
	"example.com/pokedex/internal"
	"fmt"
	"log"
	"os"
	"strings"
)

type PokedexRepl struct {
	pokedex internal.Pokedex
}

func (repl PokedexRepl) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal("failed to read input", err.Error())
		}

		input := scanner.Text()
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "" {
			continue
		}

		repl.pokedex.HandleCommand(input)
	}
}
