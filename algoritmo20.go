package main

import "fmt"

func algoritmo20(y []string) ([]string, error) {
	if len(y) == 0 {
		return nil, fmt.Errorf("Erro: Slice Vazio")
	}
	i := 0
	z := []string{}
	for i < len(y) {
		if len(y[i]) > 5 {
			z = append(z, y[i])
		}
		i++
	}
	return z, nil
}
