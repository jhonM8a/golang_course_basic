package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import (
	pk "curso_golang_platzi/src/mypackage" // El pk es un alias
	"fmt"
)

func main() {
	var myCar pk.CarPublic // Si el estruct empieza con mayuscula es publico en minuscula privado
	myCar.Brand = "Ferrai"
	myCar.Year = 2022
	fmt.Println(myCar)

	/*
		Como el struct empieza con minuscula es privado
		var myOtherCar pk.carPrivate
		fmt.Println(carPrivate)
	*/

	pk.PrintMessage() //Como comienza con Mayuscula la función es publica
	/*
		A tener en cuenta, si sale el error "is not in GOROOT" con el comando "go mod init" se puede
		solucionar

	*/
}
