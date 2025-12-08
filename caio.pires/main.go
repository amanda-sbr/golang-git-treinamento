package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Opa, meu nome é Caio e eu sou estagiário da Monks!\n\n")

	fmt.Printf("Resultado da soma: %v\n\n", soma(5, 5))

	frutas := []string{"Maçã", "Uva", "Banana"}

	for index, fruta := range frutas {
		fmt.Printf("Fruta n° %v: %s\n", index, fruta)
	}

	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}
}

func soma(a int, b int) int {
	return a + b
}
