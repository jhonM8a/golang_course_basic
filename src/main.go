package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import (
	"fmt" //	Libreria para imprimir
	"strings"
)

func isPalindromo(text string) {
	var textRevers string
	for i := len(text) - 1; i >= 0; i-- {
		textRevers += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textRevers) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindromo("Ama")

}
