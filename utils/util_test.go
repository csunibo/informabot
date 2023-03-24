package utils

import (
	"testing"
)

func TestToKebabCase(t *testing.T) {
	// Test strings
	str1 := "Logica per l'informatica"
	str2 := "Informatica e Societa"
	str3 := "Sistemi Operativi"

	// Expected results
	exp1 := "logica-per-informatica"
	exp2 := "informatica-e-societa"
	exp3 := "sistemi-operativi"

	// Test
	res1 := ToKebabCase(str1)
	res2 := ToKebabCase(str2)
	res3 := ToKebabCase(str3)

	// Check results

	if res1 != exp1 {
		t.Errorf("Expected %s, got %s", exp1, res1)
	}

	if res2 != exp2 {
		t.Errorf("Expected %s, got %s", exp2, res2)
	}

	if res3 != exp3 {
		t.Errorf("Expected %s, got %s", exp3, res3)
	}
}

func TestAccents(t *testing.T) {
	str := "Informatica e Società"
	exp := "informatica-e-societa"

	str2 := "à è ì ò ù"
	exp2 := "a-e-i-o-u"

	res := ToKebabCase(str)
	res2 := ToKebabCase(str2)

	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}

	if res2 != exp2 {
		t.Errorf("Expected %s, got %s", exp2, res2)
	}
}
