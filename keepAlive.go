package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		// Simulate a keyboard combination: Ctrl + Shift + P
		// robotgo.KeyToggle("ctrl", "down")
		// robotgo.KeyToggle("shift", "down")
		// robotgo.KeyTap("p")
		// robotgo.KeyToggle("shift", "up")
		// robotgo.KeyToggle("ctrl", "up")

		// Simulate a left mouse click
		robotgo.MoveRelative(120, 0)
		robotgo.MoveRelative(0, -120)
		robotgo.MoveRelative(-120, 0)
		robotgo.MoveRelative(0, 120)

		// Sleep for a short duration to prevent high CPU usage
		time.Sleep(20 * time.Second)
	}
}
