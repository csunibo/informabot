package utils

import (
	"fmt"
	"testing"
)

func ExampleToKebabCase() {
	fmt.Println(ToKebabCase("Hello World"))
	// Output: hello-world
}

func TestToKebabCase(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"Standard",
			"Logica per l'informatica",
			"logica-per-informatica",
		},
		{
			"Standard",
			"Informatica e Societa",
			"informatica-e-societa",
		},
		{
			"Standard",
			"Sistemi Operativi",
			"sistemi-operativi",
		},
		{
			"Accents",
			"Informatica e Società",
			"informatica-e-societa",
		},
		{
			"Accents",
			"à è ì ò ù",
			"a-e-i-o-u",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToKebabCase(tt.args); got != tt.want {
				t.Errorf("ToKebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
