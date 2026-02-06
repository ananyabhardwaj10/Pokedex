package main
import(
	"fmt"
	"os"
	"errors"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error 
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},

	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	
	"map": {
		name: "map", 
		description: "displays next page of locations",
		callback: commandMap,
	},

	"mapback": {
		name: "mapBack",
		description: "displays previous page of locations",
		callback: commandMapBack,
	},

	"explore": {
		name: "explore",
		description: "displays all the pokemons present in a particular location",
		callback: commandExplore,
	},

	}
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config, args []string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationUrl)
	if err != nil {
		return err 
	}
	c.nextLocationUrl = locationsResp.Next 
	c.previousLocationUrl = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(c *config, args []string) error {
	if c.previousLocationUrl == nil {
		return errors.New("You are on the first page")
	}

	locationsResp, err := c.pokeapiClient.ListLocations(c.previousLocationUrl)
	if err != nil {
		return err 
	}

	c.nextLocationUrl = locationsResp.Next 
	c.previousLocationUrl = locationsResp.Previous 

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(c *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide a location name.")
	}
	loc_name := args[0]

	location, err := c.pokeapiClient.GetLocation(loc_name)
	if err != nil {
		return err 
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}