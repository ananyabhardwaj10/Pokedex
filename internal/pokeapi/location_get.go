package pokeapi
import(
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func (c *Client) GetLocation(name string) (LocationArea, error) {
	var loc LocationArea
	url := baseURL + "/location-area/" + name 
	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &loc); err != nil {
			return LocationArea{}, err 
		}
		return loc, nil 
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err 
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err 
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("Bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &loc)
	if err != nil {
		return LocationArea{}, err 
	}
	return loc, nil 
}