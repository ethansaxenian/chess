package move

import (
	"fmt"
	"slices"
	"testing"

	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
	"github.com/stretchr/testify/assert"
)

var validKnightSquaresFromE4 = []string{"f6", "g5", "g3", "f2", "d2", "c3", "c5", "d6"}

func TestValidateKnightValid(t *testing.T) {
	for _, target := range validKnightSquaresFromE4 {
		assert.True(t, validateKnightMove("e4", target), target)
	}
}

func TestValidateKnightInvalid(t *testing.T) {
	for _, f := range board.Files {
		for _, r := range board.Ranks {
			target := string(f) + string(r)
			if !slices.Contains(validKnightSquaresFromE4, target) {
				assert.False(t, validateKnightMove("e4", target))
			}
		}
	}
}

func TestValidateWhitePawnValid(t *testing.T) {
	assert.True(t, validatePawnMove("a2", "a3", piece.White))
	assert.True(t, validatePawnMove("a2", "a4", piece.White))
	assert.True(t, validatePawnMove("a3", "a4", piece.White))
	assert.True(t, validatePawnMove("a7", "a8", piece.White))
	assert.True(t, validatePawnMove("d2", "c3", piece.White))
	assert.True(t, validatePawnMove("d2", "e3", piece.White))
}

func TestValidateWhitePawnInvalid(t *testing.T) {
	assert.False(t, validatePawnMove("a2", "a1", piece.White))
	assert.False(t, validatePawnMove("a2", "a2", piece.White))
	assert.False(t, validatePawnMove("a2", "a5", piece.White))
	assert.False(t, validatePawnMove("a3", "a5", piece.White))
	assert.False(t, validatePawnMove("d2", "c4", piece.White))
	assert.False(t, validatePawnMove("d2", "e4", piece.White))
	assert.False(t, validatePawnMove("d2", "b3", piece.White))
	assert.False(t, validatePawnMove("d2", "f3", piece.White))
}

func TestValidateBlackPawnValid(t *testing.T) {
	assert.True(t, validatePawnMove("a7", "a6", piece.Black))
	assert.True(t, validatePawnMove("a7", "a5", piece.Black))
	assert.True(t, validatePawnMove("a6", "a5", piece.Black))
	assert.True(t, validatePawnMove("a2", "a1", piece.Black))
}

func TestValidateBlackPawnInvalid(t *testing.T) {
	assert.False(t, validatePawnMove("a7", "a8", piece.Black))
	assert.False(t, validatePawnMove("a7", "a7", piece.Black))
	assert.False(t, validatePawnMove("a7", "a4", piece.Black))
	assert.False(t, validatePawnMove("a6", "a4", piece.Black))
}

var validKingSquaresFromE4 = []string{"e5", "f5", "f4", "f3", "e3", "d3", "d4", "d5"}

func TestValidateKingValid(t *testing.T) {
	for _, target := range validKingSquaresFromE4 {
		assert.True(t, validateKingMove("e4", target), target)
	}
}

func TestValidateKingInvalid(t *testing.T) {
	for _, f := range board.Files {
		for _, r := range board.Ranks {
			target := string(f) + string(r)
			if !slices.Contains(validKingSquaresFromE4, target) {
				assert.False(t, validateKingMove("e4", target), target)
			}
		}
	}
}

var validBishopSquaresFromE4 = []string{"f5", "g6", "h7", "f3", "g2", "h1", "d3", "c2", "b1", "d5", "c6", "b7", "a8"}

func TestValidateBishopValid(t *testing.T) {
	for _, target := range validBishopSquaresFromE4 {
		assert.True(t, validateBishopMove("e4", target), target)
	}
}

func TestValidateBishopInvalid(t *testing.T) {
	for _, f := range board.Files {
		for _, r := range board.Ranks {
			target := string(f) + string(r)
			if !slices.Contains(validBishopSquaresFromE4, target) {
				assert.False(t, validateBishopMove("e4", target), target)
			}
		}
	}
}

var validRookSquaresFromE4 = []string{"e5", "e6", "e7", "e8", "e3", "e2", "e1", "d4", "c4", "b4", "a4", "f4", "g4", "h4"}

func TestValidateRookValid(t *testing.T) {
	for _, target := range validRookSquaresFromE4 {
		assert.True(t, validateRookMove("e4", target), target)
	}
}

func TestValidateRookInvalid(t *testing.T) {
	for _, f := range board.Files {
		for _, r := range board.Ranks {
			target := string(f) + string(r)
			if !slices.Contains(validRookSquaresFromE4, target) {
				assert.False(t, validateRookMove("e4", target), target)
			}
		}
	}
}

