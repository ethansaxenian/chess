package state

import (
	"fmt"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
)

type side int

const (
	kingside side = iota
	queenside
)

type moveContext struct {
	castling            *struct{ side side }
	enPassantCapture    string
	nextEnPassantTarget string
	isCapture           bool
	isPromotion         bool
}

func getMoveContext(s State, m move.Move) moveContext {
	var mc moveContext

	isPawn := s.Piece(m.Source).Type() == piece.Pawn
	sourceColor := s.Piece(m.Source).Color()

	if s.Piece(m.Target) != piece.Empty {
		mc.isCapture = true
	}

	if isPawn && m.Target == s.EnPassantTarget {
		capturedSquare := board.AddRank(m.Target, int(sourceColor)*-1)
		mc.isCapture = true
		mc.enPassantCapture = capturedSquare

		assert.Assert(
			s.Piece(capturedSquare).Type() == piece.Pawn && s.Piece(capturedSquare).Color() == s.Piece(m.Source).Color()*-1,
			fmt.Sprintf("getMoveContext: invalid en passant capture: %s %s", m.Source, m.Target),
		)
	}

	if isPawn && m.TargetRank() == piece.MaxPawnRank[sourceColor] {
		mc.isPromotion = true
	}

	if isPawn && (m.TargetRank()-m.SourceRank())*int(sourceColor) == 2 {
		enPassantSquare := board.CoordsToSquare(int(m.TargetFile()), m.TargetRank()-int(sourceColor))
		mc.nextEnPassantTarget = enPassantSquare
	} else {
		mc.nextEnPassantTarget = noEnPassantTarget
	}

	if s.Piece(m.Source).Type() == piece.King && m.Source == piece.StartingKingSquares[sourceColor] {
		switch m.TargetFile() {
		case 'g':
			mc.castling = &struct{ side side }{kingside}
		case 'b':
			mc.castling = &struct{ side side }{queenside}
		}

		assert.Assert(
			mc.enPassantCapture == "" || mc.nextEnPassantTarget == noEnPassantTarget,
			"a double pawn move cannot also be an en passant capture!",
		)
	}

	return mc
}
