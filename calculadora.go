package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Bem-vindo à calculadora!")
}

func soma(i ...int64) int64 {
	var resultado int64
	if len(i) == 0 {
		return 0
	}
	for _, k := range i {
		resultado += k
	}
	return resultado
}

func sub(i ...int64) int64 {
	if len(i) == 0 {
		return 0
	}
	resultado := i[0]
	for _, k := range i[1:] {
		resultado -= k
	}
	return resultado
}

func mult(i ...int64) int64 {
	var resultado int64 = 1
	if len(i) == 0 {
		return 0
	}
	for _, k := range i {
		resultado *= k
	}
	return resultado
}

func div(i ...float64) float64 {
	if len(i) == 0 {
		return math.NaN()
	}
	k := len(i)
	for j := 0; j < k; j++ {
		if i[j] == 0 {
			return math.NaN()
		}
	}
	resultado := i[0]
	for _, num := range i[1:] {
		resultado /= num
	}
	return resultado
}
