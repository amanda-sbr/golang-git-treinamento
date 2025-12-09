package main

import "fmt"

func main() {
	fmt.Println(`Olá equipe, eu sou a Isarão e vou cantar uma música para vocês: 

Sapo cururu
Na beira do rio
Quando o sapo canta, ô maninha
É porque tem frio
A mulher do sapo
Deve estar lá dentro
Fazendo rendinha, ô maninha
Para o casamento
Sapo cururu
Na beira do rio
Quando o sapo canta, ô maninha
É porque tem frio
A mulher do sapo
Deve estar lá dentro
Fazendo rendinha, ô maninha
Para o casamento
Sapo cururu
Na beira do rio
Quando o sapo canta, ô maninha
É porque tem frio
A mulher do sapo
Deve estar lá dentro
Fazendo rendinha, ô maninha
Para o casamento
Sapo cururu
Na beira do rio
Quando o sapo canta, ô maninha
É porque tem frio
A mulher do sapo
Deve estar lá dentro
Fazendo rendinha, ô maninha
Para o casamento
Fazendo rendinha, ô maninha
Para o casamento`)

	//desafio 01
	fmt.Println("\n\nO resultado da soma é: ", soma(5, 5))

	//desafio 02
	slice := []string{"Uva", "Pera", "Maçã", "Salada mista"}
	fmt.Printf("\nImprimindo frutas:\n")
	for _, fruta := range slice {
		fmt.Println(fruta)
	}

	//desafio 03
	fmt.Printf("\nImprimindo asteriscos\n")
	for i := 1; i <=5; i++ {
		for j := 1; j <= i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	//desafio extra
	colegas := map[string]int{
		"Isarão": 28, 
		"Farinha": 24, 
		"Poru": 21,
		"Lê": 27,
		"Fanta": 19,
	}
	fmt.Printf("\nNome e idade dos novos estagiários:\n")
	for nome, idade := range colegas{
		fmt.Println(nome, "-", idade)
	}

}

func soma(a int, b int) int {
	return a + b
}
