package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiro é uma referência de memória

	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) //desreferênciação

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)

}