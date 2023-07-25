package main

import "fmt"

func main() {
	// Declaração de uma variável e de um ponteiro
	var num int = 42
	var ptr *int

	// Atribuição do endereço da variável 'num' ao ponteiro 'ptr'
	ptr = &num

	fmt.Println("Valor da variável 'num':", num)
	fmt.Println("Endereço de memória de 'num':", &num)
	fmt.Println("Valor apontado pelo ponteiro 'ptr':", *ptr)
	fmt.Println("Endereço de memória apontado pelo ponteiro 'ptr':", ptr)

	// Modificando o valor da variável através do ponteiro
	*ptr = 100
	fmt.Println("Novo valor da variável 'num':", num)
}
