package move

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/game"
	"github.com/ethansaxenian/chess/piece"
)

func validateMove(state game.State, src, target string) bool {
	srcPiece := state.Board.Square(src)
	targetPiece := state.Board.Square(target)

	// src is my color
	if !piece.IsColor(srcPiece, state.ActiveColor) {
		return false
	}

	// target is not my color
	if piece.IsColor(targetPiece, state.ActiveColor) {
		return false
	}

	if !validatePieceMoveWithBoard(state, srcPiece, src, target) {
		return false
	}

	return true
}

func validatePieceMove(srcPiece piece.Piece, srcSquare, targetSquare string) bool {
	slog.Debug("piece info:", "piece", srcPiece, "src", srcSquare, "target", targetSquare)
	switch piece.Value(srcPiece) {
	case piece.Pawn:
		return validatePawnMove(srcSquare, targetSquare, piece.Color(srcPiece))
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
	srcFile, srcRank := game.SquareToCoords(src)
	targetFile, targetRank := game.SquareToCoords(target)

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
	srcFile, srcRank := game.SquareToCoords(src)
	targetFile, targetRank := game.SquareToCoords(target)

	rankDiff := math.Abs(float64(targetRank - srcRank))
	fileDiff := math.Abs(float64(targetFile - srcFile))

	if !(rankDiff == 2 && fileDiff == 1) && !(rankDiff == 1 && fileDiff == 2) {
		return false
	}

	return true
}

func validateBishopMove(src, target string) bool {
	srcFile, srcRank := game.SquareToCoords(src)
	targetFile, targetRank := game.SquareToCoords(target)

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
	srcFile, srcRank := game.SquareToCoords(src)
	targetFile, targetRank := game.SquareToCoords(target)

	return (srcRank == targetRank) != (srcFile == targetFile)
}

func validateQueenMove(src, target string) bool {
	return validateBishopMove(src, target) || validateRookMove(src, target)
}

func validateKingMove(src, target string) bool {
	// TODO: castling
	srcFile, srcRank := game.SquareToCoords(src)
	targetFile, targetRank := game.SquareToCoords(target)

	rankDiff := math.Abs(float64(targetRank - srcRank))
	fileDiff := math.Abs(float64(targetFile - srcFile))

	// no movement
	if rankDiff == 0 && fileDiff == 0 {
		return false
	}

	if rankDiff > 1 || fileDiff > 1 {
		return false
	}

	return true
}

func validatePieceMoveWithBoard(s game.State, srcPiece piece.Piece, src, target string) bool {
	switch piece.Value(srcPiece) {
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
		return true
	default:
		return false
	}
}

func validatePawnMoveWithState(s game.State, src, target string) bool {
	srcPiece := s.Board.Square(src)
	srcColor := piece.Color(srcPiece)
	targetPiece := s.Board.Square(target)

	isCaptureAttempt := src[0] != target[0]
	isEnPassantAttempt := s.EnPassantTarget == target
	isOppositeColorPiece := piece.IsColor(targetPiece, srcColor*-1)
	isDoubleMove := target[1]-src[1] == 2
	jumpsOverPiece := s.Board.Square(game.AddRank(src, int(srcColor))) != piece.None

	if isCaptureAttempt {
		if isEnPassantAttempt && targetPiece != piece.None {
			return false
		}

		if !isEnPassantAttempt && !isOppositeColorPiece {
			return false
		}
	}

	if !isCaptureAttempt && targetPiece != piece.None {
		return false
	}

	if isDoubleMove && jumpsOverPiece {
		return false
	}

	return true
}

func validateBishopMoveWithState(s game.State, src, target string) bool {
	sf, sr := game.SquareToCoords(src)
	tf, tr := game.SquareToCoords(target)

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

	srcPiece := s.Board.Square(src)

	for f, r := sf+df, sr+dr; ; f, r = f+df, r+dr {
		assert.Assert(f >= 'a' && f <= 'h' && r >= 1 && r <= 8, fmt.Sprintf("%s%s: %d/%d", src, target, df, dr))
		currPiece := s.Board.Square(game.CoordsToSquare(f, r))

		if f == tf && r == tr {
			if currPiece == piece.None {
				return true
			} else if piece.IsColor(currPiece, srcPiece) {
				return false
			} else {
				return true
			}
		} else if currPiece != piece.None {
			return false
		}

	}

}

func validateRookMoveWithState(s game.State, src, target string) bool {
	sf, sr := game.SquareToCoords(src)
	tf, tr := game.SquareToCoords(target)

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

	srcPiece := s.Board.Square(src)

	for f, r := sf+df, sr+dr; ; f, r = f+df, r+dr {
		currPiece := s.Board.Square(game.CoordsToSquare(f, r))

		if f == tf && r == tr {
			if currPiece == piece.None {
				return true
			} else if piece.IsColor(currPiece, srcPiece) {
				return false
			} else {
				return true
			}
		} else if currPiece != piece.None {
			return false
		}
	}
}

func validateQueenMoveWithState(s game.State, src, target string) bool {
	return validateBishopMoveWithState(s, src, target) || validateRookMoveWithState(s, src, target)
}
