package handlers

import (
	"fmt"
)

// EmptyHandler struct
type EmptyHandler struct {
}

// Handle func
func (handler EmptyHandler) Handle(input <-chan int) {
	for value := range input {
		fmt.Println(value)
	}
}
