package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

func say(text string, c chan<- string) { //Es de buena practia indicar el tipo de canal que estamos recibiendo, si es de entrada o salida ya q mejora la eficiencia del código
	c <- text // Le pasamos al canal el texto
}

func main() {
	c := make(chan string, 1) //Con esto se crea el channel que recibira una gorutime que trabajara con string
	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c) // Sacamos el dato que tiene el canal

}
