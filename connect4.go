package main

import "github.com/bitlux/connect4/board"

func main() {

	b := board.New(6, 7)

	b.Move(1, 1)
	b.Print()

	b.Move(1, 2)
	b.Print()

	b.Move(1, 1)
	b.Print()

	b.Move(2, 2)
	b.Print()

}
