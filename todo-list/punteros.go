package todolist

import "fmt"

func Punter() {

	var x int = 10
	var p *int = &x

	fmt.Println(x)
	fmt.Println(p)
	Editar(&x)
	fmt.Println(x)
}

func Editar(x *int) {
	*x = 20
}

type Perrito struct {
	nombre string
	edad   int
}

func (m *Perrito) saluda() {
	fmt.Println("Hola mi nombre es", m.nombre)
}

func Usar() {
	k := Perrito{"Pepe", 20}
	k.saluda()
}
