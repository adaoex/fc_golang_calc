package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 10)

	if total != 20 {
		t.Errorf("Resultado inválido! Resuldado: %d. Esperado: %d", total, 20)
	}
}
