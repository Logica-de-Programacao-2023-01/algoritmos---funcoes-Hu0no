package main

import "fmt"

func soma5(x int) int {
	x += 5
	return x
}
func algoritmo18(y []int, soma5 func(int) int) (int, error) {
	if len(y) == 0 {
		return 0, fmt.Errorf("Erro: Slice vazio")
	}
	i := 0
	soma := 0
	for i < len(y) {
		y[i] = soma5(y[i])
		soma += y[i]
		i++
	}
	return soma, nil
}
