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
		z := 1
		divisores := 0
			for z <= x {
				if x%z == 0 {
					divisores += 1
				}
				z++
			}

		}
	}
}
func main() {
	x := 5
}
