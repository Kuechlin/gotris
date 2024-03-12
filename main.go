package main

import (
	"fmt"
	"os/exec"
	"time"
)

var ticker *time.Ticker

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	exec.Command("tput", "civis")

	done := make(chan bool)
	game := NewGame(&done)
	fmt.Println("> gotris")
	game.Draw()

	for {
		select {
		case <-done:
			return
		}
	}
}
