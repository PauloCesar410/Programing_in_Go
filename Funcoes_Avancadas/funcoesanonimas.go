package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprint("Recebido-> %s %d", texto, 10)
	}("Passando o Parâmetro")

	fmt.Println(retorno)
}
