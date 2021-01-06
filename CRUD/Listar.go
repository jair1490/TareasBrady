package CRUD

import (
	"GoProject/BasedeDatos"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MenuListar() (e error) {

	libros, err := BasedeDatos.ObtenerLibro()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
	} else {
		fmt.Println("======Lista de Libros=========")
		for _, Libro := range libros {
			//fmt.Printf("Idlibro: %d\n", Libro.Idlibro)
			fmt.Printf("Nombre: %s\n", Libro.Nombre)
			fmt.Printf("Autor: %s\n", Libro.Autor)
			fmt.Printf("Editorial: %s\n", Libro.Editorial)
			fmt.Println("-----------------------------------")
		}
	}

	return err
}