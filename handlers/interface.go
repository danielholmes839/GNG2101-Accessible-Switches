package handlers

// Handler interface
type Handler interface {
	Command1()
	Command2()
	Command3()
	Command4()
	Handle(input chan int) // start the handler
}
