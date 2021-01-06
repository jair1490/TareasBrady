package CRUD


import (
	"GoProject/BasedeDatos"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MenuInsertar() (e error) {

	var c BasedeDatos.Libro
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("========Insertar un nuevo libro======= \n")

	fmt.Print("Ingrese el nombre del libro: ")
	if scanner.Scan() {
		c.Nombre = scanner.Text()
	}
	fmt.Print("Ingrese su Autor: ")
	if scanner.Scan() {
		c.Autor = scanner.Text()
	}

	fmt.Print("Ingrese su Editorial :")
	if scanner.Scan() {
		c.Editorial = scanner.Text()
	}

	/*
	libros, err := BasedeDatos.Buscar(c.Nombre)
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
	} else {
		fmt.Println("======Lista de Libros=========")
		for _, Libro := range libros {
			fmt.Printf("Nombre: %s\n", Libro.Nombre)
			fmt.Printf("Autor: %s\n", Libro.Autor)
			fmt.Printf("Editorial: %s\n", Libro.Editorial)
			fmt.Println("-----------------------------------")

		}

	 */

		//if (len(libros) <= 0) {

			var respuesta string

			fmt.Print("Estas seguro de insertar este libro (S/N) :")
			if scanner.Scan() {
				respuesta = scanner.Text()
			}

			if(respuesta == "S"){

				err := BasedeDatos.Insertar(c)
				if err != nil {
					fmt.Printf("Error insertando: %v", err)
				} else {
					fmt.Println("Insertado correctamente")
				}
			}

			//}else {
			//fmt.Printf("El registro ya existe en la base de datos: \n")
		//}

	//}

	return e
}


