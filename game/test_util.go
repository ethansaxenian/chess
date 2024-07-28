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

func NewTestStateFromFEN(fen string) *State {
	s := StartingState(testPlayer{}, testPlayer{})
	s.LoadFEN(fen)

	return s
}
