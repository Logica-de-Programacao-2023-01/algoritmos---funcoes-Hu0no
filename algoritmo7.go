package main

import "fmt"

func soma5(x []int) []int {
	for y := range x {
		x[y] += 5
	}
	return x
}
func algoritmo7(x []int, f func([]int) []int) ([]int, error) {
	if len(x) == 0 {
		return nil, fmt.Errorf("The slice is empty")
	}
	resultado := soma5(x)
	return resultado, nil
}
