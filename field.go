package main

import (
	"math/rand"
	"strings"
)

const W = 10
const H = 20

var COLORS = map[int]int{
	1: 1,
	2: 2,
	3: 4,
	4: 5,
	5: 6,
	6: 7,
	7: 8,
}

type Field struct {
	Curr  Current
	Cells [W * H]int
}

// r 0 ^
// r 1 >
// r 2 v
// r 3 <
type Current struct {
	Id int
	R  int
	X  int
	Y  int
}
type Point struct {
	X int
	Y int
}

func idx(x int, y int) int {
	return idxw(x, y, W)
}
func idxw(x int, y int, w int) int {
	return x + w*y
}

func contains(arr []Point, x int, y int) bool {
	for _, p := range arr {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func nextRotation(r int, offset int) int {
	next := r + offset
	if next > 3 {
		return 0
	}
	return next
}

func (c *Current) getPoints(offsetX int, offsetY int, offsetR int) []Point {
	p := pieces[c.Id]
	m := p.ToMatrix(nextRotation(c.R, offsetR))

	points := []Point{}

	anchor := (p.Size + p.Size%2) / 2

	for y := 0; y < p.Size; y++ {
		for x := 0; x < p.Size; x++ {
			if !m[y][x] {
				continue
			}
			x1 := c.X - anchor + x + offsetX
			y1 := c.Y - anchor + y + offsetY
			points = append(points, Point{
				X: x1,
				Y: y1,
			})
		}
	}

	return points
}

func (f *Field) Place() {
	for _, p := range f.Curr.getPoints(0, 0, 0) {
		i := idx(p.X, p.Y)
		f.Cells[i] = f.Curr.Id
	}
}

func (f *Field) DropLines() int {
	var lines []int
	// find lines to drop
	for y := 0; y < H; y++ {
		full := true
		for x := 0; x < W; x++ {
			if f.Cells[idx(x, y)] == 0 {
				full = false
			}
		}
		if full {
			lines = append(lines, idx(0, y))
		}
	}

	if len(lines) == 0 {
		return 0
	}
	// build new cells array
	cells := make([]int, len(lines)*W)
	last := 0
	for _, i := range lines {
		cells = append(cells, f.Cells[last:i]...)
		last = i + W
	}
	cells = append(cells, f.Cells[last:]...)
	// apply update
	for i, val := range cells {
		f.Cells[i] = val
	}

	return len(lines)
}

func (f *Field) Spawn() {
	id := 1 + rand.Intn(7)
	f.Curr = Current{
		Id: id,
		X:  W / 2,
		Y:  0,
	}
}

func (f *Field) Collision(offsetX int, offsetY int, offsetR int) bool {
	for _, p := range f.Curr.getPoints(offsetX, offsetY, offsetR) {
		if p.X < 0 || p.X >= W || p.Y >= H {
			return true
		} else if p.Y < 0 {
			continue
		}
		i := idx(p.X, p.Y)
		if f.Cells[i] != 0 {
			return true
		}
	}
	return false
}

func (f *Field) String() []string {
	lines := []string{}
	lines = append(lines, "┌"+strings.Repeat("─", W*2)+"┐")
	points := f.Curr.getPoints(0, 0, 0)
	for y := 0; y < H; y++ {
		value := "│"
		for x := 0; x < W; x++ {
			if contains(points, x, y) {
				value += block(COLORS[f.Curr.Id])
			} else {
				i := idx(x, y)
				val := f.Cells[i]
				if val == 0 {
					value += empty()
				} else {
					value += block(COLORS[val])
				}
			}
		}
		value += "│"
		lines = append(lines, value)
	}
	lines = append(lines, "└"+strings.Repeat("─", W*2)+"┘")
	return lines
}
