package main

import "fmt"

func main() {
	fmt.Println("Olá, equipe! Meu nome é Alexandre, ou Lê, e estou pronto para começar os trabalhos!")
	fmt.Println("A soma é:", soma(3, 5))
}

func soma(a int, b int) int {
	return a + b
}
