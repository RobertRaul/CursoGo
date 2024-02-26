package main

import (
	contact "cursogo/gertorcontactos"
	"fmt"
	"math/rand/v2"
)

func main() {
	contact.Inicio()
}

func jugar() {
	numAleatorio := rand.IntN(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos <= maxIntentos {
		intentos++
		fmt.Printf("Ingrese un numero (intentos restantes %v):", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)
		if numAleatorio == numIngresado {
			fmt.Println("Felicitaciones, adivinaste el numero:")
			jugarNuevamente()

			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es MAYOR")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")

		}
	}
	fmt.Println("Se acabaron los intentos. El numero era:", numAleatorio)
	jugarNuevamente()

}

func jugarNuevamente() {
	var eleccion string
	fmt.Print("Quieres jugar nuevamente? (s/n):")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Grcias")
	default:
		fmt.Println("Eleccionar incorrecta")
		jugarNuevamente()
	}
}
