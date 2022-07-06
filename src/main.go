package main //Indica la carpeta donde esta, pero como no hay ponemos el main

import (
	"fmt" //	Libreria para imprimir
	"log"
	"strconv"
)

func ValidarParOImpar(number int) {
	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

}

func ValidateUserPassword(user, password string) bool {
	if user == "jochoa" && password == "test123" {
		return true
	} else {
		return false
	}
}

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// Text to number
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	//Reto 1: Ver si un número es par o impar
	ValidarParOImpar(3)
	//Reto 2: Validar usuario y contraseña
	userValid := ValidateUserPassword("jochoa", "test123")
	if userValid {
		fmt.Println("Usuario valido")
	}
}
