// Seminario GoLang - Tipos e Interfaces (objetos)
// https://youtu.be/zvsKp4TCM6w
package main

import "fmt"

// printable...
type printable interface {
	print()
}

type person struct {
	name string
}

type figure struct {
	name string
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name)
}

// *person es un puntero
func (p *person) print() {
	fmt.Println("[person]", p.name)
}

// func (p *person) clean() {
// 	p.name = ""
// }

func invokePrint(p printable) {
	p.print()
}

func main() {
	p := &person{name: "Lucho"}
	// p.print()
	// p.clean()
	// p.print()
	invokePrint(p)

	f := &figure{name: "Cube"}
	invokePrint(f)
}
