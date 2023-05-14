package main

import "fmt"

func algoritmo19(x int) ([]int, error) {
	if x < 0 {
		return nil, fmt.Errorf("Erro: Número inválido")
	}
	y := []int{}
	if x == 1 || x == 0 {
		return y, nil
	} else {
		z := 2
		d := 1
		divisores := 0
		for z <= x {
			divisores = 0
			d = 1
			for d <= z {
				if z%d == 0 {
					divisores += 1
				}
				d++
			}
			if divisores == 2 {
				y = append(y, z)
			}
			z++
		}
		return y, nil
	}
}
