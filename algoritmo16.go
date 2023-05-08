package main

import (
	"fmt"
	"strings"
)

func algoritmo16(x string) (string, error) {
	if len(x) == 0 {
		return "", fmt.Errorf("Erro: String vazia")
	}
	i := 0
	var y string = ""
	for i < len(x) {
		if strings.ContainsAny(string(x[i]), "AEIOUaeiou") {
			y += "*"
		} else {
			y += string(x[i])
		}
		i++
	}
	return y, nil
}
