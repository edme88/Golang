# GO

- [Instalación](#instalacion)
- [Recursos](#recursos)

## Instalación

1. Instalar go desde https://go.dev/doc/install
2. Instalar VSC https://code.visualstudio.com/ o alternativamente GoLand https://www.jetbrains.com/es-es/go/download/
3. Ejecutando `go version` en la terminal verificar la versión instalada
4. Ejecutando `go env` verificar el PATH
5. Crear el path usando `mkdir go` y acceder con `cd go`
6. Crear la carpeta de source con `mkdir src` y acceder con `cd src`
7. Ejecutar `code .` para abrir el VSC en esa carpeta

## Ejecución

1. Crear un archivo **hola.go**
2. En el archivo escribir:

```
package main

import "fmt"

func main(){
	fmt.Println("Hola mundo!")
}
```

3. Ejecutar `go run hola.go` para ejecutar
4. Ejecutar `go build hola.go` para crear el archivo ejecutable
5. Para el manejador de modulos `go mod init holamundo`. Definir y gestionar los módulos y dependencias de go.
6. Agregar paquete quote `go get rsc.io/quote`

## Recursos

- [Curso GoLand](https://github.com/alexroel/curso-golang/blob/main/sections/01-get-started-with-go.md)
- [Get Started](https://go.dev/doc/tutorial/getting-started)
- [Learn](https://learn.microsoft.com/es-es/training/modules/go-get-started/)
- [Go by Example](https://gobyexample.com/hello-world)
- [Tour](https://go.dev/tour/welcome/1)
- [Learn Info](https://learn.microsoft.com/es-es/training/modules/go-variables-functions-packages/)

type nul > contact-manager.go

go mod init github.com/edme88/greetings
se crea un go.mod para manejar los modulos
go mod edit -replace github.com/edme88/greetings=../greetings

go mod tidy quita los paquetes innecesarios y agrega los necesarios

go test //para ejecutar las pruebas de un modulo
go tets -v //para tener mas info de los tests

bajar y compilar
go build main.go

//para compilar una aplicación
go build -o hello main.go

Agregar a las variables de ambiente para poder ejecutarlo
pwd
C:\Users\Darwoft\go\src\hola-mundo\hello

nano.profile
export PATH=$PATH:Users\Darwoft\go\src\hola-mundo\hello

go install //para instalar la aplicación

## Para Instalar un paquete

Ejecutar:

```bash
go get -u github.com/edme88/greetings
```
