package estructuradatos

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListadoTareas struct {
	tareas []Tarea
}

// metodo para agregar tarea
func (l *ListadoTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para marca tarea como completado
func (l *ListadoTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tarea
func (l *ListadoTareas) editarTrea(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para eliminar tarea
func (l *ListadoTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}
func Menu() {
	//instancia de lista de tareas
	lista := ListadoTareas{}

	//instancia bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println(
			"Selecciona una opcion:\n",
			"1: Agregar Tarea\n",
			"2: Marca tarea como completada\n",
			"3: Editar Tarea\n",
			"4: Eliminar Tarea\n",
			"5: Salir Tarea",
		)
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.agregarTarea(t)
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada:")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea editar:")
			fmt.Scanln(&index)

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.editarTrea(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea eliminara:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opcion Invalida")
		}
		//listar todas la tareas
		fmt.Println("Lista de Tareas")
		fmt.Println("==================================")
		for i, t := range lista.tareas {
			fmt.Printf("Identificador: %d - Nombre: %s - Descripcion %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("==================================")

	}
}
