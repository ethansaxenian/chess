package piece

import (
	"strings"

	"github.com/ethansaxenian/chess/assert"
)

type Piece int

const (
	Empty Piece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	White Piece = 1
	Black Piece = -1
)

type Side int

const (
	Kingside Side = iota
	Queenside
)

func (p Piece) Color() Piece {
	if p > 0 {
		return White
	} else if p < 0 {
		return Black
	} else {
		return Empty
	}
}

func (p Piece) Type() Piece {
	return p * p.Color()
}

func (p Piece) FEN() string {
	var char string
	switch t := p.Type(); t {
	case Empty:
		assert.Raise("piece.None has no FEN repr")
	case Pawn:
		char = "p"
	case Knight:
		char = "n"
	case Bishop:
		char = "b"
	case Rook:
		char = "r"
	case Queen:
		char = "q"
	case King:
		char = "k"
	default:
		char = ""
	}

	if p.Color() == White {
		char = strings.ToUpper(char)
	}

	return char
}

func (p Piece) String() string {
	switch p.Type() {
	case Empty:
		return " "
	case Pawn:
		return "♟"
	case Knight:
		return "♞"
	case Bishop:
		return "♝"
	case Rook:
		return "♜"
	case Queen:
		return "♛"
	case King:
		return "♚"
	default:
		return ""
	}
}

var SlidingPieces = []Piece{Bishop, Rook, Queen}

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
	'_': Empty,
	'p': Pawn,
	'n': Knight,
	'b': Bishop,
	'r': Rook,
	'q': Queen,
	'k': King,
}

var StartingKingSquares = map[Piece]string{
	White: "e1",
	Black: "e8",
}

var StartingRookSquares = map[Piece]map[Side]string{
	White: {Kingside: "h1", Queenside: "a1"},
	Black: {Kingside: "h8", Queenside: "a8"},
}

var CastlingSquares = map[Piece]map[Side]string{
	White: {Kingside: "g1", Queenside: "c1"},
	Black: {Kingside: "g8", Queenside: "c8"},
}

var CastlingIntermediateSquares = map[Piece]map[Side][]string{
	White: {Kingside: {"f1", "g1"}, Queenside: {"d1", "c1", "b1"}},
	Black: {Kingside: {"f8", "g8"}, Queenside: {"d8", "c8", "b8"}},
}

var RookCastlingSquares = map[Piece]map[Side]string{
	White: {Kingside: "f1", Queenside: "d1"},
	Black: {Kingside: "f8", Queenside: "d8"},
}
