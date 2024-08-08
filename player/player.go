package player

import (
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
)

type Player interface {
	GetMove([]move.Move) move.Move
	State() map[string]any
	ChoosePromotionPiece(string) piece.Piece
	IsBot() bool
}
