package main

import "fmt"

// Playlist is a object..
type Playlist struct {
	songs map[int]*Song
}

// Song is a object...
type Song struct {
	ID   int
	Name string
}

// NewPlaylist is a constructor...
func NewPlaylist() Playlist {
	songs := make(map[int]*Song)
	return Playlist{
		songs,
	}
}

// Add is a method...
func (p Playlist) Add(s Song) {
	p.songs[s.ID] = &s
}

// Print is a method...
func (p Playlist) Print() {
	for _, v := range p.songs {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Name)
	}
}

// FindByID is a method...
func (p Playlist) FindByID(ID int) *Song {
	return p.songs[ID]
}

// Delete is a method...
func (p Playlist) Delete(ID int) {
	delete(p.songs, ID)
}

// Update is a method...
func (p Playlist) Update(s Song) {
	p.songs[s.ID] = &s
}
