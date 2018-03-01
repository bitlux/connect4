package player

import "github.com/bitlux/connect4/board"

type Player interface {
	Move(b board.Board) int
}