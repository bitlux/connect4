package board

import "fmt"

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
		s += "----+"
	}
	fmt.Println(s)
}

func printEmpty(cols int) {
	s := "|"
	for i := 0; i < cols; i++ {
		s += "     |"
	}
	fmt.Println(s)
}

// Print prints a Board.
func (b *Board) Print() {
	printSep(b.width)
	for h := b.height - 1; h >= 0; h-- {
		for i := 0; i < 3; i++ {
			fmt.Print("|")
			for _, cell := range b.cells[h] {
				fmt.Printf("%s%s%s%s|", cell, cell, cell, cell)
			}
			fmt.Println()
		}
		printSep(b.width)
	}
	for c := 0; c < b.width; c++ {
		fmt.Printf("  %d  ", c+1)
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
		fmt.Printf("checking (%d, %d)\n", h, col)

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
