package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

//	Libreria para imprimir

func main() {
	//switch nomal
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Par")
	default:
		fmt.Println("impar")
	}
}
