package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

const noEnPassantTarget = "-"

type State struct {
	Players         map[piece.Piece]player.Player
	Castling        map[piece.Piece][2]bool
	EnPassantTarget string
	Moves           []move
	fens            []string
	Board           board.Chessboard
	nextBoard       board.Chessboard
	ActiveColor     piece.Piece
	HalfmoveClock   int
	FullmoveNumber  int
}

func (s State) DeepCopy() *State {
	n := &State{Players: make(map[piece.Piece]player.Player), Castling: make(map[piece.Piece][2]bool)}
	for color, player := range s.Players {
		n.Players[color] = player
	}
	for color, castlingRights := range s.Castling {
		n.Castling[color] = castlingRights
	}
	n.EnPassantTarget = s.EnPassantTarget
	n.Moves = append(n.Moves, s.Moves...)
	n.Board = s.Board
	n.nextBoard = s.nextBoard
	n.ActiveColor = s.ActiveColor
	n.HalfmoveClock = s.HalfmoveClock
	n.FullmoveNumber = s.FullmoveNumber

	return n
}

func StartingState(white, black player.Player) *State {
	fen := board.StartingFEN

	s := &State{
		Players: map[piece.Piece]player.Player{
			piece.White: white,
			piece.Black: black,
		},
	}

	s.LoadFEN(fen)

	return s
}

