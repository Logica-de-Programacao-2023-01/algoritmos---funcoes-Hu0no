package main

import "fmt"

func algoritmo15(x []int, y int) (bool, error) {
	if len(x) == 0 {
		return false, fmt.Errorf("Erro: Slice vazio")
	}
	i := 0
	z := false
	for i < len(x) {
		if x[i] == y {
			z = true
			break
		}
		i++
	}
	return z, nil
}
