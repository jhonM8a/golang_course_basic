package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() { // De esta manera podemos añadirle funciones a nuestro struc
	fmt.Println(myPc.brand, "pong")
}

func (myPc *pc) duplicateRam() { // Se añade la función al struct y además accedmos al valor del struct que esta la posición de memoria
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a // & accedemos a la posicíon de memoria de la variable a

	fmt.Println(a, b)
	fmt.Println(a, *b) // El * es para acceder al valor de la variable que se encuentra en memoria

	*b = 100       // Se modifca la posición de memoria que es la misma para a y b
	fmt.Println(a) // Al imprimir a queda con otro valor a pesar de que no modificamos directamente la variable pero si la posición de memoria

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

}
