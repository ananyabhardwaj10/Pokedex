package main
import (
	"time"
	"github.com/ananyabhardwaj10/Pokedex/internal/pokeapi"
)
func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}