package main

import "testing"

func TestPLaylistAdd(t *testing.T) {
	p := NewPlaylist()
	s0 := p.FindByID(0)
	if s0 != nil {
		t.Error("La cancion con ID=0 ya existe")
	}

	p.Add(Song{0, "song-0"})
	s0 = p.FindByID(0)

	if s0 == nil {
		t.Error("La cancion con ID=0 no fue agregada")
	}

	if s0.Name != "song-0" {
		t.Error("La cancion con ID=0 no tiene la cancion correcta")

	}

}
