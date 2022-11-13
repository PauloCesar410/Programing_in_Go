package main

import "fmt"

func funcao(x float32) float32 {
	return x
}

func main() {

	var soma float32 = 0
	fmt.Println("Informe o valor do limite inferior do intervalo: ")
	var a float32
	fmt.Scanln(&a)
	fmt.Println("Informe o valor do limite superior do intervalo: ")
	var b float32
	fmt.Scanln(&b)
	var n float32
	fmt.Println("Informe a quantidade de partições do intervalo: ")
	fmt.Scanln(&n)
	var deltax float32 = (b - a) / n
	var xi float32

	for i := 0; i <= int(n); i++ {
		xi = a + float32(i)*float32(deltax)
		soma = (soma) + float32(funcao(float32(xi)))*float32(deltax)
	}

	fmt.Println("O valor da integral no intervalo [a,b] é:  ", soma)
}
