package pokeapi 
import(
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var poke Pokemon
	
	url := baseURL + "/pokemon/" + name 
	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &poke)
		if err != nil {
			return Pokemon{}, err 
		}
		return poke, nil 
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err 
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err 
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, errors.New("Bad Status Code")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err 
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &poke)
	if err != nil {
		return Pokemon{}, err 
	}

	return poke, nil
}