package main

import (
	"strings"
)

func quantidadevogais(x string) int {
	qtd := 0
	for y := range x {
		if strings.ContainsAny(string(x[y]), "aeiouAEIOU") {
			qtd++
		}
	}
	return qtd
}
