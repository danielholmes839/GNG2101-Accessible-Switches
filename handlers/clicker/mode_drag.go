package clicker

import "github.com/go-vgo/robotgo"

// SetOrigin struct
type Drag struct {
}

// Execute func
func (save *Drag) Execute(c *Clicker, x int, y int) {
	robotgo.Move(c.dragOriginX, c.dragOriginY)
	robotgo.DragMouse(x, y)
}

// GetName func
func (save *Drag) GetName() string {
	return "Drag & Drop"
}
