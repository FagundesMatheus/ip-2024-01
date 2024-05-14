package main

import "fmt"

func main() {
	//CASOS DE TESTE
	var arr01 = []int{5, 4, 3, 2, 1}
	var arr02 = []int{76, 34, 97, 12, 11, -2, 234}
	var arr03 = []int{24, 0, 4, 2, 0, 0, 6}

	//PRINTA OS CASOS DE TESTE ORDENADOS
	fmt.Println(ISort(arr01))
	fmt.Println(ISort(arr02))
	fmt.Println(ISort(arr03))
}
func ISort(nums []int) []int {

	for i := 1; i < len(nums); i++ { //INICIA UM LOOP QUE PERCORRE TODOS OS ELEMENTOS DADOS NA SLICE
		j := i                             //"J" RECEBE O VALOR DE "I" NA EXECUÇÃO ATUAL DO LOOP
		for j > 0 && nums[j] < nums[j-1] { //LOOP ENQUANTO J FOR MAIOR QUE ZERO E O ELEMENTO ATUAL FOR MENOR QUE O ELEMENTO ANTERIOR
			nums[j], nums[j-1] = nums[j-1], nums[j] //TROCA A POSIÇÃO DOS ELEMENTOS
			j--                                     //"J" TEM SEU VALOR REDUZIDO PARA VERIFICAR NOVAMENTE OS ELEMENTOS DA POSIÇÃO ANTERIOR
		}
	}
	return nums //RETORNA A SLICE ORDENADA
}
