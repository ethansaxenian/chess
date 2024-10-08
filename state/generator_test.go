package state

import (
	"fmt"
	"slices"
	"testing"

	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
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

	assert.True(t, validateKingMove("e1", "f1"), "e1 f1")
	assert.True(t, validateKingMove("e1", "d1"), "e1 d1")
	assert.True(t, validateKingMove("e8", "f8"), "e8 f8")
	assert.True(t, validateKingMove("e8", "d8"), "e8 d8")
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

func TestValidatePawnMoveWithStateValid(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/8/3ppp2/3PPP2/8/8/8 w - - 0 1")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e4", "d5")), "e4 d5")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e4", "f5")), "e4 f5")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e5", "d4")), "e5 d4")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e5", "f4")), "e5 f4")
}

func TestValidatePawnMoveWithStateInvalid(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/8/3ppp2/3PPP2/8/8/8 w - - 0 1")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("e4", "e5")), "e4 e5")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d4", "c5")), "d4 c5")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("e5", "e4")), "e5 e4")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d5", "c4")), "d5 c4")

	s.LoadFEN("8/p7/n7/8/8/8/8/8 b - - 0 1")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("a7", "a5")), "a7 a5")
}

func TestValidateBishopMoveWithStateValid(t *testing.T) {
	s := *NewTestStateFromFEN("8/1p6/6P1/8/4B3/3P4/2p5/8 w - - 0 1")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "d5")), "e4 d5")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "c6")), "e4 c6")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "b7")), "e4 b7")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "f3")), "e4 f3")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "f5")), "e4 f5")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "g2")), "e4 g2")
	assert.True(t, validateBishopMoveWithState(s, move.NewMove("e4", "h1")), "e4 h1")
}

func TestValidateBishopMoveWithStateInvalid(t *testing.T) {
	s := *NewTestStateFromFEN("8/1p6/6P1/8/4B3/3P4/2p5/8 w - - 0 1")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "g6")), "e4 g6")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "h7")), "e4 h7")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "a8")), "e4 a8")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "d3")), "e4 d3")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "c2")), "e4 c2")
	assert.False(t, validateBishopMoveWithState(s, move.NewMove("e4", "b1")), "e4 b1")
}

func TestValidateRookMoveWithStateValid(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/4P3/8/1Pp1R3/8/8/4p3 w - - 1 1")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "e5")), "e4 e5")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "d4")), "e4 d4")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "c4")), "e4 c4")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "f4")), "e4 f4")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "g4")), "e4 g4")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "h4")), "e4 h4")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "e3")), "e4 e3")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "e2")), "e4 e2")
	assert.True(t, validateRookMoveWithState(s, move.NewMove("e4", "e1")), "e4 e1")
}

func TestValidateRookMoveWithStateInvalid(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/4P3/8/1Pp1R3/8/8/4p3 w - - 1 1")
	assert.False(t, validateRookMoveWithState(s, move.NewMove("e4", "b4")), "e4 b4")
	assert.False(t, validateRookMoveWithState(s, move.NewMove("e4", "a4")), "e4 a4")
	assert.False(t, validateRookMoveWithState(s, move.NewMove("e4", "e6")), "e4 e6")
	assert.False(t, validateRookMoveWithState(s, move.NewMove("e4", "e7")), "e4 e7")
	assert.False(t, validateRookMoveWithState(s, move.NewMove("e4", "e8")), "e4 e8")
}

func TestValidateQueenMoveWithStateValid(t *testing.T) {
	s := *NewTestStateFromFEN("p7/1P2P3/8/5p2/1Pp1Q2p/8/2P3P1/1P6 w - - 1 1")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "d5")), "e4 d5")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "c6")), "e4 c6")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "e5")), "e4 e5")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "e6")), "e4 e6")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "f5")), "e4 f5")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "f4")), "e4 f4")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "g4")), "e4 g4")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "h4")), "e4 h4")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "f3")), "e4 f3")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "e3")), "e4 e3")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "e2")), "e4 e2")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "e1")), "e4 e1")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "d3")), "e4 d3")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "d4")), "e4 d4")
	assert.True(t, validateQueenMoveWithState(s, move.NewMove("e4", "c4")), "e4 c4")
}

