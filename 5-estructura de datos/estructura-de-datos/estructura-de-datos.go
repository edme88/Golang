package main

import "fmt"

func main(){
	//Matrices: Almacena conjunto de elementos del mismo tipo
	//Longitud fija - cualquier tipo de datos
	var miMatriz [5]int
	miMatriz[0]=10
	miMatriz[1]=20
	fmt.Println(miMatriz)

	var miMatriz2 = [5]int{12,33,56,77,34}
	fmt.Println(miMatriz2)

	var miMatriz3 = [...]int{12,33,56,77,34,343,34,12,43}
	fmt.Println(miMatriz3)

	for i:=0;i<len(miMatriz3);i++{
		fmt.Println(miMatriz3[i])
	}

	//walruls = morza 
	for index, value := range miMatriz3{
		fmt.Printf("Índice: %d, valor %d\n", index, value)
	}

	for _, value := range miMatriz3{
		fmt.Printf("Valor %d\n", value)
	}

	//filas columnas tipo de datos
	var matrizBidi = [3][4] int{{1,2,4},{5,6,7},{3,5,9}}
	fmt.Println(matrizBidi)
	//Rebanadas: Slice - Sección de un Arreglo
	var rebanada []int
	fmt.Println(rebanada)

	diasSemana := []string{"Domingo","Lunes", "Martes",
		"Miércoles","Jueves", "Viernes", "Sábado"}
	fmt.Println(diasSemana)

	//crear una rebanada a partir de otra rebanada
	diaRebanada := diasSemana[1:6]
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	diaRebanada = append(diaRebanada, "Osvaldo", "lalala")
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	diaRebanada = append(diaRebanada[:1], diaRebanada[2:]...)
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))


	nombres := make([]string, 5, 10)
	nombres[0]="Agus"
	nombres[4]="Final"
	fmt.Println(nombres)

	rebanada3 := []int{1,2,3,4,5}
	rebanada4 := make([]int, 5)
	fmt.Println(rebanada3)
	fmt.Println(rebanada4)
	//Copia de rebanada 3 a rebanada 4 (o sea que queda con valores, no con 0)
	copy(rebanada4, rebanada3)
	fmt.Println(rebanada3)
	fmt.Println(rebanada4)
	//Mapas
	colores := map[string]string{
		"rojo":"#FF0000",
		"verde": "#00FF00",
		"azul": "#0000FF",
	}
	fmt.Println(colores)
	fmt.Println(colores["azul"])
	colores["negro"] = "#000000"
	//No respeta el orden - Pone los elementos en orden alfabético
	fmt.Println(colores)

	valorNew, ok := colores["rojo"]
	fmt.Println(valorNew, ok)

	valorNew2, ok2 := colores["blanco"]
	fmt.Println(valorNew2, ok2)
	if ok2{
		fmt.Println("Si existe la clave")
	}else{
		fmt.Println("No existe esta clave")
	}

	//eliminar un elemneto del Mapa
	delete(colores, "rojo")
	fmt.Println(colores)

	for clavel, valore := range colores{
		fmt.Printf("Clave: %s, Valor: %s\n", clavel, valore)
	}
	//Estructuras
	

	//Punteros y métodos
	//Proyecto: Lista de tareas
}