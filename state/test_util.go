package state

import (
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

type testPlayer struct {
}

func (t testPlayer) GetMove(moves []move.Move) move.Move {
	return move.Move{}
}

func (t testPlayer) ChoosePromotionPiece(square string) piece.Piece {
	return piece.Queen
}

func (t testPlayer) IsBot() bool {
	return true
}

func NewTestStateFromFEN(fen string) *State {
	s := StartingStateFromFEN(fen, testPlayer{}, testPlayer{})
	s.headless = true

	return s
}

func NewStartingTestState(white, black player.Player) *State {
	s := StartingState(white, black)
	s.headless = true

	return s
}
