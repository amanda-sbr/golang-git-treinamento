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

}

func soma(a int, b int) int {
	return a + b
}






