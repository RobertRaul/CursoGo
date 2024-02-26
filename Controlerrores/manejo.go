package controlerrores

import (
	"fmt"
	"strconv"
)

func Control() {
	str := "123e"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("El numero es:", num)

}

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("no se puede dividir entre")
	}
	return dividendo / divisor, nil
}

func Run() {
	result, erro := divide(10, 0)
	if erro != nil {
		fmt.Println("Error:", erro)
		return
	}
	fmt.Println(result)
}
