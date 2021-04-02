package BasedeDatos

import (
	//"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
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
	Id        int	`gorm:"primary_key,AUTO_INCREMENT"`
	Nombre    string
	Autor     string
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

	var libros []Libro
	db.Find(&libros)
	
	return libros, nil
}


func Insertar(c Libro) (e error) {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(&c)
	if err != nil {
		return err
	}
	//_, err = sentenciaPreparada.Exec(c.Nombre, c.Autor, c.Editorial)
	if err != nil {
		return err
	}
	return nil
}

func Actualizar(c Libro) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Model (Libro {}).Where("Nombre =?", c.Nombre).Updates(Libro {Autor: c.Autor , Editorial: c.Editorial })

	if err != nil {
		return err
	}

	return err

}

func Eliminar(c Libro) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()


	db.Where("nombre = ?", c.Nombre).Delete(&c)
	if err != nil {
		return err
	}


	if err != nil {
		return err
	}
	return err
}

func Buscar(c Libro) ([]Libro, error) {
	libros := []Libro{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	db.Where("nombre = ?", c.Nombre).First(&libros)
	if err != nil {
		return nil, err
	}

	return libros, nil
}





/*

func BuscarLibro(c Libro) ([]Libro, error) {
	libros := []Libro{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	//filas, err := db.Query("SELECT Nombre,Autor,Editorial FROM libro where Nombre like  ?",nombre + "%")
	db.Where("nombre LIKE ?", c.Nombre).Find(&libros)

	if err != nil {
		return nil, err
	}

	//defer filas.Close()

	//var c Libro

	/*for filas.Next() {
		err = filas.Scan(&c.Nombre,&c.Autor,&c.Editorial)
		if err != nil {
			return nil, err
		}
		libros = append(libros, c)
	}

return libros, nil
}

 */
