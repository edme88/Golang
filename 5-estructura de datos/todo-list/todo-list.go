package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre string
	descripcion string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea){
	l.tareas = append(l.tareas, t)
}

func (l*ListaTareas) marcarCompletado(index int){
	l.tareas[index].completado = true
}

func (l *ListaTareas) editarTarea(index int,t Tarea){
	l.tareas[index] = t
}

func (l *ListaTareas) eliminarTarea(index int){
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main(){
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)

	for{
		var opcion int
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Marcar tarea como completada")
		fmt.Println("3. Editat Tarea")
		fmt.Println("4. Eliminar Tarea")
		fmt.Println("5. Salir")

		fmt.Println("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion{
		case 1:
			var t Tarea
			fmt.Println("Ingrese nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el indice de la tarea que desea modificar: ")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nuevo nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la nueva descripción de la tarea")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Eliminar Tarea correctamente")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción incorrecta.")
		}
		fmt.Println("Lista de tareas: ")
		fmt.Println("===================================")
		for i, t:= range lista.tareas{
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)
		}
		fmt.Println("===================================")
	}
}