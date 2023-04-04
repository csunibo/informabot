package commands

import (
	"testing"
)

func TestWeekend(t *testing.T) {
	url := "https://corsi.unibo.it/laurea/informatica/orario-lezioni/@@orario_reale_json?anno=1&start=2023-03-11&end=2023-03-11"
	result := GetTimeTable(url)

	if result != "" {
		t.Errorf("Expected empty string in weekend, got %s", result)
	}
}
