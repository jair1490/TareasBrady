package main
/*
import (
	"GoProject/BasedeDatos"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	_ "GoProject/BasedeDatos"
)

func main() {
	var c BasedeDatos.Libro
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese el Nombre del libro: ")
	if scanner.Scan() {
		c.Nombre = scanner.Text()
	}

	libros, err := BasedeDatos.BuscarLibro(c.Nombre)

	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
	} else {
		fmt.Println("======Lista de Libros=========")
		for _, Libro := range libros {
			fmt.Printf("Nombre: %s\n", Libro.Nombre)
			fmt.Printf("Dirección: %s\n", Libro.Autor)
			fmt.Printf("E-mail: %s\n", Libro.Editorial)
			fmt.Println("-----------------------------------")

		}
	}

}

 */