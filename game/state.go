package game

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

const noEnPassantTarget = "-"

type State struct {
	Players         map[piece.Piece]player.Player
	EnPassantTarget string
	Moves           []move
	Board           Chessboard
	nextBoard       Chessboard
	ActiveColor     piece.Piece
	whiteCastling   [2]bool
	blackCastling   [2]bool
	halfmoveClock   int
	fullmoveNumber  int
}

func StartingState(white, black player.Player) *State {
	return LoadFEN(StartingFEN, white, black)
}

func LoadFEN(fen string, white, black player.Player) *State {
	fenFields := strings.Fields(fen)

	b := loadFEN(fenFields[0])

	var activeColor piece.Piece
	switch fenFields[1] {
	case "w":
		activeColor = piece.White
	case "b":
		activeColor = piece.Black
	default:
		assert.Raise(fmt.Sprintf("invalid FEN active color: %s", fen))
	}

	whiteCastling := [2]bool{}
	blackCastling := [2]bool{}
	for _, char := range fenFields[2] {
		switch char {
		case 'K':
			whiteCastling[0] = true
		case 'Q':
			whiteCastling[1] = true
		case 'k':
			blackCastling[0] = true
		case 'q':
			blackCastling[1] = true
		case '-':
			continue
		default:
			assert.Raise(fmt.Sprintf("invalid FEN castling rights: %s", fen))
		}
	}

	enPassantTarget := fenFields[3]

	halfmoveClock, err := strconv.Atoi(fenFields[4])
	assert.ErrIsNil(err, fmt.Sprintf("invalid FEN halfmove clock: %s", fen))
	fullmoveNumber, err := strconv.Atoi(fenFields[5])
	assert.ErrIsNil(err, fmt.Sprintf("invalid FEN fullmove clock: %s", fen))

	s := &State{
		Board:           b,
		nextBoard:       b,
		Players:         map[piece.Piece]player.Player{piece.White: white, piece.Black: black},
		ActiveColor:     activeColor,
		Moves:           []move{},
		EnPassantTarget: enPassantTarget,
		whiteCastling:   whiteCastling,
		blackCastling:   blackCastling,
		halfmoveClock:   halfmoveClock,
		fullmoveNumber:  fullmoveNumber,
	}

	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)

	return s
}

func (s State) FEN() string {
	var fen []string
	fen = append(fen, s.Board.FEN())

	if s.ActiveColor == piece.White {
		fen = append(fen, "w")
	} else {
		fen = append(fen, "b")
	}

	var castling string
	if s.whiteCastling[0] {
		castling += "K"
	}
	if s.whiteCastling[1] {
		castling += "Q"
	}
	if s.blackCastling[0] {
		castling += "k"
	}
	if s.blackCastling[1] {
		castling += "q"
	}
	if castling == "" {
		castling = "-"
	}

	fen = append(fen, castling)
	fen = append(fen, s.EnPassantTarget)
	fen = append(fen, strconv.Itoa(s.halfmoveClock))
	fen = append(fen, strconv.Itoa(s.fullmoveNumber))

	return strings.Join(fen, " ")
}

func (s State) String() string {
	return s.FEN()
}

func (s *State) NextTurn() {
	s.ActiveColor *= -1
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
}

func (s State) ActivePlayer() player.Player {
	return s.Players[s.ActiveColor]
}

func (s State) PlayerRepr() string {
	switch s.ActiveColor {
	case piece.White:
		return "white"
	case piece.Black:
		return "black"
	default:
		log.Fatalf("s.Player must be %d or %d", piece.White, piece.Black)
		return ""
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func playerStateMsg(p player.Player) []any {
	msg := []any{}
	for k, v := range p.State() {
		msg = append(msg, k)
		msg = append(msg, v)
	}
	return msg
}

func (s State) Print() {
	clearScreen()
	s.Board.Print()
	fmt.Println(s)
}

func (s *State) handleEnPassantAvailable(m move) {
	isDoubleMove := int(m.targetRank-m.srcRank)*int(m.srcColor) == 2

	if m.srcValue == piece.Pawn && isDoubleMove {
		enPassantSquare := CoordsToSquare(int(m.targetFile), m.targetRank-int(m.srcColor))
		s.EnPassantTarget = enPassantSquare
	} else {
		s.EnPassantTarget = noEnPassantTarget
	}
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
}

func (s *State) handleEnPassantCapture(m move) {
	if m.srcValue == piece.Pawn && m.target == s.EnPassantTarget {
		capturedSquare := AddRank(m.target, -1*int(m.srcColor))
		capturedPiece := s.Board.Square(capturedSquare)

		assert.Assert(
			piece.Value(capturedPiece) == piece.Pawn && piece.IsColor(capturedPiece, m.srcColor*-1),
			fmt.Sprintf("handleEnPassantCapture: invalid capture: %s %s", m.src, m.target),
		)

		s.nextBoard[squareToIndex(capturedSquare)] = piece.None
		assert.AddContext("FEN", s.FEN())
		assert.AddContext("moves", s.Moves)
	}
}

func (s *State) handlePromotion(m move) {
	if m.srcValue == piece.Pawn && m.targetRank == piece.MaxPawnRank[m.srcColor] {
		p := s.ActivePlayer().ChoosePromotionPiece(m.target)
		s.nextBoard[squareToIndex(m.target)] = p * m.srcColor
	}
}

func (s *State) MakeMove(src, target string) {
	m := newMove(*s, src, target)
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.AddContext("move", m)

	s.Moves = append(s.Moves, m)
	s.handleEnPassantCapture(m)
	s.handleEnPassantAvailable(m)
	s.nextBoard.MakeMove(m.src, m.target)
	s.handlePromotion(m)
	s.Board = s.nextBoard
	if piece.IsColor(s.ActiveColor, piece.Black) {
		s.fullmoveNumber++
	}

	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.DeleteContext("move")
}
