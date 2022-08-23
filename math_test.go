package main 

import "testing"


func TestSoma(t *testing.T) {

	total := soma(15, 17)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}