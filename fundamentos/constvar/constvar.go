package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// bloco de construção de constantes
	const (
		a = 1
		b = 2
	)

	// bloco de construção de variáveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// declarar de uma vez
	var e, f bool = true, false
	fmt.Println(e, f)

	// declarar de uma vez reduzida
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

	// sempre precisa usar a variável
}
