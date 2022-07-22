package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import (
	"fmt"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func main() { //La función main corre en una goruntime

	say("Hello")
	/*
		Con la palabra go indica que la función corre en una gorutime, es decir de manera cocurrente.
		Que va a pasar, se va a imprimir la palabra Hello pero el World no. Esto, debido a que
		el main acaba su ejecución y mure la gorutime, El World se ejecuta en otro hilo diferente
		al del main.
	*/
	go say("World")

	/*
		De esta manera se puede hacer que se imprima el world en el mismo hiilo del main, pero no es recomendable
	*/
	time.Sleep(time.Second * 1)

}
