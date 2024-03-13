package main

import (
	"fmt"
	"os"
	"time"
)

type Game struct {
	done   *chan bool
	ticker *time.Ticker
	field  Field
	logs   []string
}

func NewGame(done *chan bool) *Game {
	game := Game{
		done:   done,
		ticker: time.NewTicker(time.Second),
		field: Field{
			Curr: Current{
				Id: 1,
				X:  W / 2,
				Y:  0,
			},
		},
	}

	go game.update()
	go game.inputs()

	return &game
}

func (g *Game) Draw() {
	lines := g.field.String()
	for i, line := range lines {
		val := "  " + line
		if len(g.logs) > i {
			val += "  - " + g.logs[i]
		}
		fmt.Println(val)
	}
}

func (g *Game) Redraw() {
	clear_lines(H + 2)
	g.Draw()
}

func (g *Game) update() {
	for {
		select {
		case <-g.ticker.C:
			g.moveDown()
		}
	}
}

func (g *Game) log(msg string) {
	if len(g.logs) >= 10 {
		g.logs = append(g.logs[1:], msg)
	} else {
		g.logs = append(g.logs, msg)
	}
}

// ESC  [
// 27   91 __
// - up    65
// - down  66
// - right 67
// - left  68
func (g *Game) inputs() {
	var b []byte = make([]byte, 3)
	for {
		os.Stdin.Read(b)
		if b[0] == 27 && b[1] == 91 {
			switch b[2] {
			case 65:
				g.rotate()
			case 66:
				g.moveDown()
			case 67:
				g.moveRight()
			case 68:
				g.moveLeft()
			}
		}
	}
}

func (g *Game) rotate() {
	if g.field.Collision(0, 0, 1) {
		return
	}
	g.field.Curr.R = nextRotation(g.field.Curr.R, 1)
	g.Redraw()
}

func (g *Game) moveDown() {
	if g.field.Collision(0, 1, 0) {
		// place block
		g.log("place block")
		g.field.Place()
		lines := g.field.DropLines()
		if lines > 0 {
			g.log(fmt.Sprintf("%d lines dropped", lines))
		}
		// Spawn next block
		g.field.Spawn()
		g.Redraw()
		if g.field.Collision(0, 1, 0) {
			g.End()
		}
	} else {
		// move
		g.log("move down")
		g.field.Curr.Y += 1
		g.Redraw()
	}
}

func (g *Game) moveLeft() {
	if g.field.Collision(-1, 0, 0) {
		return
	}
	g.log("move left")
	g.field.Curr.X -= 1
	g.Redraw()
}

func (g *Game) moveRight() {
	if g.field.Collision(1, 0, 0) {
		return
	}
	g.log("move right")
	g.field.Curr.X += 1
	g.Redraw()
}

func (g *Game) End() {
	// end game
	fmt.Println("gg")
	*g.done <- true
}
