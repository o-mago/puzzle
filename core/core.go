package core

import (
	"fmt"
	"math/rand"
)

var DEFAULT_TABLE = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
var KEYS = map[string]string{
	"up":    "w",
	"left":  "a",
	"down":  "s",
	"right": "d",
	"quit":  "q",
}

type Play struct {
	Table    [3][3]int
	Keys     map[string]string
	EmptyRow int
	EmptyCol int
}

func NewPlay() *Play {
	t, x, y := generateRandomTable()
	return &Play{
		Table:    t,
		Keys:     KEYS,
		EmptyRow: x,
		EmptyCol: y,
	}
}

func generateRandomTable() ([3][3]int, int, int) {
	t := DEFAULT_TABLE
	s := 3
	xEmpty := 0
	yEmpty := 0

	for i, r := range t {
		for j := range r {
			x := rand.Intn(s)
			y := rand.Intn(s)
			t[i][j], t[x][y] = t[x][y], t[i][j]

			if t[i][j] == 0 {
				xEmpty = i
				yEmpty = j
			}

			if t[x][y] == 0 {
				xEmpty = x
				yEmpty = y
			}
		}
	}

	return t, xEmpty, yEmpty
}

func (p *Play) IsWin() bool {
	return p.Table == DEFAULT_TABLE
}

func (p *Play) Up() error {
	if p.EmptyRow == 0 {
		return fmt.Errorf("can't move up")
	}

	p.Table[p.EmptyRow][p.EmptyCol], p.Table[p.EmptyRow-1][p.EmptyCol] = p.Table[p.EmptyRow-1][p.EmptyCol], p.Table[p.EmptyRow][p.EmptyCol]
	p.EmptyRow = p.EmptyRow - 1

	return nil
}

func (p *Play) Down() error {
	if p.EmptyRow == 2 {
		return fmt.Errorf("can't move down")
	}

	p.Table[p.EmptyRow][p.EmptyCol], p.Table[p.EmptyRow+1][p.EmptyCol] = p.Table[p.EmptyRow+1][p.EmptyCol], p.Table[p.EmptyRow][p.EmptyCol]
	p.EmptyRow = p.EmptyRow + 1

	return nil
}

func (p *Play) Left() error {
	if p.EmptyCol == 0 {
		return fmt.Errorf("can't move left")
	}

	p.Table[p.EmptyRow][p.EmptyCol], p.Table[p.EmptyRow][p.EmptyCol-1] = p.Table[p.EmptyRow][p.EmptyCol-1], p.Table[p.EmptyRow][p.EmptyCol]
	p.EmptyCol = p.EmptyCol - 1

	return nil
}

func (p *Play) Right() error {
	if p.EmptyCol == 2 {
		return fmt.Errorf("can't move right")
	}

	p.Table[p.EmptyRow][p.EmptyCol], p.Table[p.EmptyRow][p.EmptyCol+1] = p.Table[p.EmptyRow][p.EmptyCol+1], p.Table[p.EmptyRow][p.EmptyCol]
	p.EmptyCol = p.EmptyCol + 1

	return nil
}
