package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt" //	Libreria para imprimir

func main() {
	// Array
	var array [4]int // Es inmutable en su tamaño
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array)) //len -> Tamaño, cap ->Capacidad

	// slices
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])   //Imprime el valor de esa posición
	fmt.Println(slice[:3])  // Imprime los elementos hasta la posición 3. La posición 3 es exclusiva
	fmt.Println(slice[2:4]) // Imprime los elemtnos desde la posición 2 hasta la 4. La posición 2 es inclusiva y la 4 exclusiva
	fmt.Println(slice[4:])  // Imprime desde el elemento 4. La posición 4 es inclusiva

	// Apped
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Los 3 puntos indica que lo va a descomprimirt, es decir agrega cada elemento de forma independiente en la lista
	fmt.Println(slice)
}
