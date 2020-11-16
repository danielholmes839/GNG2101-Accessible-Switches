package clicker

// SetOrigin struct
type SetOrigin struct {
}

// Execute func
func (save *SetOrigin) Execute(clicker *Clicker, x int, y int) {
	clicker.dragOriginX = x
	clicker.dragOriginY = y
}

// GetName func
func (save *SetOrigin) GetName() string {
	return "Set Drag & Drop Origin"
}