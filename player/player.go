package player

import "github.com/bitlux/connect4/board"

// Player represents a player.
type Player interface {
	Move(b *board.Board) int
}
