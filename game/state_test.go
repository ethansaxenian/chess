package game

import (
	"strings"
	"testing"

	"github.com/ethansaxenian/chess/piece"
	"github.com/stretchr/testify/assert"
)

func TestHandleEnPassantAvailable(t *testing.T) {
	s := NewTestState(StartingFEN)
	s.handleEnPassantAvailable(newMove(*s, "e2", "e4"))
	assert.Equal(t, "e3", s.EnPassantTarget)

	s.handleEnPassantAvailable(newMove(*s, "e7", "e5"))
	assert.Equal(t, "e6", s.EnPassantTarget)

	s.handleEnPassantAvailable(newMove(*s, "e2", "e3"))
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)

	s.handleEnPassantAvailable(newMove(*s, "e7", "e6"))
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)
}

func TestHandleEnPassantCapture(t *testing.T) {
	s := NewTestState("8/8/8/3Pp3/8/8/8/8 w - e6 0 1")
	s.MakeMove("d5", "e6")
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)
	assert.Equal(t, piece.Pawn*piece.White, s.Board.Square("e6"))
	assert.Equal(t, piece.None, s.Board.Square("e5"))
	assert.Equal(t, "8/8/4P3/8/8/8/8/8 w - - 0 1", s.FEN())
}

func TestLoadFen(t *testing.T) {
	s := NewTestState(StartingFEN)

	assert.Equal(t, strings.Fields(StartingFEN)[0], s.Board.FEN())
	assert.Equal(t, piece.White, s.ActiveColor)
	assert.Equal(t, [2]bool{true, true}, s.whiteCastling)
	assert.Equal(t, [2]bool{true, true}, s.blackCastling)
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)
	assert.Equal(t, 0, s.halfmoveClock)
	assert.Equal(t, 1, s.fullmoveNumber)
}

func TestToFen(t *testing.T) {
	s := NewTestState(StartingFEN)
	assert.Equal(t, StartingFEN, s.FEN())

	fen := "4k2r/6r1/8/8/8/8/3R4/R3K3 w Qk - 0 1"
	s = NewTestState(fen)
	assert.Equal(t, fen, s.FEN())
}

func TestHandlePromotion(t *testing.T) {
	s := NewTestState("8/4P3/8/8/8/8/4p3/8 w - - 0 1")
	s.handlePromotion(newMove(*s, "e7", "e8"))
	assert.Equal(t, piece.Queen*piece.White, s.nextBoard.Square("e8"))
	s.handlePromotion(newMove(*s, "e7", "d8"))
	assert.Equal(t, piece.Queen*piece.White, s.nextBoard.Square("d8"))
	s.handlePromotion(newMove(*s, "e7", "f8"))
	assert.Equal(t, piece.Queen*piece.White, s.nextBoard.Square("f8"))

	s.handlePromotion(newMove(*s, "e2", "e1"))
	assert.Equal(t, piece.Queen*piece.Black, s.nextBoard.Square("e1"))
	s.handlePromotion(newMove(*s, "e2", "d1"))
	assert.Equal(t, piece.Queen*piece.Black, s.nextBoard.Square("d1"))
	s.handlePromotion(newMove(*s, "e2", "f1"))
	assert.Equal(t, piece.Queen*piece.Black, s.nextBoard.Square("f1"))
}
