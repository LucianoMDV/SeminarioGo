// Seminario GoLang - 01.1 Introducción al lenguaje
// https://youtu.be/34Lj004wtzU
package main

import (
	"errors"
	"fmt"
)

func suma(a, b int) (int, error) {
	if a < b {
		return a, errors.New("el primer valor es menor que el segundo valor")
	}
	return a + b, nil
}

//Suma es una funcion publica y hay que documentarla
/**
 * esto es otra manera de hacer funciones y ademas como
 * seria si quisiera que sea publica va con las primer letra
 * en mayuscula
**/
func Suma(a, b int) (c int, e error) {
	c = a + b
	e = nil
	if a < b {
		e = errors.New("el primer valor es menor que el segundo valor")
	}

	return
}

var num1 = 4
var num2 = 5

func main() {
	fmt.Println("hello, world")
	r, err := Suma(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
