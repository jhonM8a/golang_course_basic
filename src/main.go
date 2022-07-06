package main //Indica la carpeta donde esta, pero como no hay ponemos el main

import "fmt" //	Libreria para imprimir

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Jhon")
	value := returnValue(3)
	fmt.Println("value: ", value)
	value1, value2 := doubleReturn(2) // usando _ le decimos a go que no nos interesa la variable que se retorna
	fmt.Println("value1, value2: ", value1, value2)

}
