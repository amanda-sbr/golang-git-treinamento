package main

import "fmt"

func main() {
	fmt.Printf("Opa, meu nome é Caio e eu sou estagiário da Monks!\n\n")

	fmt.Printf("Resultado da soma: %v\n\n", soma(5, 5))

	frutas := []string{"Maçã", "Uva", "Banana"}

	for index, fruta := range frutas {
		fmt.Printf("Fruta n° %v: %s\n", index, fruta)
	}
}

func soma(a int, b int) int {
	return a + b
}
