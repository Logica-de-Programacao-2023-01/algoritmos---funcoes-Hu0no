package main

import (
	"sort"
)

func segundomaior(x []int) int {
	sort.Ints(x)
	return x[len(x)-2]
}
