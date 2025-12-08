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

	for _, fruta := range frutas{
		fmt.Println(fruta);
	}


	// Desafio 3

	var qtdLinhas int;
	fmt.Println("Insira a quantidade de linhas a serem impressas");
	fmt.Scan(&qtdLinhas);

	for i := 1; i <= qtdLinhas; i++{
		for j := 1; j <= i; j++{
			fmt.Print("*");
		}

		fmt.Println(); // pra quebrar linha no final
	}

	// Desafio 4 (Bônus)

	pessoas := map[string]int{}; // nome --> idade
	var option int;

	for{
		fmt.Println(`Escolha uma opçao:
	0. Quebrar loop
	1. Adicionar novo nome e idade`);

		fmt.Scan(&option);

		if option == 0{
			break;
		}

		var nome string;
		var idade int;

		fmt.Scan(&nome, &idade);

		pessoas[nome] = idade;

	}

	for nome, idade := range pessoas{
		fmt.Println(nome, idade);
	}

}




func soma(n1 int , n2 int) int{
	return n1 + n2;
}
