package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c)) //len --> Cantidad de gorutimes que hay en el channel y cap --> capacidad maxima de gorutimes que puede almacenar el channel

	// Close
	// Close le indica al runtime de go que ese canal no va a recibir ni un dato más, auque tenga la capacidad para hacerlo.
	close(c)

	// Range
	// Range ayuda a recorrer todos los mensajes que hay en ese channel
	for message := range c {
		fmt.Println(message)
	}

	// Select
	// Select nos ayuda cuando manejamos multiples canales y no sabemos cual va responde primero
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ { //Es necesario tener presente cuantos channel y mensajes vayamos a tener
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1:", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2:", m2)
		}
	}
}
