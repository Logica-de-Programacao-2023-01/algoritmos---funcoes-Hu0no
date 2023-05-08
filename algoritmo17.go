package main

import (
	"fmt"
	"sort"
	"strings"
)

func algoritmo17(x []string) (string, error) {
	if len(x) == 0 {
		return "", fmt.Errorf("Erro: Slice vazio")
	}
	i := 0
	for i < len(x) {
		strings.ToLower(x[i])
		i++
	}
	sort.Strings(x)
	i = 0
	y := ""
	for i < len(x) {
		y += x[i]
		y += " "
		i++
	}
	return y, nil
}
