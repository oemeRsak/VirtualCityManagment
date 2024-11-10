package main

// Main struct
type veheicle struct {
	id            int
	name          string
	typ           veheicle_typ
	priority      int
	position      [2]int
	destionation  [2]int
	direction     string
	communication chan string
}

// Main type struct
type veheicle_typ struct {
	id       int
	name     string
	priority int
	category string
}
