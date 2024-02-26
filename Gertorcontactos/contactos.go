package gertorcontactos

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contactos struct {
	Name  string `json: "name"`
	Email string `json: "email"`
	Phone string `json: "phone"`
}

// guarda copntactos en un archivbo json
func SaveContactToFile(contact []Contactos) error {
	file, erro := os.Create("contactos.json")
	if erro != nil {
		return erro
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	erro = encoder.Encode(contact)
	if erro != nil {
		return erro
	}
	return nil
}

//cargar contactos desde un archivo json

func LoadContactosFromFile(contact *[]Contactos) error {
	file, err := os.Open("contactos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contact)
	if err != nil {
		return err
	}
	return nil
}

func Inicio() {
	//slice de contactos
	var contacts []Contactos

	//cargar contactos existente desdel archivo
	err := LoadContactosFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	//crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== GESTOR DE CONTACTOS ===\n",
			"1. Agregar un Contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		switch option {
		case 1:

			var c Contactos
			fmt.Printf("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Printf("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Printf("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := SaveContactToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}
		case 2:
			fmt.Println("======================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("======================")
		case 3:
			return
		default:
			fmt.Println("Opcion Invalida")
		}

	}
}
