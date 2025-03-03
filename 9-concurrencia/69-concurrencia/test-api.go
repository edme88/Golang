package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhrereintheinternet.com", //ejemplo de api que no existe
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api) //con la palabra "go" se hace concurrente en lugar de secuencial
	}

	time.Sleep(1 * time.Second) //esto se lo agregamos para que no diga: Ejecutar tomo: 0
	//El tiempo de suspenso no es la mejor forma de esperar. Se deberia usar canales

	elapsed := time.Since(start)
	fmt.Printf("Ejecutar tomo: %v", elapsed.Seconds())
	//Tiempos de ejecución SIN concurrencia
	//3.8892083
	//2.8918481
	//2.9676041
	//3.2982599
	//Tiempos de ejecución CON concurrencia
	//1.0001756

	//SIN CONCURRENCIA TIENEN ESTE ORDEN
// 	SUCCESS: https://management.azure.com está funcionando
// SUCCESS: https://dev.azure.com está funcionando
// SUCCESS: https://api.github.com está funcionando
// SUCCESS: https://outlook.office.com está funcionando
// ERROR: https://api.somewhrereintheinternet.com está caído
// SUCCESS: https://graph.microsoft.com está funcionan

//CON CONCURRENCIA TIENEN ESTE ORDEN
// ERROR: https://api.somewhrereintheinternet.com está caído
// SUCCESS: https://api.github.com está funcionando
// SUCCESS: https://management.azure.com está funcionando
// SUCCESS: https://dev.azure.com está funcionando
// SUCCESS: https://graph.microsoft.com está funcionando
// SUCCESS: https://outlook.office.com está funcionando

}

func checkAPI(api string){
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR: %s está caído\n", api)
		return
	}
	fmt.Printf("SUCCESS: %s está funcionando\n", api)
}