package main

import "fmt"
type Persona struct  {
	nombre string
	edad int
	correo string
}

func main(){
	//Estructuras
	var p Persona
	p.nombre = "Agus"
	p.edad = 36
	p.correo = "agus@email.com"

	fmt.Println(p)

	per := Persona {"Alex", 28, "alex@email.com"}
	fmt.Println(per)
	//Punteros y métodos
	//Acceder a una variable y modificarla
	var x int = 10
	//accede a la referncia de memoria
	var z *int = &x
	fmt.Println(&x, z)

	var t int = 10
	fmt.Println(t)
	editar(&t)
	fmt.Println(t)

	p.diHola()
	per.diHola()
	//Proyecto: Lista de tareas
}

func editar(algo *int){
	*algo = 20
}

//Método, pertenece a una estrutura
//Las funciones son independientes
func (p *Persona) diHola(){
	fmt.Println("Hola, mi nombre es: ", p.nombre)
}