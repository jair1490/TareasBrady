package CRUD


import (
	"GoProject/BasedeDatos"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MenuElimnar() (e error) {

	libros, err := BasedeDatos.ObtenerLibro()
	if err != nil {
		fmt.Printf("Error obteniendo Libros: %v", err)
	} else {
		fmt.Println("======Lista de Libros=========")
		for _, Libro := range libros {
			fmt.Printf("Nombre: %s\n", Libro.Nombre)
			fmt.Printf("Autor: %s\n", Libro.Autor)
			fmt.Printf("Editorial: %s\n", Libro.Editorial)
			fmt.Println("-----------------------------------")
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	var c BasedeDatos.Libro

	fmt.Print("Ingrese el Nombre a eliminar: ")
	if scanner.Scan() {
		c.Nombre = scanner.Text()
	}

	var respuesta string

	fmt.Print("Estas seguro de eliminar este libro (S/N) :")
	if scanner.Scan() {
		respuesta = scanner.Text()
	}

	if (respuesta == "S") {

		nil := BasedeDatos.Eliminar(c)
		if err != nil {
			fmt.Printf("Error Eliminado: %v", err)
		} else {
			fmt.Println("Eliminado correctamente")
		}

	}

	return err
}

