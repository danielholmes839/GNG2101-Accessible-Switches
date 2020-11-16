package clicker

// Mode interface. 
type Mode interface {
	Execute(clicker *Clicker, x int, y int)
	GetName() string
}