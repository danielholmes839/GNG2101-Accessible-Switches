package listeners

// Listener interface
type Listener interface {
	Listen(input chan<- int, delay int)
}
