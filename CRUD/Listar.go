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
		for _, libro := range libros {
			fmt.Printf("id: %d\n", libro.Id)
			fmt.Printf("Nombre: %s\n", libro.Nombre)
			fmt.Printf("Autor: %s\n", libro.Autor)
			fmt.Printf("Editorial: %s\n", libro.Editorial)
			fmt.Println("-----------------------------------")
		}
	}

	return err
}