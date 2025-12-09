package main

import (
	"fmt"
	"strings"
)

func main() {
	// Início
	fmt.Println("Olá, equipe! Meu nome é Alexandre, ou Lê, e estou pronto para começar os trabalhos!")
	// Desafio 1
	fmt.Println("A soma é:", soma(3, 5))
	// Desafio 2
	fmt.Println()
	frutas := []string{"Morango", "Abacate", "Laranja"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
	// Desafio 3
	fmt.Println()
	for i := 1; i < 6; i++ {
		linha := strings.Repeat("*", i)
		fmt.Println(linha)
	}
	// Desafio bônus
	fmt.Println()
	colegas := map[string]int{
		"lê":      27,
		"fanta":   19,
		"isarão":  26,
		"farinha": 21,
		"poru":    22,
	}

	for k, v := range colegas {
		fmt.Printf("%s, %d anos\n", k, v)
	}

}

// função soma
func soma(a int, b int) int {
	return a + b
}
