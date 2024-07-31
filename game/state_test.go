package game

import (
	"strings"
	"testing"

	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
	"github.com/stretchr/testify/assert"
)

func TestHandleEnPassantAvailable(t *testing.T) {
	s := NewTestStateFromFEN(board.StartingFEN)
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
	s := NewTestStateFromFEN("8/8/8/3Pp3/8/8/8/8 w - e6 0 1")
	s.MakeMove("d5", "e6")
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)
	assert.Equal(t, piece.Pawn*piece.White, s.Board.Square("e6"))
	assert.Equal(t, piece.Empty, s.Board.Square("e5"))
	assert.Equal(t, "8/8/4P3/8/8/8/8/8 b - - 0 1", s.FEN())
}

func TestLoadFen(t *testing.T) {
	s := NewTestStateFromFEN(board.StartingFEN)

	assert.Equal(t, strings.Fields(board.StartingFEN)[0], s.Board.FEN())
	assert.Equal(t, piece.White, s.ActiveColor)
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.Black])
	assert.Equal(t, noEnPassantTarget, s.EnPassantTarget)
	assert.Equal(t, 0, s.HalfmoveClock)
	assert.Equal(t, 1, s.FullmoveNumber)
}

func TestToFen(t *testing.T) {
	s := NewTestStateFromFEN(board.StartingFEN)
	assert.Equal(t, board.StartingFEN, s.FEN())

	fen := "4k2r/6r1/8/8/8/8/3R4/R3K3 w Qk - 0 1"
	s = NewTestStateFromFEN(fen)
	assert.Equal(t, fen, s.FEN())
}

func TestHandlePromotion(t *testing.T) {
	s := NewTestStateFromFEN("8/4P3/8/8/8/8/4p3/8 w - - 0 1")

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

func TestHandleCastle(t *testing.T) {
	s := NewTestStateFromFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")

	s.handleCastle(newMove(*s, "e1", "g1"))
	assert.Equal(t, piece.Rook*piece.White, s.nextBoard.Square("f1"))

	s.handleCastle(newMove(*s, "e1", "c1"))
	assert.Equal(t, piece.Rook*piece.White, s.nextBoard.Square("d1"))

	s.handleCastle(newMove(*s, "e8", "g8"))
	assert.Equal(t, piece.Rook*piece.Black, s.nextBoard.Square("f8"))

	s.handleCastle(newMove(*s, "e8", "c8"))
	assert.Equal(t, piece.Rook*piece.Black, s.nextBoard.Square("d8"))
}

func TestHandleUpdateCastlingRights(t *testing.T) {
	fen := "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1"
	s := NewTestStateFromFEN(fen)

	s.handleUpdateCastlingRights(newMove(*s, "e1", "e2"))
	assert.Equal(t, [2]bool{false, false}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.handleUpdateCastlingRights(newMove(*s, "e8", "e7"))
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{false, false}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("a1", "a2")
	s.handleUpdateCastlingRights(newMove(*s, "a1", "a2"))
	assert.Equal(t, [2]bool{true, false}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("h1", "h2")
	s.handleUpdateCastlingRights(newMove(*s, "h1", "h2"))
	assert.Equal(t, [2]bool{false, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("a8", "a7")
	s.handleUpdateCastlingRights(newMove(*s, "a8", "a7"))
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, false}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("h8", "h7")
	s.handleUpdateCastlingRights(newMove(*s, "h8", "h7"))
	assert.Equal(t, [2]bool{true, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{false, true}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("a1", "a8")
	s.handleUpdateCastlingRights(newMove(*s, "a1", "a8"))
	assert.Equal(t, [2]bool{true, false}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{true, false}, s.Castling[piece.Black])

	s.LoadFEN(fen)
	s.nextBoard.MakeMove("h8", "h1")
	s.handleUpdateCastlingRights(newMove(*s, "h8", "h1"))
	assert.Equal(t, [2]bool{false, true}, s.Castling[piece.White])
	assert.Equal(t, [2]bool{false, true}, s.Castling[piece.Black])
}

func TestUndo(t *testing.T) {
	tests := map[string]struct {
		startingFEN string
		move        [2]string
	}{
		"Starting FEN": {
			startingFEN: board.StartingFEN,
			move:        [2]string{"e2", "e4"},
		},
		"castling": {
			startingFEN: "rn2kbnr/p1pq1p1p/b2p2p1/1p2p3/2P1P1P1/3B1P1P/PP1PN3/RNBQK2R w KQkq - 0 1",
			move:        [2]string{"a7", "a5"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewTestStateFromFEN(test.startingFEN)
			s.MakeMove(test.move[0], test.move[1])
			s.Undo()
			assert.Equal(t, test.startingFEN, s.FEN())
			assert.Equal(t, s.fens[len(s.fens)-1], s.FEN())
		})
	}
}

func TestUndoSecondTurn(t *testing.T) {
	s := NewTestStateFromFEN(board.StartingFEN)
	s.MakeMove("e2", "e4")
	s.MakeMove("e7", "e5")
	s.Undo()
	s.Undo()
	assert.Equal(t, s.fens[len(s.fens)-1], s.FEN())
	assert.Equal(t, board.StartingFEN, s.FEN())
}
