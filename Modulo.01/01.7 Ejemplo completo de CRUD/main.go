// Seminario GoLang - 01.7 Ejemplo completo de CRUD
// https://youtu.be/KlIG_LrifHQ
package main

import "fmt"

func main() {
	//CRUD
	p := NewPlaylist()
	p.Add(Song{0, "song-0"})
	p.Add(Song{1, "song-1"})
	p.Add(Song{2, "song-2"})

	p.Print()
	s0 := p.FindByID(0)
	if s0 != nil {
		fmt.Println("Se encuentra el ID=0")
	}

	p.Delete(0)

	p.Print()

	p.Update(Song{1, "song1"})
	p.Print()
}
