package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Olá, equipe! Meu nome é Alexandre, ou Lê, e estou pronto para começar os trabalhos!")

	fmt.Println("A soma é:", soma(3, 5))

	fmt.Println()
	frutas := []string{"Morango", "Abacate", "Laranja"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	fmt.Println()
	for i := 1; i < 6; i++ {
		linha := strings.Repeat("*", i)
		fmt.Println(linha)
	}

}

func soma(a int, b int) int {
	return a + b
}
