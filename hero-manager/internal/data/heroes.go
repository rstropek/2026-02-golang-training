package data

import (
	"encoding/json"
	"strings"
	"time"
)

type Hero struct {
	ID        int64     `json:"id"`
	FirstSeen time.Time `json:"firstSeen"`
	Name      string    `json:"name"`
	CanFly    bool      `json:"canFly"`
	RealName  string    `json:"realName,omitempty"`
	Abilities []string  `json:"-"`
	Version   int32     `json:"version"`
}

func (h Hero) MarshalJSON() ([]byte, error) {
	// Build a comma-separated string with abilities.
	abilities := strings.Join(h.Abilities, ", ")

	type HeroAlias Hero // Break the recursion by creating an alias type

	extendedHero := struct {
		HeroAlias
		Abilities string `json:"abilities"`
	}{
		HeroAlias: HeroAlias(h),
		Abilities: abilities,
	}

	return json.Marshal(extendedHero)
}
