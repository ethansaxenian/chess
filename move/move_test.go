package move

import (
	"fmt"
	"slices"
	"testing"

	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
	"github.com/stretchr/testify/assert"
)

func TestValidateSquareValid(t *testing.T) {
	assert.True(t, validateBounds("a1"))
	assert.True(t, validateBounds("h8"))
	assert.True(t, validateBounds("d4"))
}

func TestValidateSquareInvalid(t *testing.T) {
	assert.False(t, validateBounds("a0"))
	assert.False(t, validateBounds("h9"))
	assert.False(t, validateBounds("i4"))
	assert.False(t, validateBounds("1a"))
	assert.False(t, validateBounds("a"))
	assert.False(t, validateBounds(""))
	assert.False(t, validateBounds("foo"))
}

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
}

func TestValidateWhitePawnInvalid(t *testing.T) {
	assert.False(t, validatePawnMove("a2", "a1", piece.White))
	assert.False(t, validatePawnMove("a2", "a2", piece.White))
	assert.False(t, validatePawnMove("a2", "a5", piece.White))
	assert.False(t, validatePawnMove("a3", "a5", piece.White))
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
			for target, valid := range targetMoves {
				assert.Equal(t, valid, validatePieceMovement(p, src, target), fmt.Sprintf("piece: %v, src: %v, target: %v, valid: %v", p, src, target, valid))
			}
		}
	}
}
