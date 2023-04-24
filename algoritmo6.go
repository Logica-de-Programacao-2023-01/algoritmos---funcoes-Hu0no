package main

import "fmt"

func algoritmo6(x []string) (string, error) {
	if len(x) == 0 {
		fmt.Errorf("cannot concatenate 0 words")
	}
	var y string = ""
	for _, z := range x {
		y += z
		y += ", "
	}
	return y, nil
}
