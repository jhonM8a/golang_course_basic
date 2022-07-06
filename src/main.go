package main //Indica la carpeta donde esta, pero como no hay ponemos el main

import "fmt" //	Libreria para imprimir

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %d cursos\n", nombre, cursos) // %v cuando no sabemos el tipo de dato

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos) // Guarda el resultado en un string
	fmt.Println(message)

	//Tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
