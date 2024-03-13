package main

const (
	PcO int = 1
	PcS     = 2
	PcZ     = 3
	PcT     = 4
	PcL     = 5
	PcJ     = 6
	PcI     = 7
)

type Piece struct {
	Id    int
	Size  int
	Cells []bool
}

var pieceO = Piece{
	Id:   PcO,
	Size: 2,
	Cells: []bool{
		true, true,
		true, true,
	},
}
var pieceS = Piece{
	Id:   PcS,
	Size: 3,
	Cells: []bool{
		false, true, true,
		true, true, false,
		false, false, false,
	},
}
var pieceZ = Piece{
	Id:   PcZ,
	Size: 3,
	Cells: []bool{
		true, true, false,
		false, true, true,
		false, false, false,
	},
}
var pieceT = Piece{
	Id:   PcT,
	Size: 3,
	Cells: []bool{
		false, true, false,
		true, true, true,
		false, false, false,
	},
}
var pieceL = Piece{
	Id:   PcJ,
	Size: 3,
	Cells: []bool{
		false, false, true,
		true, true, true,
		false, false, false,
	},
}
var pieceJ = Piece{
	Id:   PcJ,
	Size: 3,
	Cells: []bool{
		true, false, false,
		true, true, true,
		false, false, false,
	},
}
var pieceI = Piece{
	Id:   PcI,
	Size: 4,
	Cells: []bool{
		false, false, false, false,
		true, true, true, true,
		false, false, false, false,
		false, false, false, false,
	},
}

var pieces = map[int]*Piece{
	PcO: &pieceO,
	PcS: &pieceS,
	PcZ: &pieceZ,
	PcT: &pieceT,
	PcL: &pieceL,
	PcJ: &pieceJ,
	PcI: &pieceI,
}

// r * 90 = deg
func (p *Piece) ToMatrix(r int) [][]bool {
	m := make([][]bool, p.Size)
	for y := 0; y < p.Size; y++ {
		row := make([]bool, p.Size)
		for x := 0; x < p.Size; x++ {
			row[x] = p.Cells[idxw(x, y, p.Size)]
		}
		m[y] = row
	}
	return rotate(m, r)
}

func rotate[T interface{}](m [][]T, r int) [][]T {
	switch r {
	case 1:
		// +90
		m = transpose(m)
		reverseRows(m)
	case 2:
		// +180
		m = rotate(rotate(m, 1), 1)
	case 3:
		// -90
		m = transpose(m)
		reverseCols(m)
	}
	return m
}

func transpose[T any](matrix [][]T) [][]T {
	w, h := len(matrix[0]), len(matrix)
	res := make([][]T, w)
	for i := range res {
		res[i] = make([]T, h)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			res[x][y] = matrix[y][x]
		}
	}
	return res
}

func reverseRows[T any](m [][]T) {
	for y := 0; y < len(m); y++ {
		for i, j := 0, len(m[y])-1; i < j; i, j = i+1, j-1 {
			m[y][i], m[y][j] = m[y][j], m[y][i]
		}
	}
}
func reverseCols[T any](m [][]T) {
	for x := 0; x < len(m[0]); x++ {
		for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
			m[i][x], m[j][x] = m[j][x], m[i][x]
		}
	}
}
