package piece

const (
	None int = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	White int = 1
	Black int = -1
)

var CharToPiece = map[rune]int{
	'_': None,
	'p': Pawn,
	'n': Knight,
	'b': Bishop,
	'r': Rook,
	'q': Queen,
	'k': King,
}

var PieceToChar = map[int]rune{
	None:   ' ',
	Pawn:   '♟',
	Knight: '♞',
	Bishop: '♝',
	Rook:   '♜',
	Queen:  '♛',
	King:   '♚',
}