func TestValidateQueenMoveWithStateInvalid(t *testing.T) {
	s := *NewTestStateFromFEN("p7/1P2P3/8/5p2/1Pp1Q2p/8/2P3P1/1P6 w - - 1 1")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "b7")), "e4 b7")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "a8")), "e4 a8")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "e7")), "e4 e7")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "e8")), "e4 e8")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "g6")), "e4 g6")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "h7")), "e4 h7")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "g2")), "e4 g2")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "h1")), "e4 h1")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "c2")), "e4 c2")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "b1")), "e4 b1")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "b4")), "e4 b4")
	assert.False(t, validateQueenMoveWithState(s, move.NewMove("e4", "a4")), "e4 a4")
}

func TestValidatePawnMoveWithStateEnPassant(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/8/3Pp3/8/8/8/8 w - e6 0 1")
	assert.Equal(t, "e6", s.EnPassantTarget)
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("d5", "e6")), "d5 e6")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("d5", "d6")), "d5 d6")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d5", "c6")), "d5 c6")
	s.EnPassantTarget = "-"
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d5", "e6")), "d5 e6")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("d5", "d6")), "d5 d6")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d5", "c6")), "d5 c6")

	s = *NewTestStateFromFEN("8/8/4p3/3P4/8/8/8/8 w - e6 0 1")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d4", "e6")), "d4 e6")

	s = *NewTestStateFromFEN("8/8/8/8/3Pp3/8/8/8 b - d3 0 1")
	assert.Equal(t, "d3", s.EnPassantTarget)
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e4", "d3")), "e4 d3")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e4", "e3")), "e4 e3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("e4", "f3")), "e4 f3")
	s.EnPassantTarget = "-"
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("e4", "d3")), "e4 d3")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("e4", "e3")), "e4 e3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("e4", "f3")), "e4 f3")
}

func TestValidatePawnMoveWithStateJumpOverPiece(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/8/8/2Pp4/Pp6/PPPP4/8 w - - 0 1")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("a2", "a3")), "a2 a3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("a2", "a4")), "a2 a4")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("b2", "b3")), "b2 b3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("b2", "b4")), "b2 b4")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("c2", "c3")), "c2 c3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("c2", "c4")), "c2 c4")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("d2", "d3")), "d2 d3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d2", "d4")), "d2 d4")
}

func TestValidatePawnMoveWithStateMaxRanks(t *testing.T) {
	s := *NewTestStateFromFEN("8/8/8/8/2Pp4/Pp6/PPPP4/8 w - - 0 1")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("a2", "a3")), "a2 a3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("a2", "a4")), "a2 a4")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("b2", "b3")), "b2 b3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("b2", "b4")), "b2 b4")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("c2", "c3")), "c2 c3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("c2", "c4")), "c2 c4")
	assert.True(t, validatePawnMoveWithState(s, move.NewMove("d2", "d3")), "d2 d3")
	assert.False(t, validatePawnMoveWithState(s, move.NewMove("d2", "d4")), "d2 d4")
}

func TestValidateKingMoveWithStateCastlingValid(t *testing.T) {
	s := *NewTestStateFromFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e1", "g1")), "e1 g1")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e1", "c1")), "e1 c1")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e8", "g8")), "e8 g8")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e8", "c8")), "e8 c8")

	assert.True(t, validateKingMoveWithState(s, move.NewMove("e1", "f1")), "e1 f1")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e1", "d1")), "e1 d1")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e8", "f8")), "e8 f8")
	assert.True(t, validateKingMoveWithState(s, move.NewMove("e8", "d8")), "e8 c8")
}

func TestValidateKingMoveWithStateCastlingInvalidRights(t *testing.T) {
	s := *NewTestStateFromFEN("r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "g1")), "e1 g1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "c1")), "e1 c1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "g8")), "e8 g8")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "c8")), "e8 c8")
}

func TestValidateKingMoveWithStateCastlingInvalidBlockingPieces(t *testing.T) {
	s := *NewTestStateFromFEN("r2qkb1r/8/8/8/8/8/8/R2QKB1R w KQkq - 0 1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "g1")), "e1 g1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "c1")), "e1 c1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "g8")), "e8 g8")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "c8")), "e8 c8")

	s = *NewTestStateFromFEN("r1b1k1nr/8/8/8/8/8/8/R1B1K1NR w KQkq - 0 1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "g1")), "e1 g1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "c1")), "e1 c1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "g8")), "e8 g8")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "c8")), "e8 c8")

	s = *NewTestStateFromFEN("rn2k2r/8/8/8/8/8/8/RN2K2R w KQkq - 0 1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e1", "c1")), "e1 c1")
	assert.False(t, validateKingMoveWithState(s, move.NewMove("e8", "c8")), "e8 c8")
}

