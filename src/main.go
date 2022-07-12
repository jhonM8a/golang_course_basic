package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

/**
La estructura de datos " Struct " tiene un método llamado " String " , que podemos sobrescribir para personalizar la salida a consola de los datos del struct.

*/
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %D GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)

}
func main() {
	myPc := pc{ram: 16, disk: 100, brand: "msi"}

	fmt.Println(myPc)
}
