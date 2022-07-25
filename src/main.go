package main //Indica la carpeta donde esta, pero como no hay ponemos el main
import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//Instanciar echo
	e := echo.New()

	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
