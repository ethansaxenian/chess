package piece

import (
	"math"
	"strings"

	"github.com/ethansaxenian/chess/assert"
)

type Piece int

const (
	None Piece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

var SlidingPieces = []Piece{Bishop, Rook, Queen}

const (
	White Piece = 1
	Black Piece = -1
)

var StartingPawnRanks = map[Piece]int{
	Pawn * White: 2,
	Pawn * Black: 7,
}

var MaxPawnRank = map[Piece]int{
	Pawn * White: 8,
	Pawn * Black: 1,
}

var AllPieces = []Piece{Pawn, Knight, Bishop, Rook, Queen, King}
var AllColors = []Piece{White, Black}
var PossiblePromotions = []Piece{Knight, Bishop, Rook, Queen}

var CharToPiece = map[rune]Piece{
	'_': None,
	'p': Pawn,
	'n': Knight,
	'b': Bishop,
	'r': Rook,
	'q': Queen,
	'k': King,
}

var PieceToChar = map[Piece]rune{
	None:   ' ',
	Pawn:   '♟',
	Knight: '♞',
	Bishop: '♝',
	Rook:   '♜',
	Queen:  '♛',
	King:   '♚',
}

var PieceToRepr = map[Piece]string{
	None:   "",
	Pawn:   "",
	Knight: "n",
	Bishop: "b",
	Rook:   "r",
	Queen:  "q",
	King:   "k",
}

func FENRepr(piece Piece) string {
	var char string
	switch v := Value(piece); v {
	case Pawn:
		char = "p"
	case None:
		assert.Raise("piece.None has no FEN repr")
	default:
		char = PieceToRepr[v]
	}

	if IsColor(piece, White) {
		char = strings.ToUpper(char)
	}

	return char
}

func Value(p Piece) Piece {
	return Piece(math.Abs(float64(p)))
}

func Color(p Piece) Piece {
	if p > 0 {
		return White
	} else if p < 0 {
		return Black
	} else {
		return None
	}
}

func IsColor(p, c Piece) bool {
	return p*c > 0
}
