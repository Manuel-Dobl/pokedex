package pokeapi

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type DetailLocationArea struct {
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonDetails struct {
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Name           string  `json:"name"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

// this has the number on BaseStat like "45"
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

// this is the name of the stat itself like "attack"
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// this just has a slot and then the actual types
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// name has tghe actual type like "flying"
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
