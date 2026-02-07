package main
import (
	"time"
	"github.com/ananyabhardwaj10/Pokedex/internal/pokeapi"
)
func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}