func StartingStateFromFEN(fen string, white, black player.Player) *State {
	s := &State{
		Players: map[piece.Piece]player.Player{
			piece.White: white,
			piece.Black: black,
		},
	}

	s.LoadFEN(fen)

	return s
}
func (s *State) LoadFEN(fen string) {
	fenFields := strings.Fields(fen)

	s.Board = board.LoadFEN(fenFields[0])
	s.nextBoard = s.Board

	var activeColor piece.Piece
	switch fenFields[1] {
	case "w":
		activeColor = piece.White
	case "b":
		activeColor = piece.Black
	default:
		assert.Raise(fmt.Sprintf("invalid FEN active color: %s", fen))
	}
	s.ActiveColor = activeColor

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
	s.Castling = map[piece.Piece][2]bool{
		piece.White: whiteCastling,
		piece.Black: blackCastling,
	}

	s.EnPassantTarget = fenFields[3]

	halfmoveClock, err := strconv.Atoi(fenFields[4])
	assert.ErrIsNil(err, fmt.Sprintf("invalid FEN halfmove clock: %s", fen))
	s.HalfmoveClock = halfmoveClock

	fullmoveNumber, err := strconv.Atoi(fenFields[5])
	assert.ErrIsNil(err, fmt.Sprintf("invalid FEN fullmove clock: %s", fen))
	s.FullmoveNumber = fullmoveNumber

	s.fens = append(s.fens, fen)

	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
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
	if s.Castling[piece.White][0] {
		castling += "K"
	}
	if s.Castling[piece.White][1] {
		castling += "Q"
	}
	if s.Castling[piece.Black][0] {
		castling += "k"
	}
	if s.Castling[piece.Black][1] {
		castling += "q"
	}
	if castling == "" {
		castling = "-"
	}

	fen = append(fen, castling)
	fen = append(fen, s.EnPassantTarget)
	fen = append(fen, strconv.Itoa(s.HalfmoveClock))
	fen = append(fen, strconv.Itoa(s.FullmoveNumber))

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

func (s State) ActivePlayerRepr() string {
	return fmt.Sprintf("%s (%s)", s.ActivePlayer(), piece.ColorToRepr[s.ActiveColor])
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
	// fmt.Println(s)
}

func (s *State) handleEnPassantAvailable(m move) {
	isDoubleMove := int(m.targetRank-m.srcRank)*int(m.srcColor) == 2

	if m.srcValue == piece.Pawn && isDoubleMove {
		enPassantSquare := board.CoordsToSquare(int(m.targetFile), m.targetRank-int(m.srcColor))
		s.EnPassantTarget = enPassantSquare
	} else {
		s.EnPassantTarget = noEnPassantTarget
	}
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
}

func (s *State) handleEnPassantCapture(m move) {
	if m.srcValue == piece.Pawn && m.target == s.EnPassantTarget {
		capturedSquare := board.AddRank(m.target, -1*int(m.srcColor))
		capturedPiece := s.Board.Square(capturedSquare)

		assert.Assert(
			piece.Value(capturedPiece) == piece.Pawn && piece.IsColor(capturedPiece, m.srcColor*-1),
			fmt.Sprintf("handleEnPassantCapture: invalid capture: %s %s", m.src, m.target),
		)

		s.nextBoard[board.SquareToIndex(capturedSquare)] = piece.Empty
		assert.AddContext("FEN", s.FEN())
		assert.AddContext("moves", s.Moves)
	}
}

func (s *State) handlePromotion(m move) {
	if m.srcValue == piece.Pawn && m.targetRank == piece.MaxPawnRank[m.srcColor] {
		p := s.ActivePlayer().ChoosePromotionPiece(m.target)
		s.nextBoard[board.SquareToIndex(m.target)] = p * m.srcColor
	}
}

func (s *State) handleUpdateCastlingRights(m move) {
	assert.AddContext("move", m)

	// rook movement
	for color, startingSquares := range piece.StartingRookSquares {
		castlingRights, ok := s.Castling[color]
		assert.Assert(ok, fmt.Sprintf("invalid castling rights: color %d not found: %v", color, s.Castling))

		for i, square := range startingSquares {
			if s.nextBoard.Square(square) != piece.Rook*color {
				castlingRights[i] = false
			}
		}

		s.Castling[color] = castlingRights
	}

	// king movement
	if m.srcValue == piece.King {
		s.Castling[m.srcColor] = [2]bool{false, false}
	}
}

func (s *State) handleCastle(m move) {
	// castling occurs
	if m.srcValue == piece.King && m.src == piece.StartingKingSquares[m.srcColor] {

		for i, castlingTarget := range piece.CastlingSquares[m.srcColor] {
			if m.target == castlingTarget {
				rookSrc := piece.StartingRookSquares[m.srcColor][i]
				rookTarget := piece.RookCastlingSquares[m.srcColor][i]
				assert.Assert(s.Board.Square(rookSrc) == piece.Rook*m.srcColor, "no rook found when castling")
				s.nextBoard.MakeMove(rookSrc, rookTarget)
			}
		}

	}
}

func (s *State) MakeMove(src, target string) {
	m := newMove(*s, src, target)
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.AddContext("move", m)

	isCapture := s.Board.Square(target) != piece.Empty

	s.Moves = append(s.Moves, m)
	s.handleEnPassantCapture(m)
	s.handleEnPassantAvailable(m)
	s.nextBoard.MakeMove(m.src, m.target)
	s.handleCastle(m)
	s.handleUpdateCastlingRights(m)
	s.handlePromotion(m)
	s.Board = s.nextBoard

	if piece.IsColor(s.ActiveColor, piece.Black) {
		s.FullmoveNumber++
	}

	if piece.Value(m.srcPiece) != piece.Pawn && !isCapture {
		s.HalfmoveClock++
	} else {
		s.HalfmoveClock = 0
	}

	s.fens = append(s.fens, s.FEN())

	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.DeleteContext("move")
}

func (s *State) Undo() {
	assert.Assert(len(s.fens) > 1, fmt.Sprintf("cannot undo move? there are %d fens", len(s.fens)))

	// remove the last TWO fens (current turn and previous turn)
	prevFEN := s.fens[len(s.fens)-2]
	s.fens = s.fens[:len(s.fens)-2]
	s.Moves = s.Moves[:len(s.Moves)-1]

	// this will add prevFEN back to s.fens
	s.LoadFEN(prevFEN)
}
