package chat

import (
	"fmt"

	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/config"
	"github.com/jmoiron/sqlx"
)

//Producto is a object...
type Producto struct {
	ID     int64
	Nombre string
	Precio int64
}

//Service is a interface...
type Service interface {
	AddProducto(Producto) (Producto, error)
	FindByID(int) (*Producto, error)
	FindAll() []*Producto
	RemoveByID(int) (bool, error)
	UpdateProducto(Producto) (bool, error)
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New is a...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddProducto(p Producto) (Producto, error) {
	sqlStatement := "INSERT INTO producto (nombre, precio) VALUES (?,?)"

	res, err := s.db.Exec(sqlStatement, p.Nombre, p.Precio)
	if err != nil {
		return p, err
	}

	p.ID, _ = res.LastInsertId()
	fmt.Println(res.LastInsertId())

	return p, err
}

func (s service) FindByID(ID int) (*Producto, error) {
	var p Producto
	sqlStatement := "SELECT * FROM producto WHERE ID=?"
	if err := s.db.Get(&p, sqlStatement, ID); err != nil {
		return nil, err
	}

	return &p, nil
}

func (s service) FindAll() []*Producto {
	var list []*Producto
	// list = append(list, &Producto{0, "Hello word"})
	if err := s.db.Select(&list, "SELECT * FROM producto"); err != nil {
		panic(err)
	}
	return list
}

func (s service) UpdateProducto(p Producto) (bool, error) {

	sqlStatement := "UPDATE producto SET nombre = ?, precio = ? WHERE ID = ?"

	_, err := s.db.Exec(sqlStatement, p.Nombre, p.Precio, p.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s service) RemoveByID(ID int) (bool, error) {
	sqlStatement := "DELETE FROM producto WHERE ID = ?"
	_, err := s.db.Exec(sqlStatement, ID)
	if err != nil {
		return false, err
	}

	return true, nil

}
