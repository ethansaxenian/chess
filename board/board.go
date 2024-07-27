package board

import (
	"fmt"
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

var precomputedSquareToCoords = map[string][2]int{}

func init() {
	for i, f := range Files {
		for j, r := range Ranks {
			precomputedSquareToCoords[string(f)+string(r)] = [2]int{i + 1, j + 1}
		}
	}
}

func SquareToCoords(square string) (int, int) {
	coords, ok := precomputedSquareToCoords[square]
	assert.Assert(ok, fmt.Sprintf("%s is not a valid square", square))

	return coords[0], coords[1]
}

func squareToIndex(square string) int {
	assert.Assert(len(square) == 2, fmt.Sprintf("invalid square: %s", square))

	rank := strings.IndexByte(Ranks, square[1])
	assert.Assert(rank != -1, fmt.Sprintf("invalid rank: %d", rank))

	file := strings.IndexByte(Files, square[0])
	assert.Assert(file != -1, fmt.Sprintf("invalid file: %d", file))

	return rank*boardLength + file
}

func indexToSquare(index int) string {
	return string(Files[index%8]) + string(Ranks[index/8])
}

type chessboard [64]piece.Piece

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

			var color piece.Piece
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

func (b chessboard) Print() {
	whiteSquare := "\033[48;2;194;167;120m"
	whitePiece := "\033[38;2;255;255;255m"

	blackSquare := "\033[48;2;131;99;69m"
	blackPiece := "\033[38;2;0;0;0m"

	reset := "\033[0m"

	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
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

			char := string(piece.PieceToChar[piece.Value(p)])
			fmt.Print(squareColor, pieceColor, " ", char, " ", reset)
		}
		fmt.Println()
	}

	fmt.Println("   a  b  c  d  e  f  g  h")
}

func (b *chessboard) MakeMove(src, dest string) {
	srcIndex := squareToIndex(src)
	destIndex := squareToIndex(dest)

	b[destIndex] = b[srcIndex]
	b[srcIndex] = piece.None
}

func (b chessboard) Square(square string) piece.Piece {
	return b[squareToIndex(square)]
}

func (b chessboard) Squares() map[string]piece.Piece {
	squares := map[string]piece.Piece{}
	for i, p := range b {
		squares[indexToSquare(i)] = p
	}
	return squares
}
