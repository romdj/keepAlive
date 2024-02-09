package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		// Simulate a keyboard combination: Ctrl + Shift + P
		robotgo.KeyToggle("ctrl", "down")
		robotgo.KeyToggle("shift", "down")
		robotgo.KeyTap("p")
		robotgo.KeyToggle("shift", "up")
		robotgo.KeyToggle("ctrl", "up")

		// Simulate a left mouse click
		robotgo.Click("left", false)

		// Sleep for a short duration to prevent high CPU usage
		time.Sleep(90 * time.Second)
	}
}
