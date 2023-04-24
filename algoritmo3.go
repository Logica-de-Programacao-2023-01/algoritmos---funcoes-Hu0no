package main

func concatenação(x []string) string {
	y := 0
	var resultado string
	resultado = ""
	for y < len(x) {
		resultado += " "
		resultado += x[y]
		y++
	}
	return resultado
}
