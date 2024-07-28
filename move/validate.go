package move

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
)

func validateMove(state board.State, src, target string) bool {
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

	if !validatePieceMoveWithBoard(state.Board, srcPiece, src, target) {
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
	// TODO: en passant, promotions
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
	// TODO: castling
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

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

func validatePieceMoveWithBoard(b board.Chessboard, srcPiece piece.Piece, src, target string) bool {
	switch piece.Value(srcPiece) {
	case piece.Pawn:
		return validatePawnMoveWithBoard(b, src, target)
	case piece.Knight:
		return true
	case piece.Bishop:
		return validateBishopMoveWithBoard(b, src, target)
	case piece.Rook:
		return validateRookMoveWithBoard(b, src, target)
	case piece.Queen:
		return validateQueenMoveWithBoard(b, src, target)
	case piece.King:
		return true
	default:
		return false
	}
}

func validatePawnMoveWithBoard(b board.Chessboard, src, target string) bool {
	// TODO en passant
	srcPiece := b.Square(src)
	targetPiece := b.Square(target)

	// diagonal pawn moves
	if src[0] != target[0] && !piece.IsColor(targetPiece, piece.Color(srcPiece)*-1) {
		return false
	}

	// can't capture forward
	if src[0] == target[0] && targetPiece != piece.None {
		return false
	}

	return true
}

func validateBishopMoveWithBoard(b board.Chessboard, src, target string) bool {
	sf, sr := board.SquareToCoords(src)
	tf, tr := board.SquareToCoords(target)

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

	srcPiece := b.Square(src)

	for f, r := sf+df, sr+dr; f >= 'a' && f <= 'h' && r >= 1 && r <= 8; f, r = f+df, r+dr {
		currPiece := b.Square(board.CoordsToSquare(f, r))

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

	return false

}

func validateRookMoveWithBoard(b board.Chessboard, src, target string) bool {
	sf, sr := board.SquareToCoords(src)
	tf, tr := board.SquareToCoords(target)

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

	srcPiece := b.Square(src)

	for f, r := sf+df, sr+dr; ; f, r = f+df, r+dr {
		currPiece := b.Square(board.CoordsToSquare(f, r))

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

func validateQueenMoveWithBoard(b board.Chessboard, src, target string) bool {
	return validateBishopMoveWithBoard(b, src, target) || validateRookMoveWithBoard(b, src, target)
}
