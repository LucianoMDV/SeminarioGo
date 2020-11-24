// Seminario GoLang - 01.2 Estructuras de datos (Array, Slices y Maps)
// https://youtu.be/8dJIf7QE-JM
package main

import "fmt"

func main() {
	fmt.Println("-----Array-----")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("-----Slices-----")

	// var l []int
	l := make([]int, 100)
	l = append(l, 10)
	fmt.Printf("%p\n", l)
	l = append(l, 100)
	fmt.Printf("%p\n", l)
	l = append(l, 1000)
	fmt.Printf("%p\n", l)
	for i, v := range l {
		fmt.Println(i, v)
	}

	fmt.Println("-----Maps-----")

	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	m[0] = "c"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
