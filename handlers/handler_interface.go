package handlers

// Handler interface
type Handler interface {
	Handle(input <-chan int) // start the handler
}
