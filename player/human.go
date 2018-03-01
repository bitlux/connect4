package player

import (
	"fmt"

	"github.com/bitlux/connect4/board"
)

// Human is a human.
type Human struct{}

// Move satisfies Player.
func (Human) Move(*board.Board) int {
	fmt.Print("> ")
	var m int
	fmt.Scanf("%d", &m)
	return m
}
