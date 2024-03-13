package main

import (
	"os/exec"
	"time"
)

var ticker *time.Ticker

// vsp +term
func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	exec.Command("tput", "civis")

	done := make(chan bool)
	game := NewGame(&done)

	title()
	game.Draw()

	for {
		select {
		case <-done:
			return
		}
	}
}
