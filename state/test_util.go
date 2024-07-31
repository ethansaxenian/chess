package state

import (
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
)

type testPlayer struct {
}

func (t testPlayer) GetMove(moves []move.Move) move.Move {
	return move.Move{}
}

func (t testPlayer) State() map[string]any {
	return map[string]any{}
}

func (t testPlayer) ChoosePromotionPiece(square string) piece.Piece {
	return piece.Queen
}

func NewTestStateFromFEN(fen string) *State {
	s := StartingStateFromFEN(fen, testPlayer{}, testPlayer{})

	return s
}