func TestGenerateMovesDoesntChangeState(t *testing.T) {
	fen := "rnbqkbnr/p1ppp1pp/1p6/5p2/2P5/P7/1P1PPPPP/RNBQKBNR w Kkq f6 0 3"
	s := *NewTestStateFromFEN(fen)
	s.GeneratePossibleMoves()
	assert.Equal(t, fen, s.FEN())
}

func TestGeneratePossibleMoves(t *testing.T) {
	tests := map[string]struct {
		fen           string
		possibleMoves []move.Move
	}{
		"block check": {
			fen:           "3pkp2/3p1p2/2n5/4Q3/8/8/8/8 b - - 0 1",
			possibleMoves: []move.Move{move.NewMove("c6", "e5"), move.NewMove("c6", "e7")},
		},
		"move into check": {
			fen:           "7k/Q7/8/8/8/8/8/8 b - - 0 1",
			possibleMoves: []move.Move{move.NewMove("h8", "g8")},
		},
		"checkmate": {
			fen:           "R6k/Q7/8/8/8/8/8/8 b - - 0 1",
			possibleMoves: []move.Move{},
		},
		"pin": {
			fen:           "Q5pk/8/8/8/8/8/8/6R1 b - - 0 1",
			possibleMoves: []move.Move{move.NewMove("h8", "h7")},
		},
		"stalemate": {
			fen:           "Q5pk/R7/8/8/8/8/8/8 b - - 0 1",
			possibleMoves: []move.Move{},
		},
		"a": {
			fen:           "2R1qk2/1Q6/8/8/8/8/8/8 b - - 1 37",
			possibleMoves: []move.Move{move.NewMove("e8", "c8"), move.NewMove("e8", "d8"), move.NewMove("f8", "g8")},
		},
		"only move take queen with king": {
			fen:           "rnbqkbnr/pp1p1Qp1/7p/2p1p3/4P3/8/PPPP1PPP/RNB1KBNR b KQkq - 0 4",
			possibleMoves: []move.Move{move.NewMove("e8", "f7")},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := *NewTestStateFromFEN(test.fen)
			moves := s.GeneratePossibleMoves()
			assert.Len(t, moves, len(test.possibleMoves))
			assert.Equal(t, test.possibleMoves, moves)
		})
	}
}

func TestGeneratePossibleMovesNotIn(t *testing.T) {
	tests := map[string]struct {
		fen              string
		notPossibleMoves []move.Move
	}{
		"a": {
			fen:              "r1R2k1r/nQ6/P4ppp/P2Pqp1n/3BP3/1N5P/3bK1P1/5BR1 b - - 0 36",
			notPossibleMoves: []move.Move{move.NewMove("e8", "e5")},
		},
		"b": {
			fen:              "r1bqk2r/p1p2p1p/n3p1pn/1p1p2bQ/1P2P3/N2P3P/PRP2PP1/2B1KBNR w Kkq - 0 1",
			notPossibleMoves: []move.Move{move.NewMove("a7", "a5")},
		},
		"c": {
			fen:              "rR1qkb1r/3ppp2/7n/1b1n1Ppp/5B2/2PP2PN/4PK1P/1N1Q1B1R w q - 1 15",
			notPossibleMoves: []move.Move{move.NewMove("d8", "c7")},
		},
		"d": {
			fen:              "r1bqk1n1/p1pp1pr1/n6p/1p4p1/1b5P/QNPp2P1/PP2PP2/R1B1KBNR w KQq - 4 11",
			notPossibleMoves: []move.Move{move.NewMove("c3", "c4")},
		},
		"e": {
			fen:              board.StartingFEN,
			notPossibleMoves: []move.Move{move.NewMove("g2", "h3")},
		},
		"white can't capture own en passant": {
			fen:              "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR w KQkq d3 0 1",
			notPossibleMoves: []move.Move{move.NewMove("c2", "d3")},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := *NewTestStateFromFEN(test.fen)
			assert.Equal(t, test.fen, s.FEN())
			moves := s.GeneratePossibleMoves()
			for _, m := range test.notPossibleMoves {
				assert.NotContains(t, moves, m)
			}
		})
	}
}
