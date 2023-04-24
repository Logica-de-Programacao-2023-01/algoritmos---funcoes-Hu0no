package main

import "fmt"

func media(x []int) (float64, error) {
	soma := 0.0
	if len(x) == 0 {
		return 0, fmt.Errorf("Cannot calculate the media of 0 numbers")
	}
	for y := range x {
		soma += float64(x[y])
		y++
	}
	return soma / float64(len(x)), nil
}
