package main

import (
	"GoProject/CRUD"
	_ "GoProject/CRUD"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	creditos := `==========================================================
	
	LIBRERIA DEL MUÑECO PC GAMER LO ES TODO 
==========================================================`
	fmt.Println(creditos)
	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Listar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`

	var eleccion int
	var _ error
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		switch eleccion {
		case 1:
			_ = CRUD.MenuListar()
		case 2:
			_ = CRUD.MenuInsertar()

			/*

		case 3:
			_ = CRUD.MenuActualizar()
		case 4:
			_ = CRUD.MenuElimnar()
		*/
		}
	}

	var respuesta string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Estas seguro de salir (S/N) :")
	if scanner.Scan() {
		respuesta = scanner.Text()
	}

	if (respuesta == "S") {
	os.Exit(1)
	}

}