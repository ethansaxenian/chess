package state

import (
	"testing"

	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
	"github.com/stretchr/testify/assert"
)

func TestGetMoveContext(t *testing.T) {
	tests := map[string]struct {
		startingFEN string
		move        move.Move
		expected    moveContext
	}{
		"white move nothing": {
			startingFEN: board.StartingFEN,
			move:        move.NewMove("e2", "e3"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				isPromotion:         false,
				castling:            nil,
			},
		},
		"black move nothing": {
			startingFEN: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1",
			move:        move.NewMove("e7", "e6"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         false,
			},
		},
		"white double pawn move": {
			startingFEN: board.StartingFEN,
			move:        move.NewMove("e2", "e4"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: "e3",
				castling:            nil,
				isPromotion:         false,
			},
		},
		"black double pawn move": {
			startingFEN: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1",
			move:        move.NewMove("e7", "e5"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: "e6",
				castling:            nil,
				isPromotion:         false,
			},
		},
		"white en passant capture": {
			startingFEN: "8/8/8/3pP3/8/8/8/8 w - d6 0 1",
			move:        move.NewMove("e5", "d6"),
			expected: moveContext{
				isCapture:           true,
				enPassantCapture:    "d5",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         false,
			},
		},
		"black en passant capture": {
			startingFEN: "8/8/8/8/3pP3/8/8/8 b - e3 0 1",
			move:        move.NewMove("d4", "e3"),
			expected: moveContext{
				isCapture:           true,
				enPassantCapture:    "e4",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         false,
			},
		},
		"white kingside castle": {
			startingFEN: "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1",
			move:        move.NewMove("e1", "g1"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            &struct{ side piece.Side }{piece.Kingside},
				isPromotion:         false,
			},
		},
		"white queenside castle": {
			startingFEN: "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1",
			move:        move.NewMove("e1", "c1"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            &struct{ side piece.Side }{piece.Queenside},
				isPromotion:         false,
			},
		},
		"black kingside castle": {
			startingFEN: "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1",
			move:        move.NewMove("e8", "g8"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            &struct{ side piece.Side }{piece.Kingside},
				isPromotion:         false,
			},
		},
		"black queenside castle": {
			startingFEN: "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1",
			move:        move.NewMove("e8", "c8"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            &struct{ side piece.Side }{piece.Queenside},
				isPromotion:         false,
			},
		},
		"white promotion": {
			startingFEN: "8/7P/8/8/8/8/p7/8 w - - 0 1",
			move:        move.NewMove("h7", "h8"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         true,
			},
		},
		"black promotion": {
			startingFEN: "8/7P/8/8/8/8/p7/8 b - - 0 1",
			move:        move.NewMove("a2", "a1"),
			expected: moveContext{
				isCapture:           false,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         true,
			},
		},
		"white promotion + capture": {
			startingFEN: "6n1/7P/8/8/8/8/p7/1N6 w - - 0 1",
			move:        move.NewMove("h7", "g8"),
			expected: moveContext{
				isCapture:           true,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         true,
			},
		},
		"black promotion + capture": {
			startingFEN: "6n1/7P/8/8/8/8/p7/1N6 w - - 0 1",
			move:        move.NewMove("a2", "b1"),
			expected: moveContext{
				isCapture:           true,
				enPassantCapture:    "",
				nextEnPassantTarget: noEnPassantTarget,
				castling:            nil,
				isPromotion:         true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewTestStateFromFEN(test.startingFEN)
			mc := getMoveContext(*s, test.move)
			assert.Equal(t, test.expected, mc)
		})
	}
}
