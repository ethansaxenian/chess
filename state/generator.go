package state

import (
	"fmt"
	"math"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
)

var precomputedPieceMoves = map[piece.Piece]map[string][]string{}

func init() {
	for _, p := range piece.AllPieces {
		for _, c := range piece.AllColors {
			pieceMap := map[string][]string{}
			for _, sf := range board.Files {
				for _, sr := range board.Ranks {
					src := string(sf) + string(sr)
					assert.Assert(src != "e0", "foobar")
					for _, tf := range board.Files {
						for _, tr := range board.Ranks {
							target := string(tf) + string(tr)
							if validatePieceMove(p*c, src, target) {
								pieceMap[src] = append(pieceMap[src], target)
							}
						}
					}
				}
			}
			precomputedPieceMoves[p*c] = pieceMap
		}
	}
}

func generateTmpMoves(state State) []move.Move {
	moves := []move.Move{}

	for source, p := range state.board.Squares() {
		if p == piece.Empty {
			continue
		}

		for _, target := range precomputedPieceMoves[p][source] {
			if validateMove(state, source, target) {
				moves = append(moves, move.NewMove(source, target))
			}
		}
	}

	move.SortMoves(moves)

	return moves
}

func validateMove(state State, src, target string) bool {
	srcPiece := state.Piece(src)
	targetPiece := state.Piece(target)

	// src is my color
	if srcPiece.Color() != state.ActiveColor {
		return false
	}

	// target is not my color
	if targetPiece.Color() == state.ActiveColor {
		return false
	}

	if !validatePieceMoveWithState(state, srcPiece, src, target) {
		return false
	}

	return true
}

func validatePieceMove(srcPiece piece.Piece, srcSquare, targetSquare string) bool {
	switch srcPiece.Type() {
	case piece.Pawn:
		return validatePawnMove(srcSquare, targetSquare, srcPiece.Color())
	case piece.Knight:
		return validateKnightMove(srcSquare, targetSquare)
	case piece.Bishop:
		return validateBishopMove(srcSquare, targetSquare)
	case piece.Rook:
		return validateRookMove(srcSquare, targetSquare)
	case piece.Queen:
		return validateQueenMove(srcSquare, targetSquare)
	case piece.King:
		return validateKingMove(srcSquare, targetSquare)
	default:
		return false
	}
}

func validatePawnMove(src, target string, color piece.Piece) bool {
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	fileDiff := math.Abs(float64(srcFile - targetFile))
	rankDiff := (targetRank - srcRank) * int(color)

	if fileDiff > 1 {
		return false
	} else if fileDiff == 1 && rankDiff != 1 {
		return false
	}

	startingRank, ok := piece.StartingPawnRanks[piece.Pawn*color]
	assert.Assert(ok, fmt.Sprintf("invalid piece color for pawn: %d", color))

	if srcRank == startingRank && (rankDiff != 1 && rankDiff != 2) {
		return false
	} else if srcRank != startingRank && rankDiff != 1 {
		return false
	}

	return true
}

func validateKnightMove(src, target string) bool {
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	rankDiff := math.Abs(float64(targetRank - srcRank))
	fileDiff := math.Abs(float64(targetFile - srcFile))

	if !(rankDiff == 2 && fileDiff == 1) && !(rankDiff == 1 && fileDiff == 2) {
		return false
	}

	return true
}

func validateBishopMove(src, target string) bool {
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	rankDiff := math.Abs(float64(targetRank - srcRank))
	fileDiff := math.Abs(float64(targetFile - srcFile))

	if rankDiff == 0 && fileDiff == 0 {
		return false
	}

	if rankDiff != fileDiff {
		return false
	}

	return true
}

func validateRookMove(src, target string) bool {
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	return (srcRank == targetRank) != (srcFile == targetFile)
}

func validateQueenMove(src, target string) bool {
	return validateBishopMove(src, target) || validateRookMove(src, target)
}

func validateKingMove(src, target string) bool {
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	rankDiff := math.Abs(float64(targetRank - srcRank))
	fileDiff := math.Abs(float64(targetFile - srcFile))

	// no movement
	if rankDiff == 0 && fileDiff == 0 {
		return false
	}

	if rankDiff > 1 {
		return false
	}

	if fileDiff > 2 {
		return false
	}

	if src != "e1" && src != "e8" && fileDiff > 1 {
		return false
	}

	// castling possibilities
	if src == "e1" && target != "c1" && target != "g1" && fileDiff > 1 {
		return false
	}
	if src == "e8" && target != "c8" && target != "g8" && fileDiff > 1 {
		return false
	}

	return true
}

