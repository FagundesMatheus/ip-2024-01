package main

import "fmt"

func main() {
	//CASOS DE TESTE
	var arr01 = []int{5, 4, 3, 2, 1}
	var arr02 = []int{76, 34, 97, 12, 11, -2, 234}
	var arr03 = []int{24, 0, 4, 2, 0, 0, 6}

	//PRINTA OS CASOS DE TESTE ORDENADOS
	fmt.Println(SelectionSort(arr01))
	fmt.Println(SelectionSort(arr02))
	fmt.Println(SelectionSort(arr03))
}

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ { //PERCORRE TODOS OS ELEMENTOS DO SLICE
		currMinIndex := i                    //ADICIONA O ELEMENTO ATUAL COMO O ""MENOR"" E GUARDA SEU INDICE
		for j := i + 1; j < len(nums); j++ { //PERCORRE OS ELEMENTOS VERIFICANDO SE TEM ALGUM MENOR QUE NUMS[I]
			if nums[j] < nums[currMinIndex] {
				currMinIndex = j //GUARDA O INDEX DO MENOR NÚMERO ATUAL
			}
		}
		if currMinIndex != i { //SE O INDEX DO MENOR NÚMERO FOI ALTERADO
			nums[i], nums[currMinIndex] = nums[currMinIndex], nums[i] //TROCA DE LUGAR O MENOR NÚMERO ENCONTRADO COM O NÚMERO DO COMEÇO DESSA EXECUÇÃO
		}
	}
	return nums //RETORNA O SLICE ORDENADO
}
