package main

import "fmt"

func algoritmo11(x int) (bool, error) {
	if x < 0 {
		return false, fmt.Errorf("Cannot use negative numbers")
	}
	if x == 0 || x == 1 {
		return false, nil
	} else {
		z := 1
		qtd_divisores := 0
		for z <= x {
			if x%z == 0 {
				qtd_divisores += 1
			}
			z += 1
		}
		if qtd_divisores == 2 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
