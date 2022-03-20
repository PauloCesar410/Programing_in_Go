package main

import "fmt"

func main() {

	fmt.Println("Estruturas de controle")

	fmt.Println("Digite um número")
	var numero uint
	fmt.Scanln(&numero)

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	if outronumero := numero; outronumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}

}
