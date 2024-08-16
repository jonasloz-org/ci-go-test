package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := sub(15, 15)
	if total != 0 {
		t.Errorf("Resultado da subtração inválido: Resultado %d. Esperado: %d", total, 0)
	}
}

func TestTimes(t *testing.T) {
	total := times(15, 15)
	if total != 225 {
		t.Errorf("Resultado da multiplicação inválido: Resultado %d. Esperado: %d", total, 225)
	}
}

func TestSumX(t *testing.T) {
	total := sumX(15, 15)
	if total != 45 {
		t.Errorf("Resultado da sumX inválido: Resultado %d. Esperado: %d", total, 45)
	}
}
