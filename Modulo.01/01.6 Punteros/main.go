// Seminario GoLang - 01.6 Punteros
// https://youtu.be/qDyyZ5Ji1AE
// Repo con diagramas y codigo: https://github.com/juanpablopizarro/tudai-pointers/tree/master
package main

import "fmt"

// Person is...
type Person struct {
	Name string
}

// & devuelve la direccion de memoria de una variable
// * a nivel definicion de funcnion indica que se requiere una posicion de memoria
// * a nivel operador hace referencia al valor que contiene el puntero
func (x *Person) changeName(name string) {
	x.Name = name
	fmt.Printf("%p %v\n", &x, x)
}

func main() {
	p := Person{Name: "Juan"}
	fmt.Printf("%p %v\n", &p, p)
	p.changeName("Pedro")
	fmt.Printf("%p %v\n", &p, p)
}
