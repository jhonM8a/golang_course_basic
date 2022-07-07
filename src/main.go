package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

//	Libreria para imprimir

func main() {
	// Defer
	defer fmt.Println("Hola") // Antes de morir el código va ejecutar esta línea es como un finally en java
	fmt.Println("Mundo")

	// Continue se usa cuando algo nos interese que continue, por ejemplo cuando ocurre un error pero tu quieres seguir con el ciclo
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Es 8")
			break

		}
	}

}
