package main

import "fmt"

func algoritmo8(x []int) ([]int, error) {
	if len(x) == 0 {
		return nil, fmt.Errorf("The slice is empty")
	}
	var i = 0
	newlist := []int{}
	for i < len(x) {
		if x[i]%2 == 0 {
			newlist = append(newlist, x[i])
		}
		i++
	}
	return newlist, nil
}
func main() {
	x := []int{0, 1, 2, 3, 4, 5}
	resultado, err := algoritmo8(x)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}
}
