package main

import (
	"github.com/gosuri/uiprogress"
	"time"
)

func main() {
	uiprogress.Start() // Begin render

	myBar := uiprogress.AddBar(100) // Add bar to track progress
	myBar.AppendCompleted()         // print completion time upon completion
	myBar.PrependElapsed()          // print elapsed time before bar

	for myBar.Incr() {
		time.Sleep(time.Second / 20)
	}

}
