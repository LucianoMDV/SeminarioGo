// Seminario GoLang - 01.4 Tipo Funcion
// https://youtu.be/LQxzHcN-Xoo
package main

import "fmt"

/**
* 1ra alternativa de hacer lo mismo
**/
// type close func()

/**
* 2ra y 3ra alternativa de hacer lo mismo
**/
type close func(s string)

func invokeClose(c close) {
	/**
	* 1ra alternativa de hacer lo mismo
	**/
	// c()

	/**
	* 2ra y 3ra alternativa de hacer lo mismo
	**/
	c("hello world2")
}
func main() {
	/**
	* 1ra alternativa de hacer lo mismo
	**/
	// f := func(s string) {
	// 	fmt.Println("hello world")
	// }
	// invokeClose(f)

	/**
	* 2da alternativa de hacer lo mismo
	**/
	f := func(s string) {
		fmt.Println(s)
	}
	invokeClose(f)

	/**
	* 3ra alternativa de hacer lo mismo
	**/
	invokeClose(func(s string) {
		fmt.Println(s)
	})
}
