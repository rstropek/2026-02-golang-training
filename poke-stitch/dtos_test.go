package main

import (
	"reflect"
	"testing"
)

func TestPokemonSpriteURLs(t *testing.T) {
	t.Run("filters empty", func(t *testing.T) {
		p := Pokemon{
			Sprites: PokemonSprites{
				BackDefault:  "back-default",
				BackFemale:   "",
				BackShiny:    "back-shiny",
				FrontDefault: "front-default",
			},
		}

		want := []string{
			"back-default",
			"back-shiny",
			"front-default",
		}

		got := p.SpriteURLs()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("SpriteURLs() mismatch: got %v, want %v", got, want)
		}
	})

	t.Run("all empty", func(t *testing.T) {
		p := Pokemon{}

		got := p.SpriteURLs()
		if len(got) != 0 {
			t.Fatalf("SpriteURLs() expected empty slice, got %v", got)
		}
	})

	t.Run("preserves order", func(t *testing.T) {
		p := Pokemon{
			Sprites: PokemonSprites{
				BackDefault:      "back-default",
				BackFemale:       "back-female",
				BackShiny:        "back-shiny",
				BackShinyFemale:  "back-shiny-female",
				FrontDefault:     "front-default",
				FrontFemale:      "front-female",
				FrontShiny:       "front-shiny",
				FrontShinyFemale: "front-shiny-female",
			},
		}

		want := []string{
			"back-default",
			"back-female",
			"back-shiny",
			"back-shiny-female",
			"front-default",
			"front-female",
			"front-shiny",
			"front-shiny-female",
		}

		got := p.SpriteURLs()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("SpriteURLs() mismatch: got %v, want %v", got, want)
		}
	})
}
