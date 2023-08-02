package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	valor := soma(4, 4, 8, 8)
	expected := int64(24)
	if valor != expected {
		t.Errorf("Soma: resultado esperado %d, mas obteve %d", expected, valor)
	}
}

func TestSub(t *testing.T) {
	res := sub(8, 2, 1)
	expected := int64(5)
	if res != expected {
		t.Errorf("Subtração: resultado esperado %d, mas obteve %d", expected, res)
	}
}

func TestMult(t *testing.T) {
	val := mult(2, 3, 4, 5)
	expected := int64(120)
	if val != expected {
		t.Errorf("Multiplicação: resultado esperado %d, mas obteve %d", expected, val)
	}
}

func TestDiv(t *testing.T) {
	mx := div(10, 2)
	expected := 5.0
	if mx != expected {
		t.Errorf("Divisão: resultado esperado %f, mas obteve %f", expected, mx)
	}
}

func TestDivPorZero(t *testing.T) {
	mx := div(10, 0)
	if mx == mx {
		t.Error("Divisão por zero: deveria retornar NaN, mas obteve valor numérico")
	}
}
