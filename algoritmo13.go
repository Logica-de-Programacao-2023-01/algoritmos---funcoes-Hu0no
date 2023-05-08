package main

import (
	"fmt"
	"strconv"
	"strings"
)

func algoritmo13(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Erro: Número inválido")
	}
	string := strconv.Itoa(n)
	digitos := strings.Split(string, "")
	soma := 0
	i := 0
	for i < len(digitos) {
		digito, err := strconv.Atoi(digitos[i])
		if err != nil {
			return 0, err
		}
		soma += digito
		i++
	}
	return soma, nil
}