func TestValidateQueenValid(t *testing.T) {
	for _, target := range slices.Concat(validBishopSquaresFromE4, validRookSquaresFromE4) {
		assert.True(t, validateQueenMove("e4", target), target)
	}
}

func TestValidateQueenInvalid(t *testing.T) {
	for _, f := range board.Files {
		for _, r := range board.Ranks {
			target := string(f) + string(r)
			if !slices.Contains(slices.Concat(validBishopSquaresFromE4, validRookSquaresFromE4), target) {
				assert.False(t, validateQueenMove("e4", target), target)
			}
		}
	}
}

func TestValidatePrecomputedMoves(t *testing.T) {
	for p, moves := range precomputedPieceMoves {
		for src, targetMoves := range moves {
			for _, target := range targetMoves {
				assert.True(t, validatePieceMove(p, src, target), fmt.Sprintf("piece: %v, src: %v, target: %v", p, src, target))
			}
		}
	}
}

func TestValidatePawnMoveWithBoardValid(t *testing.T) {
	b := board.LoadFEN("8/8/8/3ppp2/3PPP2/8/8/8 w - - 0 1")
	assert.True(t, validatePawnMoveWithBoard(b, "e4", "d5"), "e4 d5")
	assert.True(t, validatePawnMoveWithBoard(b, "e4", "f5"), "e4 f5")
	assert.True(t, validatePawnMoveWithBoard(b, "e5", "d4"), "e5 d4")
	assert.True(t, validatePawnMoveWithBoard(b, "e5", "f4"), "e5 f4")
}

func TestValidatePawnMoveWithBoardInvalid(t *testing.T) {
	b := board.LoadFEN("8/8/8/3ppp2/3PPP2/8/8/8 w - - 0 1")
	assert.False(t, validatePawnMoveWithBoard(b, "e4", "e5"), "e4 e5")
	assert.False(t, validatePawnMoveWithBoard(b, "d4", "c5"), "d4 c5")
	assert.False(t, validatePawnMoveWithBoard(b, "e5", "e4"), "e5 e4")
	assert.False(t, validatePawnMoveWithBoard(b, "d5", "c4"), "d5 c4")
}

func TestValidateBishopMoveWithBoardValid(t *testing.T) {
	b := board.LoadFEN("8/1p6/6P1/8/4B3/3P4/2p5/8 w - - 0 1")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "d5"), "e4 d5")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "c6"), "e4 c6")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "b7"), "e4 b7")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "f3"), "e4 f3")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "f5"), "e4 f5")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "g2"), "e4 g2")
	assert.True(t, validateBishopMoveWithBoard(b, "e4", "h1"), "e4 h1")
}

func TestValidateBishopMoveWithBoardInvalid(t *testing.T) {
	b := board.LoadFEN("8/1p6/6P1/8/4B3/3P4/2p5/8 w - - 0 1")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "g6"), "e4 g6")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "h7"), "e4 h7")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "a8"), "e4 a8")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "d3"), "e4 d3")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "c2"), "e4 c2")
	assert.False(t, validateBishopMoveWithBoard(b, "e4", "b1"), "e4 b1")
}

func TestValidateRookMoveWithBoardValid(t *testing.T) {
	b := board.LoadFEN("8/8/4P3/8/1Pp1R3/8/8/4p3 w - - 1 1")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "e5"), "e4 e5")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "d4"), "e4 d4")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "c4"), "e4 c4")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "f4"), "e4 f4")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "g4"), "e4 g4")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "h4"), "e4 h4")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "e3"), "e4 e3")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "e2"), "e4 e2")
	assert.True(t, validateRookMoveWithBoard(b, "e4", "e1"), "e4 e1")
}

func TestValidateRookMoveWithBoardInvalid(t *testing.T) {
	b := board.LoadFEN("8/8/4P3/8/1Pp1R3/8/8/4p3 w - - 1 1")
	assert.False(t, validateRookMoveWithBoard(b, "e4", "b4"), "e4 b4")
	assert.False(t, validateRookMoveWithBoard(b, "e4", "a4"), "e4 a4")
	assert.False(t, validateRookMoveWithBoard(b, "e4", "e6"), "e4 e6")
	assert.False(t, validateRookMoveWithBoard(b, "e4", "e7"), "e4 e7")
	assert.False(t, validateRookMoveWithBoard(b, "e4", "e8"), "e4 e8")

	b = board.LoadFEN("8/8/8/8/8/8/7P/7R w - - 1 1")
	assert.False(t, validateRookMoveWithBoard(b, "h1", "h6"), "h1 h6")
}
