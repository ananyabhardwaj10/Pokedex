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
		cmd := words[0]
		command, ok := getCommands()[cmd]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println("err")
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}