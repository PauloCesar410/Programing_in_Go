package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	//números inteiros
	var numero int64 = -100000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//int32=rune
	var numero3 rune = 1234
	fmt.Println(numero3)

	//byte=uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//números reais
	var numeroReal11 float32 = 123.42
	fmt.Println(numeroReal11)

	//inferência
	numeroReal2 := 1234.77
	fmt.Println(numeroReal2)

	//Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//reorna o valor da tabela ASCI
	char := 'B'
	fmt.Println(char)

	//tipo booleano
	var booleano1 bool = true
	fmt.Print(booleano1)

	//tipo erro
	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

}
