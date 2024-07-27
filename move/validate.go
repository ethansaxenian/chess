package move

import (
	"fmt"
	"log/slog"
	"math"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
)

func validateMove(move string, state board.State) bool {
	if len(move) != 4 {
		return false
	}

	srcSquare := move[:2]
	targetSquare := move[2:]

	if !validateBounds(srcSquare) {
		return false
	}

	if !validateBounds(targetSquare) {
		return false
	}

	srcPiece := state.Board.Square(srcSquare)
	targetPiece := state.Board.Square(targetSquare)

	if !validateSrc(srcPiece, state) {
		return false
	}

	if !validateTarget(targetPiece, state) {
		return false
	}

	if !validatePieceMovement(srcPiece, srcSquare, targetSquare) {
		return false
	}

	return true
}

func validateSrc(src piece.Piece, state board.State) bool {
	// no piece on square
	if src == piece.None {
		return false
	}

	if !piece.IsColor(src, state.CurrColor) {
		return false
	}

	return true
}

func validateTarget(target piece.Piece, state board.State) bool {
	if piece.IsColor(target, state.CurrColor) {
		return false
	}

	return true
}

func validateBounds(square string) bool {
	if len(square) != 2 {
		return false
	}

	rank := strings.IndexByte(board.Ranks, square[1])
	if rank == -1 {
		return false
	}

	file := strings.IndexByte(board.Files, square[0])
	if file == -1 {
		return false
	}

	return true
}

func validatePieceMovement(srcPiece piece.Piece, srcSquare, targetSquare string) bool {
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
	// TODO: captures, en passant, promotions
	srcFile, srcRank := board.SquareToCoords(src)
	targetFile, targetRank := board.SquareToCoords(target)

	if srcFile != targetFile {
		return false
	}

	rankDiff := (targetRank - srcRank) * int(color)

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
