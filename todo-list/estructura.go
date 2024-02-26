package todolist

import "fmt"

type Mascota struct {
	nombre string
	edad   int
}

func Estruc() {

	var m Mascota
	m.nombre = "Tomas"
	m.edad = 20

	mas := Mascota{nombre: "Pepe", edad: 5}

	fmt.Println(mas)

}
