package main

import "fmt"

func main(){

	fmt.Println("Olá, equipe! Eu sou o Fanta, será um prazer aprender e trabalhar com vocês! :)");

	var n1 int;
	var n2 int;

	fmt.Println("Insira os valores de n1 e n2");
	fmt.Scan(&n1, &n2);
	fmt.Println("A soma é:", soma(n1, n2));


}


func soma(n1 int , n2 int) int{
	return n1 + n2;
}