func validatePieceMoveWithState(s State, srcPiece piece.Piece, src, target string) bool {
	switch srcPiece.Type() {
	case piece.Pawn:
		return validatePawnMoveWithState(s, src, target)
	case piece.Knight:
		return true
	case piece.Bishop:
		return validateBishopMoveWithState(s, src, target)
	case piece.Rook:
		return validateRookMoveWithState(s, src, target)
	case piece.Queen:
		return validateQueenMoveWithState(s, src, target)
	case piece.King:
		return validateKingMoveWithState(s, src, target)
	default:
		return false
	}
}

func validatePawnMoveWithState(s State, src, target string) bool {
	srcPiece := s.Piece(src)
	srcColor := srcPiece.Color()
	targetPiece := s.Piece(target)

	isCaptureAttempt := src[0] != target[0]
	isEnPassantAttempt := s.EnPassantTarget == target
	isOppositeColorPiece := targetPiece.Color() == srcColor*-1
	isDoubleMove := int(target[1])-int(src[1]) == 2*int(srcColor)
	jumpsOverPiece := s.Piece(board.AddRank(src, int(srcColor))) != piece.Empty

	if isCaptureAttempt {
		if isEnPassantAttempt && targetPiece != piece.Empty {
			return false
		}

		if !isEnPassantAttempt && !isOppositeColorPiece {
			return false
		}
	}

	if !isCaptureAttempt && targetPiece != piece.Empty {
		return false
	}

	if isDoubleMove && jumpsOverPiece {
		return false
	}

	return true
}

func validateBishopMoveWithState(s State, src, target string) bool {
	sf, sr := board.SquareToCoords(src)
	tf, tr := board.SquareToCoords(target)

	if sf == tf || sr == tr {
		return false
	}

	var df, dr int

	if tf > sf {
		df = 1
	} else {
		df = -1
	}

	if tr > sr {
		dr = 1
	} else {
		dr = -1
	}

	srcPiece := s.Piece(src)

	for f, r := sf+df, sr+dr; ; f, r = f+df, r+dr {
		assert.Assert(f >= 'a' && f <= 'h' && r >= 1 && r <= 8, fmt.Sprintf("%s%s: %d/%d", src, target, df, dr))
		currPiece := s.Piece(board.CoordsToSquare(f, r))

		if f == tf && r == tr {
			if currPiece == piece.Empty {
				return true
			} else if currPiece.Color() == srcPiece.Color() {
				return false
			} else {
				return true
			}
		} else if currPiece != piece.Empty {
			return false
		}

	}

}

func validateRookMoveWithState(s State, src, target string) bool {
	sf, sr := board.SquareToCoords(src)
	tf, tr := board.SquareToCoords(target)

	if sf != tf && sr != tr {
		return false
	}

	var df, dr int

	if sf > tf {
		df = -1
	} else if sf < tf {
		df = 1
	}

	if sr > tr {
		dr = -1
	} else if sr < tr {
		dr = 1
	}

	srcPiece := s.Piece(src)

	for f, r := sf+df, sr+dr; ; f, r = f+df, r+dr {
		currPiece := s.Piece(board.CoordsToSquare(f, r))

		if f == tf && r == tr {
			if currPiece == piece.Empty {
				return true
			} else if currPiece.Color() == srcPiece.Color() {
				return false
			} else {
				return true
			}
		} else if currPiece != piece.Empty {
			return false
		}
	}
}

func validateQueenMoveWithState(s State, src, target string) bool {
	return validateBishopMoveWithState(s, src, target) || validateRookMoveWithState(s, src, target)
}

func validateKingMoveWithState(s State, src, target string) bool {
	color := s.Piece(src).Color()

	startingSquare := piece.StartingKingSquares[color]
	castlingSquares := piece.CastlingSquares[color]
	intermediateSquares := piece.CastlingIntermediateSquares[color]
	castlingRights := s.Castling[color]

	if src == startingSquare {
		for i := range []int{0, 1} {
			// not trying to castle
			if target != castlingSquares[i] {
				continue
			}

			// can't castle
			if !castlingRights[i] {
				return false
			}

			// blocking pieces
			for _, square := range intermediateSquares[i] {
				if s.Piece(square) != piece.Empty {
					return false
				}
			}
		}
	}

	return true
}
