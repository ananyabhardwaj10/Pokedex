package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/ananyabhardwaj10/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client 
	nextLocationUrl *string 
	previousLocationUrl *string
}

func startRepl (c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		cmd := words[0]
		args := words[1:]
		command, ok := getCommands()[cmd]
		if ok {
			err := command.callback(c, args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}