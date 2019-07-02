package main

import (
	"fmt"
	"time"
)

const (
	BASE = "["
	FILL = "="
	//CAP  = ">"
	FIN = "] ... COMPLETE!"
)

func main() {
	// initialize loop for testing terminal output
	x := 0
	// run loop for x
	for x <= 100 {
		switch x {
		case 0:
			fmt.Print(BASE)
		case 100:
			fmt.Print(FIN)
		default:
			fmt.Printf(FILL)
		}
		time.Sleep(time.Second / 25)
		x += 1
	}
}
