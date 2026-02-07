package pokeapi

type Pokemon struct {
    BaseExperience int    `json:"base_experience"`
    Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats []struct {
    	BaseStat int `json:"base_stat"`
    	Stat     struct {
        Name string `json:"name"`
    	} `json:"stat"`
	}
	Types []struct {
    	Type struct {
        Name string `json:"name"`
    	} `json:"type"`
	}

}