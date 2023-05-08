package main

import "fmt"

func algoritmo14(x []int, y []int) ([]int, error) {
	if len(x) == 0 || len(y) == 0 {
		return nil, fmt.Errorf("Erro: Slice vazio")
	}
	i := 0
	for i < len(y) {
		x = append(x, y[i])
		i++
	}
	return x, nil
}
