package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2022} // Forma de instanciar
	fmt.Println(myCar)

	//Otra forma de instaciar
	var otherCar car
	otherCar.brand = "Ferrary"
	fmt.Println(otherCar)

}
