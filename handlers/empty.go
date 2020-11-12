package handlers

import (
	"fmt"
)

// EmptyHandler struct
type EmptyHandler struct {
}

// Command1 func
func (handler EmptyHandler) Command1() {
	fmt.Println("Command 1")
}

// Command2 func
func (handler EmptyHandler) Command2() {
	fmt.Println("Command 2")
}

// Command3 func
func (handler EmptyHandler) Command3() {
	fmt.Println("Command 3")
}

// Command4 func
func (handler EmptyHandler) Command4() {
	fmt.Println("Command 4")
}

// Start func
func (handler EmptyHandler) Start() {
	fmt.Println("Start")
}