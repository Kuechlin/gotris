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
	level  int
	score  int
}

func NewGame(done *chan bool) *Game {
	game := Game{
		done:   done,
		ticker: time.NewTicker(getSpeed(0)),
		level:  0,
		score:  0,
		field:  Field{},
	}
	game.field.Spawn()
	game.field.Spawn()

	go game.update()
	go game.inputs()

	return &game
}

func (g *Game) Draw() {
	left := g.field.Display()
	right := []string{
		"# preview",
	}
	right = append(right, g.field.Preview()...)

	right = append(right,
		"",
		"level: "+colored(5, fmt.Sprint(g.level)),
		"score: "+colored(5, fmt.Sprint(g.score)),
	)

	// add logs
	right = append(right, "", "# logs")
	for _, l := range g.logs {
		right = append(right, "- "+l)
	}

	for i, line := range left {
		val := "  " + line
		if i < len(right) {
			val += "    " + right[i]
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

func (g *Game) scoreLines(count int) {
	switch count {
	case 1:
		g.score += 40
	case 2:
		g.score += 100
	case 3:
		g.score += 300
	case 4:
		g.score += 1200
	}
	g.log(fmt.Sprintf("%d lines dropped", count))
	next := getLevel(g.score)
	if next != g.level {
		g.level = next
		g.ticker.Reset(getSpeed(next))
		g.log("level up")
	}
}

func getLevel(score int) int {
	return score / 200
}

func getSpeed(level int) time.Duration {
	if level > 100 {
		return 100 * time.Millisecond
	}
	p := float64(100-level) / 100
	val := 900*p + 100
	return time.Duration(val) * time.Millisecond
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
			g.scoreLines(lines)
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
