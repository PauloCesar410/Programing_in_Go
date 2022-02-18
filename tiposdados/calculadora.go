package main

import (
	"fmt"
)

func menu() {

	fmt.Println("Qual operação deseja fazer?")
	fmt.Println("1)Soma")
	fmt.Println("2)Subtração")
}

func main() {

	var a, b, resposta int

	fmt.Println("Bem-vindo(a) a calculadora!")

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scanln(&a)

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scanln(&b)

	menu()
	fmt.Scanln(&resposta)
	switch resposta {
	case 1:
		var soma int = a + b
		fmt.Println(soma)
	case 2:
		var subtracao int = a - b
		fmt.Println(subtracao)
	default:
		fmt.Println("Número inválido!")
	}
}
