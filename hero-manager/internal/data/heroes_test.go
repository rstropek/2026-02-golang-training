package data

import (
	"encoding/json"
	"testing"
	"time"
)

func TestHero_MarshalJSON_Abilities(t *testing.T) {
	// Arrange
	hero := Hero{
		ID:        1,
		FirstSeen: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		Name:      "Superman",
		CanFly:    true,
		RealName:  "Clark Kent",
		Abilities: []string{"Flight", "Super Strength", "Heat Vision"},
		Version:   1,
	}

	// Act
	jsonData, err := json.Marshal(hero)
	if err != nil {
		t.Fatalf("Failed to marshal hero: %v", err)
	}

	// Assert
	var result map[string]any
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	abilities, ok := result["abilities"].(string)
	if !ok {
		t.Errorf("Expected abilities to be a string, got %T", result["abilities"])
	}

	expected := "Flight, Super Strength, Heat Vision"
	if abilities != expected {
		t.Errorf("Expected abilities to be %q, got %q", expected, abilities)
	}
}
