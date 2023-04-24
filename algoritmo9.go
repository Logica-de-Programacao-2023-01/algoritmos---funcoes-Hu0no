package main

import (
	"fmt"
	"strings"
)

func algoritmo9(x string) ([]string, error) {
	if x == "" {
		return nil, fmt.Errorf("The string is empty")
	}
	y := strings.Split(x, " ")
	return y, nil
}
