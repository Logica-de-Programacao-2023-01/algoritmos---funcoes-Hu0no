package main

func algoritmo5(x []int, y int) int {
	var i int = 0
	for i < len(x) {
		if x[i] == y {
			return i
		}
		i++
	}
	return -1
}
