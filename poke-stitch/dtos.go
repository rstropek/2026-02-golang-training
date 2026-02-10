package main

// pokemon matches only the PokeAPI fields this sample uses.
// Go's JSON decoder ignores unknown fields by default.
type Pokemon struct {
	Sprites PokemonSprites `json:"sprites"`
}

// pokemonSprites mirrors selected sprite URLs from the API payload.
// Struct tags map Go field names to snake_case JSON keys.
type PokemonSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

func (p Pokemon) SpriteURLs() []string {
	all := []string{
		p.Sprites.BackDefault,
		p.Sprites.BackFemale,
		p.Sprites.BackShiny,
		p.Sprites.BackShinyFemale,
		p.Sprites.FrontDefault,
		p.Sprites.FrontFemale,
		p.Sprites.FrontShiny,
		p.Sprites.FrontShinyFemale,
	}

	filtered := make([]string, 0, len(all))
	for _, url := range all {
		if url == "" {
			continue
		}
		filtered = append(filtered, url)
	}

	return filtered
}
