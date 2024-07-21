package board

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/piece"
)

const boardLength = 8

const (
	Files = "abcdefgh"
	Ranks = "12345678"
)

const StartingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func SquareToIndex(square string) int {
	assert.Assert(len(square) == 2, fmt.Sprintf("invalid square: %s", square))

	rank := strings.IndexByte(Ranks, square[1])
	assert.Assert(rank != -1, fmt.Sprintf("invalid rank: %d", rank))

	file := strings.IndexByte(Files, square[0])
	assert.Assert(file != -1, fmt.Sprintf("invalid file: %d", file))

	return rank*boardLength + file
}

type chessboard [64]int

func LoadFEN(fen string) chessboard {
	piecePlacement := strings.Split(fen, " ")[0]

	file := 0
	rank := 7

	var board chessboard

	for _, char := range piecePlacement {
		if char == '/' {
			rank--
			file = 0
		} else if unicode.IsDigit(char) {
			file += int(char - '0')
		} else {
			p := piece.CharToPiece[unicode.ToLower(char)]

			var color int
			if unicode.IsUpper(char) {
				color = piece.White
			} else {
				color = piece.Black
			}

			board[rank*boardLength+file] = p * color
			file++
		}
	}

	return board
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (b chessboard) Print() {
	clearScreen()
	whiteSquare := "\033[48;2;194;167;120m"
	whitePiece := "\033[38;2;255;255;255m"

	blackSquare := "\033[48;2;131;99;69m"
	blackPiece := "\033[38;2;0;0;0m"

	reset := "\033[0m"

	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			p := b[rank*boardLength+file]

			var pieceColor string
			if p < 0 {
				pieceColor = blackPiece
			} else if p > 0 {
				pieceColor = whitePiece
			}

			var squareColor string
			if (rank%2 == 0 && file%2 != 0) || (rank%2 != 0 && file%2 == 0) {
				squareColor = whiteSquare
			} else {
				squareColor = blackSquare
			}

			char := string(piece.PieceToChar[int(math.Abs(float64(p)))])
			fmt.Print(squareColor, pieceColor, " ", char, " ", reset)
		}
		fmt.Println()
	}
}

func (b *chessboard) MakeMove(src, dest string) {
	srcIndex := SquareToIndex(src)
	destIndex := SquareToIndex(dest)

	b[destIndex] = b[srcIndex]
	b[srcIndex] = piece.None
}
