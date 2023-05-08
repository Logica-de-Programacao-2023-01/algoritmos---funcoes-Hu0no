package main

import (
	"fmt"
	"math"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	maximo := math.Inf(-1) // Define o máximo como o menor número possível
	minimo := math.Inf(1)  // Define o mínimo como o maior número possível
	var soma, media, denominador float64
	if len(numbers) <= 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia.") // erro caso a lista seja igual a 0
	}
	for i := 0; i < len(numbers); i++ {
		soma += float64(numbers[i]) //soma de todos os numeros
		denominador++               //soma dos denominadores ate o len(numbers)
		valorMax := numbers[i]
		valorMin := numbers[i]
		if float64(valorMax) > maximo {
			maximo = float64(valorMax)
		}
		if float64(valorMin) < minimo {
			minimo = float64(valorMin)
		}
	}
	media = (soma / (denominador))
	return int(maximo), int(minimo), media, nil
}
func main() {
	slice := []int{}
	fmt.Print(FindMinMaxAverage(slice))
}
