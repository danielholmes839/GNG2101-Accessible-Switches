package clicker

import "fmt"

// Clicker type
type Clicker struct {
	dragOriginX int
	dragOriginY int

	mode  int
	modes []Mode
}

// NewClicker constructor
func NewClicker() *Clicker {
	modes := []Mode{
		&Click{name: "Right Click", side: "right", double: false},
		&Click{name: "Left Click", side: "left", double: false},
		&Click{name: "Double Right Click", side: "right", double: true},
		&Click{name: "Double Left Click", side: "left", double: true},
		&SetOrigin{},
		&Drag{},
	}
	return &Clicker{dragOriginX: 0, dragOriginY: 0, mode: 0, modes: modes}
}

// Next func
func (c *Clicker) Next() {
	c.mode++
	c.mode %= len(c.modes)
	fmt.Println("Selected: ", c.modes[c.mode].GetName())
}

// Click func
func (c *Clicker) Click(x int, y int) {
	mode := c.modes[c.mode]
	mode.Execute(c, x, y)
	fmt.Println("Click: ", mode.GetName())
}

// SpecificClick func
func (c *Clicker) SpecificClick(x int, y int, s int) {
	mode := c.modes[s]
	mode.Execute(c, x, y)
	fmt.Println("Specific Click: ", mode.GetName())
}
