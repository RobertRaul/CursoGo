package controlerrores

import (
	"fmt"
	"os"
)

func Inicio() {

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	_, err = file.Write([]byte("QUE MIERDA A XD"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
