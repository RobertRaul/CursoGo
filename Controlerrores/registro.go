package controlerrores

import (
	"log"
	"os"
)

func Registrologs() {

	// log.SetPrefix("funcionRegisdtrolos():")
	// log.Print("eSTE ES UNA PRUENA")
	// log.Println("SALTO DE LINEA?")
	// log.Fatal("Registr de errores")

	log.SetPrefix("Registrologs():")
	file, err := os.OpenFile("errores.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Print("Soy un primer error!")
}
