package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

type figura2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 4, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces
	myInterface := []interface{}{"hola", 12, 4.9}
	fmt.Println(myInterface...)
}
