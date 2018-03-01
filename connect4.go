package main

import (
	"log"

	"github.com/bitlux/connect4/board"
	"github.com/bitlux/connect4/player"
)

func main() {

	b := board.New(6, 7)

	players := []player.Player{player.Human{},
		player.Human{}}

	for i := 0; i < 10; i++ {
		b.Print()

		p := players[i%2]
		move := p.Move(b)
		err := b.Move(move, i%2+1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
