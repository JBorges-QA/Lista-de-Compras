package main

import (
	"fmt"
)

var lista []string

// Adicionar itens a lista de compras
func adicionarItens(itens ...string) []string {
	lista = append(lista, itens...)
	return lista
}

// Remover item da lista de compras
func removerItem(indice int) []string {
	indice = indice - 1
	fmt.Printf("Item %v removido com sucesso! \n", lista[indice])
	lista = append(lista[:indice], lista[indice+1:]...)
	return lista
}

// Exibir itens da lista
func exibirItens(lista []string) {
	fmt.Println("==========LISTA============")
	for i, item := range lista {
		fmt.Printf("%v - %v \n", i+1, item)
	}
	fmt.Println("===========================")
}
func main() {
	lista := adicionarItens("Arroz", "Feijão", "Carne", "Leite", "Mel")

	exibirItens(lista)

	lista = removerItem(5)

	exibirItens(lista)
}
