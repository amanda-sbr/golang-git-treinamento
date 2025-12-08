package main

import "fmt"

func main() {

	fmt.Println("Olá, equipe! Meu nome é Alexandre, ou Lê, e estou pronto para começar os trabalhos!")

	fmt.Println("A soma é:", soma(3, 5))

	fmt.Println()
	frutas := []string{"Morango", "Abacate", "Laranja"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

}

func soma(a int, b int) int {
	return a + b
}
