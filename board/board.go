package board

import (
	"fmt"
	"os"
	"os/exec"
)

type cell int

func (c cell) String() string {
	if c == 1 {
		return "#"
	}
	if c == 2 {
		return "."
	}
	return " "
}

// Board is a board
type Board struct {
	height, width int
	cells         [][]cell
}

// New creates a new Board.
func New(height, width int) *Board {
	b := &Board{height: height, width: width}
	for i := 0; i < height; i++ {
		b.cells = append(b.cells, make([]cell, width))
	}
	return b
}

func printSep(cols int) {
	s := "+"
	for i := 0; i < cols; i++ {
		s += "-----+"
	}
	fmt.Println(s)
}

func printRow(cells []cell, printSpace bool) {
	fmt.Print("|")
	for _, cell := range cells {
		c := cell.String()
		if printSpace {
			c = " "
		}
		fmt.Printf("%s%s%s%s%s|", c, cell, cell, cell, c)
	}
	fmt.Println()
}

// Print prints a Board.
func (b *Board) Print() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	printSep(b.width)
	for h := b.height - 1; h >= 0; h-- {
		printRow(b.cells[h], true)
		printRow(b.cells[h], false)
		printRow(b.cells[h], false)
		printRow(b.cells[h], true)
		printSep(b.width)
	}
	for c := 0; c < b.width; c++ {
		fmt.Printf("   %d  ", c+1)
	}
	fmt.Println()
}

// Move updates a Board with a move.
func (b *Board) Move(column, player int) error {
	col := column - 1
	if col < 0 || col >= b.width {
		return fmt.Errorf("move out of bounds: %d > %d", column, b.width+1)
	}

	made := false
	for h := 0; h < b.height; h++ {
		if b.cells[h][col] != 0 {
			continue
		}
		b.cells[h][col] = cell(player)
		made = true
		break
	}

	if !made {
		return fmt.Errorf("column %d is full", column)
	}
	return nil
}
