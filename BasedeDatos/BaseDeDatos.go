package BasedeDatos

import (
	//"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func ObtenerBaseDeDatos() (*gorm.DB, error)  {
	usuario := "admin"
	pass := "B@pemos1490"
	host := "tcp(localhost)"
	nombreBaseDeDatos := "librerias"
	db, err := gorm.Open( "mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		fmt.Printf("Error al conectar el mysql: %v", err)
	}
	return db,nil
}

type Libro struct {
	//Idlibro int
	Nombre string
	Autor string
	Editorial string
}


func (libro *Libro) TableName() string {
	return "libro"
}

func ObtenerLibro() ([]Libro, error) {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var libros []Libro

	db.Find(&libros)

	log.Print(&libros)

	if err != nil {
		return nil, err
	}

	/*var c Libro

	for libros.Next() {
		err = var libros.Scan(&c.Nombre, &c.Autor, &c.Editorial)
		if err != nil {
			return nil, err
		}
		var libros = append(var libros, c)
	}*/

	return libros, nil
}

func Insertar(c Libro) (e error) {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(c)
	if err != nil {
		return err
	}
	//_, err = sentenciaPreparada.Exec(c.Nombre, c.Autor, c.Editorial)
	if err != nil {
		return err
	}
	return nil
}


/*
func BuscarLibro(nombre string) ([]Libro, error) {
	libros := []Libro{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	filas, err := db.Query("SELECT Nombre,Autor,Editorial FROM libro where Nombre like  ?",nombre + "%")

	if err != nil {
		return nil, err
	}

	defer filas.Close()

	var c Libro

	for filas.Next() {
		err = filas.Scan(&c.Nombre,&c.Autor,&c.Editorial)
		if err != nil {
			return nil, err
		}
		libros = append(libros, c)
	}

	return libros, nil
}

func Eliminar(nombre string) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Query("DELETE FROM libro WHERE Nombre = ?",nombre)
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	if err != nil {
		return err
	}
	return err
}

func Buscar(nombre string) ([]Libro, error) {
	libros := []Libro{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	filas, err := db.Query("SELECT Nombre,Autor,Editorial FROM libro where Nombre =  ?",nombre)

	if err != nil {
		return nil, err
	}

	defer filas.Close()

	var c Libro

	for filas.Next() {
		err = filas.Scan(&c.Nombre,&c.Autor,&c.Editorial)
		if err != nil {
			return nil, err
		}
		libros = append(libros, c)
	}

	return libros, nil
}

func Actualizar(c Libro) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE libro SET Autor = ?, Editorial = ? WHERE Nombre = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
		_, err = sentenciaPreparada.Exec(c.Autor, c.Editorial, c.Nombre)
	return err

}
 */
