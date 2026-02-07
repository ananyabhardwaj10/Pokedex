package main
import(
	"fmt"
	"os"
	"errors"
	"math/rand"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error 
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

	"catch": {
		name: "catch",
		description: "helps catch a pokemon",
		callback: commandCatch,
	},

	"inspect": {
		name: "inspect",
		description: "displays stats of the pokemon",
		callback: commandInspect,
	},

	}
}

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config, args ...string) error {
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

func commandMapBack(c *config, args ...string) error {
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

func commandExplore(c *config, args ...string) error {
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

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please provide a Pokemon name.")
	}

	pokeName := args[0]

	pokemon, err := c.pokeapiClient.GetPokemon(pokeName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	max := pokemon.BaseExperience
	if max < 1 {
		max = 1
	}
	roll := rand.Intn(max)

	if roll > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil 
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	c.caughtPokemon[pokemon.Name] = pokemon 

	return nil
}

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Provide valid Pokemon Name.")
	}
	
	value, exists := c.caughtPokemon[args[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Height: %d\n", value.Height)
		fmt.Printf("Weight: %d\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
    		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range value.Types {
    		fmt.Println("  -", typeInfo.Type.Name)
		}

	}
	return nil
}