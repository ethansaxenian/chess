package game

import "github.com/ethansaxenian/chess/piece"

type testPlayer struct {
}

func (t testPlayer) GetMove(moves [][2]string) (string, string) {
	return "", ""
}

func (t testPlayer) State() map[string]any {
	return map[string]any{}
}

func (t testPlayer) ChoosePromotionPiece(square string) piece.Piece {
	return piece.Queen
}

func NewTestState(fen string) *State {
	s := LoadFEN(fen, testPlayer{}, testPlayer{})

	return s
}
