package game

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

const noEnPassantTarget = "-"

type State struct {
	Players         map[piece.Piece]player.Player
	Castling        map[piece.Piece][2]bool
	EnPassantTarget string
	Moves           []move
	Board           Chessboard
	nextBoard       Chessboard
	ActiveColor     piece.Piece
	halfmoveClock   int
	fullmoveNumber  int
}

func StartingState(white, black player.Player) *State {
	s := &State{Players: map[piece.Piece]player.Player{piece.White: white, piece.Black: black}}

	s.LoadFEN(StartingFEN)

	return s
}

func (s *State) LoadFEN(fen string) {
	fenFields := strings.Fields(fen)

	s.Board = loadFEN(fenFields[0])
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
	s.halfmoveClock = halfmoveClock

	fullmoveNumber, err := strconv.Atoi(fenFields[5])
	assert.ErrIsNil(err, fmt.Sprintf("invalid FEN fullmove clock: %s", fen))
	s.fullmoveNumber = fullmoveNumber

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

		s.nextBoard[squareToIndex(capturedSquare)] = piece.Empty
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

func (s *State) handleGameEnd(m move) {
	if m.targetPiece == piece.King {
		s.Print()
		fmt.Printf("%s wins\n", s.ActivePlayerRepr())
		fmt.Println(s.Moves)
		os.Exit(0)
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
	s.handleCastle(m)
	s.handleUpdateCastlingRights(m)
	s.handlePromotion(m)
	s.Board = s.nextBoard
	if piece.IsColor(s.ActiveColor, piece.Black) {
		s.fullmoveNumber++
	}
	s.handleGameEnd(m)

	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.DeleteContext("move")
}
