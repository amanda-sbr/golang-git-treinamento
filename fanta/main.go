package main

import "fmt"

func main(){

	// Exercicio 0
	fmt.Println("Olá, equipe! Eu sou o Fanta, será um prazer aprender e trabalhar com vocês! :)");


	// Desafio 1
	var n1 int;
	var n2 int;

	fmt.Println("Insira os valores de n1 e n2");
	fmt.Scan(&n1, &n2);
	fmt.Println("A soma é:", soma(n1, n2));


	// Desafio 2

	frutas := []string{"Maçã", "Uva", "Pera", "Abacaxi", "Banana"};

	for i := 0; i < len(frutas); i++ {
		fmt.Println(frutas[i]);
	}
	
}


func soma(n1 int , n2 int) int{
	return n1 + n2;
}
