package player

import (
	"github.com/ethansaxenian/chess/piece"
)

type Player interface {
	GetMove([][2]string) (string, string)
	State() map[string]any
	ChoosePromotionPiece(string) piece.Piece
}
