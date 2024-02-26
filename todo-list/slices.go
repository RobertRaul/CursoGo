package todolist

import "fmt"

func Rebana() {
	diasSemana := []string{"A", "B", "C", "E"}

	diaSeparado := diasSemana[0:4]
	fmt.Println(diaSeparado)

	//AGREGAR
	diaSeparado = append(diaSeparado, "OJOS", "PIES", "ILEY", "OMAET")
	fmt.Println(diaSeparado)

	// quitar partes del SLICE desde la posicion 4 hasta la 6
	diaSeparado = append(diaSeparado[0:4], diaSeparado[6:]...)
	fmt.Println(diaSeparado)

	fmt.Println(len(diaSeparado)) //longitud
	fmt.Println(cap(diaSeparado)) //capacidad

	nombres := make([]string, 5, 10)
	nombres[0] = "Robert"
	fmt.Println(nombres)

	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := make([]int, 6)

	fmt.Println(copy(slice1, slice2))
}
