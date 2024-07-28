package main

import (
	"example.com/pokedex/internal"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	dex := internal.NewPokedex()
	repl := PokedexRepl{pokedex: dex}
	repl.Start()
}
