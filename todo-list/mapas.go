package todolist

import "fmt"

func Mapitas() {
	colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colores)
	colores["negro"] = "#00000"
	fmt.Println(colores["rojo"])
	fmt.Println(colores)

	valor, ok := colores["amarillo"]
	if ok {
		fmt.Println(valor + " Si existe")
	} else {
		fmt.Println("No existe esta clave")
	}
	fmt.Println(valor, ok)
	delete(colores, "azul")
	fmt.Println(colores)

	for clave, valor := range colores {
		fmt.Printf(" Clave: %v Valor:%v\n", clave, valor)
	}
}
