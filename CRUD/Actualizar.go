package CRUD

import (
	"GoProject/BasedeDatos"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MenuActualizar() (e error){

	var c BasedeDatos.Libro
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa el nombre del libro actualizar:")
	if scanner.Scan() {
		c.Nombre = scanner.Text()
	}

	fmt.Println("Ingresa el nuevo autor:")
	if scanner.Scan() {
		c.Autor = scanner.Text()
	}
	fmt.Println("Ingresa la nueva editorial:")
	if scanner.Scan() {
		c.Editorial = scanner.Text()
	}

	var respuesta string

	fmt.Print("Estas seguro de actualizar la informacion de este libro (S/N) :")
	if scanner.Scan() {
		respuesta = scanner.Text()
	}



	if(respuesta == "S"){

		err := BasedeDatos.Actualizar(c)
		if err != nil {
			fmt.Printf("Error actualizando: %v", err)
		} else {
			fmt.Println("Actualizado correctamente")
		}

	}


	return e
}

