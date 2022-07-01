package main //Indica la carpeta donde esta, pero como no hay ponemos el main

import "fmt" //	Libreria para imprimir

func main() {
	// Delcaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.15
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Delacación de variables entera: Si las variables no se usan dará erros al ejecutar o compilar
	base := 12 // los dos puntos indican que es una variable que no ha sido delcarada anteriormete
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values:Es un valor que tiene por defecto al crear una variable y no se le asigna un valor
	var a int     // valor por defecto 0
	var b float64 // 0
	var c string  // valor por decto no nulo, como un string vacio
	var d bool    //false
	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado // Go asigna el tipo de esta manera
	fmt.Println("Area cuadrado:", areaCuadrado)
}
