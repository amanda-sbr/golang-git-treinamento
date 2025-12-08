package main

import "fmt"

func main() {
	fmt.Println("Olá, mundo! Meu nome é Poru e estou pronto para começar!")
	fmt.Println(soma(1, 2))

	// slice de strings com 3 frutas
	frutas := []string{"maçã", "banana", "laranja"}

	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	// print do triangulo de * 
	for i := 0; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n")

	// map onde a chave eh uma string (nome) e o valor eh um int (idade) 
	// do pessoal que estava na call
	idades := map[string]int{
		"Fanta": 20,
		"Poru": 21,
		"Ana": 22,
		"Brennda": 23,
		"Farinha": 24,
		"Duda": 25,
		"Isarao": 26,
		"Lisso": 27,
		"Le": 28,
		"Madu": 29,
	}

	for nome, idade := range idades {
		fmt.Println(nome, idade)
	}

}

func soma(a int, b int) int {
	return a + b
}






