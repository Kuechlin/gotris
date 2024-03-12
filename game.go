package main

import (
	"fmt"
	"math/rand"
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
			Curr: Piece{
				Id: 1,
				X:  0,
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
				fmt.Println("up")
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

func (g *Game) moveDown() {
	if g.field.CollisionDown() {
		// place block
		i := idx(g.field.Curr.X, g.field.Curr.Y)
		g.field.Cells[i] = g.field.Curr.Id
		// Spawn next block
		g.Spawn()
	} else {
		// move
		g.log("move down")
		g.field.Curr.Y += 1
		g.Redraw()
	}
}

func (g *Game) moveLeft() {
	if g.field.CollisionLeft() {
		return
	}
	g.log("move left")
	g.field.Curr.X -= 1
	g.Redraw()
}

func (g *Game) moveRight() {
	if g.field.CollisionRight() {
		return
	}
	g.log("move right")
	g.field.Curr.X += 1
	g.Redraw()
}

func (g *Game) Spawn() {
	p := Piece{
		X:  W / 2,
		Y:  0,
		Id: 1 + rand.Intn(4),
	}
	i := idx(p.X, p.Y)

	if g.field.Cells[i] != 0 {
		// end game
		fmt.Println("gg")
		*g.done <- true
		return
	}

	g.log(fmt.Sprintf("span %d at %d, %d", p.Id, p.X, p.Y))
	g.field.Curr = p
	g.Redraw()
}
