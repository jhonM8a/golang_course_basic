package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

//	Libreria para imprimir

func main() {
	//switch como se encuentra generalmente
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Par")
	default:
		fmt.Println("impar")
	}

	//switch sin condición
	value := 200
	switch {
	case value > 100:
		fmt.Println("Menor a 100")
	case value < 0:
		fmt.Println("Menor a 0")
	default:
		fmt.Println("Sin condición")
	}
}
