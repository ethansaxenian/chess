package board

import (
	"fmt"
	"strconv"
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

func SquareToCoords(square string) (int, int) {
	f := int(square[0])
	r, err := strconv.Atoi(string(square[1]))
	assert.NilError(err, fmt.Sprintf("invalid square: %s", square))
	assert.Assert(square[0] >= 'a' && square[0] <= 'h', fmt.Sprintf("SquareToCoords: %s", square))
	return f, r
}

func CoordsToSquare(f, r int) string {
	assert.Assert(f >= 'a' && f <= 'h' && r >= 1 && r <= 8, fmt.Sprintf("CoordsToSquare: %d %d", f, r))
	return string(rune(f)) + strconv.Itoa(r)
}

func squareToIndex(square string) int {
	file := int(square[0]) - 97
	rank, err := strconv.Atoi(string(square[1]))
	assert.NilError(err, fmt.Sprintf("invalid square: %s", square))
	index := (rank-1)*boardLength + file
	assert.Assert(index >= 0 && index < 64, fmt.Sprintf("squareToIndex: %s -> %d", square, index))
	return index
}

func indexToSquare(index int) string {
	f := rune(index%8 + 97)
	r := index/8 + 1
	assert.Assert(f >= 'a' && f <= 'h' && r >= 1 && r <= 8, fmt.Sprintf("indexToSquare: %d", index))

	return string(f) + strconv.Itoa(r)
}

type Chessboard [64]piece.Piece

func LoadFEN(fen string) Chessboard {
	piecePlacement := strings.Split(fen, " ")[0]

	file := 0
	rank := 7

	var board Chessboard

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

func (b Chessboard) Print() {
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

func (b *Chessboard) MakeMove(src, dest string) {
	srcIndex := squareToIndex(src)
	destIndex := squareToIndex(dest)

	b[destIndex] = b[srcIndex]
	b[srcIndex] = piece.None
}

func (b Chessboard) Square(square string) piece.Piece {
	return b[squareToIndex(square)]
}

func (b Chessboard) Squares() map[string]piece.Piece {
	squares := map[string]piece.Piece{}
	for i, p := range b {
		squares[indexToSquare(i)] = p
	}
	return squares
}
