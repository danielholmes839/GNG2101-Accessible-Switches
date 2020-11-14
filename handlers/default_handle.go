package handlers

// Triggers commands 1-4 of the handler
func defaultHandle(input chan int, h Handler) {
	for button := range input {
		switch button {
		case 1:
			h.Command1()
		case 2:
			h.Command2()
		case 3:
			h.Command3()
		case 4:
			h.Command4()
		}
	}
}
