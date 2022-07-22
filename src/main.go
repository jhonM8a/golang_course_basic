package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() //Esta linea se va a ejecutar hasta el final de la funcion, y de esta forma libera el gorutine del WaitGroup
	fmt.Println(text)
}

func main() { //La función main corre en una goruntime

	var wg sync.WaitGroup //El paquete sync permite interacturar de forma primitiva con las gorutine. Variable que acomula un conjunto de gorutines y los va liberando poco a poco

	fmt.Println("Hello") //say("Hello")
	wg.Add(1)            //Indicamos que vamos a agregar 1 Gorutine al WaitGroup para que espere su ejecucion antes de que la gurutine base (main) muera, y así le de tiempo a la siguiente gorutine de ejecutarse

	go say("World", &wg)

	wg.Wait() // Funcion del WaitGroup que sirve para decirle al gorutine principal (main) que espere hasta que todas las gorutine del WaitGroup finalicen, es decir, hasta que se ejecute 'defer wg.Done()' en cada una de las goroutines
}
