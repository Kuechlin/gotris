package main

import (
	"strings"
)

const W = 10
const H = 20

var COLORS = map[int]int{
	0: 0,
	1: 10,
	2: 20,
	3: 40,
	4: 50,
}

type Field struct {
	Curr  Piece
	Cells [W * H]int
}

type Piece struct {
	Id int
	X  int
	Y  int
}

func idx(x int, y int) int {
	return x + W*y
}

func (f *Field) CollisionDown() bool {
	y := f.Curr.Y + 1
	next := idx(f.Curr.X, y)
	return y >= H || f.Cells[next] != 0
}
func (f *Field) CollisionLeft() bool {
	x := f.Curr.X - 1
	next := idx(x, f.Curr.Y)
	return x < 0 || f.Cells[next] != 0
}
func (f *Field) CollisionRight() bool {
	x := f.Curr.X + 1
	next := idx(x, f.Curr.Y)
	return x >= W || f.Cells[next] != 0
}

func (f *Field) String() []string {
	lines := []string{}
	lines = append(lines, "┌"+strings.Repeat("─", W*2)+"┐")
	piece := idx(f.Curr.X, f.Curr.Y)
	for y := 0; y < H; y++ {
		value := "│"
		for x := 0; x < W; x++ {
			i := idx(x, y)
			if i == piece {
				value += block(COLORS[f.Curr.Id])
			} else {
				val := f.Cells[idx(x, y)]
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
