package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import "fmt"

func main() {
	/**
	Nota: Los mapas son efectivos cuando tenemos dos o mas listas que se relacionan entre ellas,
	ya que con los maps tenemos acceso mediante llave valor y son más eficientes que manejar
	slices o arrays ya que de forma nativa el map implementa cocurrencia para poder interactuar
	entre las opeaciones que usemos con el map
	*/
	m := make(map[string]int) // Asi se delcara un mapa que tiene como llave un string y valor un int
	m["jose"] = 14
	m["pepito"] = 30
	fmt.Println(m)

	// Recorrer un map
	for i, valor := range m { // No imprime en un orden en especifico ya que cuando se añaden datos se hacen de manera concurrente
		fmt.Println(i, valor)
	}

	//Get un valor
	value := m["jose"] // Si no existe la llave te devuelve el mismo valor que tiene zero value, en este caso como es int un 0
	fmt.Println(value)

	//Get un valor con ok
	value2, ok := m["jose"] // El ok indica si la llave existe o no el diccionario
	fmt.Println(value2, ok)
}
