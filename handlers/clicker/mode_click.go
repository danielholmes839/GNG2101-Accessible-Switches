package clicker

import "github.com/go-vgo/robotgo"

// Click struct
type Click struct {
	name   string
	side   string
	double bool
}

// Execute func
func (click *Click) Execute(clicker *Clicker, x int, y int) {
	robotgo.MoveMouse(x, y)
	robotgo.MouseClick(click.side, click.double)
}

// GetName func
func (click *Click) GetName() string {
	return click.name
}
