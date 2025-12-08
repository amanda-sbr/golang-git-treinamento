package main

import (
	"fmt"
	"strings"
)

func main() {
	//Inicio
	fmt.Printf("Opa, meu nome é Caio e eu sou estagiário da Monks!\n\n")

	//Desafio 1
	fmt.Printf("Resultado da soma: %v\n\n", soma(5, 5))

	//Desafio 2
	frutas := []string{"Maçã", "Uva", "Banana"}

	for index, fruta := range frutas {
		fmt.Printf("Fruta n° %v: %s\n", index, fruta)
	}

	//Desafio 3
	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}

	//Bônus (não sei se as idades estão certas)
	fmt.Println()
	idadesColegas := map[string]int{
		"Poru":    21,
		"Fanta":   18,
		"Farinha": 24,
		"Lê":      30,
	}
	for nome, idade := range idadesColegas {
		fmt.Printf("Nome:%s, Idade: %v\n", nome, idade)
	}
}

func soma(a int, b int) int {
	return a + b
}
