package main

import (
	"fmt"
	"bufio"
	"os"
)

func startRepl () {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		command := words[0]
		fmt.Println("Your command was: ", command)
	}
}