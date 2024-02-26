package todolist

import "fmt"

func Matrices() {
	matriz := [...]int{0, 1, 2, 3, 4}
	matriz[0] = 322
	//iteracion modo 1
	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}
	//iteracion modo 2
	for indice, value := range matriz {
		fmt.Printf("Indice: %v valor: %v\n", indice, value)
	}

	matrizBidimencional := [3][3]int{{1, 2, 3}, {6, 5, 4}, {8, 9, 7}}

	fmt.Println(matrizBidimencional)

}
