package main //Indica la carpeta donde esta, pero como no hay ponemos el main

import "fmt" //	Libreria para imprimir

func main() {
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------")
	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("-------")
	// for forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
