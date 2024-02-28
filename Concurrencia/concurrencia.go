package concurrencia

import (
	"fmt"
	"net/http"
	"time"
)

func Concu() {

	start := time.Now()

	apis :=
		[]string{
			"https://management.azure.com",
			"https://dev.azure.com",
			"https://api.github.com",
			"https://outlook.office.com/",
			"https://api.somewhereintheinternet.com",
			"https://graph.microsoft.com",
		}
	canales := make(chan string)

	for _, urls := range apis {
		go checkApi(urls, canales)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-canales)
	}

	elapsed := time.Since(start)
	fmt.Printf("Listo Tomo  %v segundos\n", elapsed.Seconds())

}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: el api %s esta caido!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCES: el api %s esta operativo\n", api)
}
