package main

import (
	"fmt"
	"sort"
)

func algoritmo10(x []int) ([]int, error) {
	if len(x) == 0 {
		return nil, fmt.Errorf("The slice is empty")
	}
	sort.Ints(x)
	return x, nil
}
