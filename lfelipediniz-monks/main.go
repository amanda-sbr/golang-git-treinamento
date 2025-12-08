package main

import "fmt"

func main() {
	fmt.Println("\nOlá, mundo! Meu nome é Poru e estou pronto para começar!")

	resultadoSoma := soma(1, 2)
	fmt.Println("\nDesafio 1: Resultado da soma:\n1 + 2 = ",resultadoSoma)

	fmt.Println("\nDesafio 2: Lista de frutas com Slice")
	frutas := []string{"maçã", "banana", "laranja"}
	imprimirSlice(frutas)

	fmt.Println("\nDesafio 3: Triângulo de asteriscos")
	imprimirTriangulo(5)

	fmt.Println("\nDesafio 4: Idades da equipe")
	idadesDaEquipe := criarMapIdades()
	imprimirMap(idadesDaEquipe)
}

// --------------------------------------------------
// Funções dos desafios 
// --------------------------------------------------

// soma retorna a soma de dois inteiros
func soma(a int, b int) int {
	return a + b
}

// itera e imprime os elementos de um slice de strings
func imprimirSlice(itens []string) {
	fmt.Println("--- Lista de Frutas ---")
	for i, item := range itens {
		fmt.Printf("%d: %s\n", i, item)
	}
}

// imprime um triângulo de asteriscos com a altura especificada
func imprimirTriangulo(altura int) {
	fmt.Println("--- Triângulo de Asteriscos ---")
	for i := 0; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// cria e retorna um map de nomes para idades
func criarMapIdades() map[string]int {
	// pessoal que estava na call
	return map[string]int{
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
}

// itera e imprime as chaves e valores de um map de string para int
func imprimirMap(dados map[string]int) {
	fmt.Println("--- Idades Inventadas da Equipe ---")
	for nome, idade := range dados {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}
}