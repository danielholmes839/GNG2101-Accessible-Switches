package handlers

import (
	"fmt"
)

// Handler interface
type Handler interface {
	Command1()
	Command2()
	Command3()
	Command4()
}

// DefaultHandler struct
type DefaultHandler struct {
}

// Command1 func
func (handler DefaultHandler) Command1() {
	fmt.Println("Command 1")
}

// Command2 func
func (handler DefaultHandler) Command2() {
	fmt.Println("Command 2")
}

// Command3 func
func (handler DefaultHandler) Command3() {
	fmt.Println("Command 3")
}

// Command4 func
func (handler DefaultHandler) Command4() {
	fmt.Println("Command 4")
}
