package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// canal := make(chan int)
	// canal <- 15
	// valor := <- canal

	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhrereintheinternet.com", //ejemplo de api que no existe
		"https://graph.microsoft.com",
	}

	//canales: comunicación de go-runtime
	// permite enviar y recibir valores entre los go-runtimes
	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)

	for i := 0; i < len(apis);i++ {
		fmt.Println(<- ch)
	}
	// Si se agregan más "fmt.Println(<- ch)" la aplicación NO termina
	
	elapsed := time.Since(start)
	fmt.Printf("Ejecutar tomo: %v", elapsed.Seconds())
	//Ejecutar tomo: 0.9943115
	//Ejecutar tomo: 0.9306811
}

func checkAPI(api string, ch chan string){
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: %s está caído\n", api) //!Cambio a Sprintf
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s está funcionando\n", api) //!Cambio a Sprintf
}