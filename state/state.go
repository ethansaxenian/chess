package state

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

const noEnPassantTarget = "-"

type gameOverState int

const (
	no gameOverState = iota
	whiteWin
	blackWin
	stalemate
	draw
)

func (g gameOverState) String() string {
	switch g {
	case whiteWin:
		return "white wins!"
	case blackWin:
		return "black wins!"
	case stalemate:
		return "stalemate"
	case draw:
		return "draw"
	default:
		assert.Raise("The game is not over?")
		return ""
	}
}

type State struct {
	Players         map[piece.Piece]player.Player
	Castling        map[piece.Piece]map[piece.Side]bool
	EnPassantTarget string
	Moves           []move.Move
	fens            []string
	Board           board.Chessboard
	nextBoard       board.Chessboard
	ActiveColor     piece.Piece
	HalfmoveClock   int
	FullmoveNumber  int
	headless        bool
}

func StartingState(white, black player.Player) *State {
	return StartingStateFromFEN(board.StartingFEN, white, black)
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

	whiteCastling := map[piece.Side]bool{}
	blackCastling := map[piece.Side]bool{}
	for _, char := range fenFields[2] {
		switch char {
		case 'K':
			whiteCastling[piece.Kingside] = true
		case 'Q':
			whiteCastling[piece.Queenside] = true
		case 'k':
			blackCastling[piece.Kingside] = true
		case 'q':
			blackCastling[piece.Queenside] = true
		case '-':
			continue
		default:
			assert.Raise(fmt.Sprintf("invalid FEN castling rights: %s", fen))
		}
	}
	s.Castling = map[piece.Piece]map[piece.Side]bool{
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
	if s.Castling[piece.White][piece.Kingside] {
		castling += "K"
	}
	if s.Castling[piece.White][piece.Queenside] {
		castling += "Q"
	}
	if s.Castling[piece.Black][piece.Kingside] {
		castling += "k"
	}
	if s.Castling[piece.Black][piece.Queenside] {
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

func (s State) Piece(square string) piece.Piece {
	return s.Board.Square(square)
}

func (s State) String() string {
	assert.Assert(s.Board == s.nextBoard, "dasfsdg")
	return s.FEN()
}

func (s State) ActivePlayer() player.Player {
	return s.Players[s.ActiveColor]
}

func (s State) PlayerRepr(color piece.Piece) string {
	var colorStr string
	switch color {
	case piece.White:
		colorStr = "white"
	case piece.Black:
		colorStr = "black"
	}

	return fmt.Sprintf("%s (%s)", s.Players[color], colorStr)
}

func (s State) ActivePlayerRepr() string {
	return s.PlayerRepr(s.ActiveColor)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (s State) Print() {
	if s.headless {
		return
	}

	clearScreen()
	s.Board.Print()
	fmt.Println(s)
}

func (s *State) handleEnPassantAvailable(_ move.Move, mc moveContext) {
	s.EnPassantTarget = mc.nextEnPassantTarget
	assert.AddContext("FEN", s.FEN())
}

func (s *State) handleEnPassantCapture(_ move.Move, mc moveContext) {
	if mc.enPassantCapture != "" {
		s.nextBoard[board.SquareToIndex(mc.enPassantCapture)] = piece.Empty
		assert.AddContext("FEN", s.FEN())
	}
}

func (s *State) handlePromotion(m move.Move, mc moveContext) {
	if mc.isPromotion {
		p := s.ActivePlayer().ChoosePromotionPiece(m.Target)
		s.nextBoard[board.SquareToIndex(m.Target)] = p * s.Piece(m.Source).Color()
		assert.AddContext("FEN", s.FEN())
	}
}

func (s *State) handleUpdateCastlingRights(m move.Move, _ moveContext) {
	assert.AddContext("move", m)

	// rook movement
	for color, startingSquares := range piece.StartingRookSquares {
		castlingRights, ok := s.Castling[color]
		assert.Assert(ok, fmt.Sprintf("invalid castling rights: color %d not found: %v", color, s.Castling))

		for side, square := range startingSquares {
			if s.nextBoard.Square(square) != piece.Rook*color {
				castlingRights[side] = false
			}
		}

		s.Castling[color] = castlingRights
	}

	// king movement
	if s.Piece(m.Source).Type() == piece.King {
		color := s.Piece(m.Source).Color()
		s.Castling[color][piece.Kingside] = false
		s.Castling[color][piece.Queenside] = false
	}
}

func (s *State) handleCastle(m move.Move, mc moveContext) {
	// castling occurs
	if mc.castling != nil {
		color := s.Piece(m.Source).Color()
		rookSource := piece.StartingRookSquares[color][mc.castling.side]
		rookTarget := piece.RookCastlingSquares[color][mc.castling.side]
		assert.Assert(s.Piece(rookSource) == piece.Rook*color, "no rook found when castling")
		s.nextBoard.MakeMove(move.NewMove(rookSource, rookTarget))
	}
}

func (s *State) MakeMove(m move.Move) {
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.AddContext("move", m)

	mc := getMoveContext(*s, m)

	var isPawnMove = s.Piece(m.Source).Type() == piece.Pawn

	s.nextBoard.MakeMove(m)
	s.handleEnPassantCapture(m, mc)
	s.handleCastle(m, mc)
	s.handlePromotion(m, mc)
	s.handleEnPassantAvailable(m, mc)
	s.handleUpdateCastlingRights(m, mc)
	s.Board = s.nextBoard
	s.Moves = append(s.Moves, m)

	if s.ActiveColor.Color() == piece.Black {
		s.FullmoveNumber++
	}

	if isPawnMove || mc.isCapture {
		s.HalfmoveClock = 0
	} else {
		s.HalfmoveClock++
	}

	s.ActiveColor *= -1
	s.fens = append(s.fens, s.FEN())
	assert.AddContext("FEN", s.FEN())
	assert.AddContext("moves", s.Moves)
	assert.DeleteContext("move")
}

func (s *State) Undo() {
	numFens := len(s.fens)

	assert.Assert(numFens > 1, fmt.Sprintf("cannot undo move? there are %d fens", len(s.fens)))

	assert.AddContext("last recorded FEN", s.fens[numFens-1])
	assert.AddContext("current FEN", s.FEN())
	assert.Assert(s.fens[numFens-1] == s.FEN(), "Undo: FEN mismatch")
	assert.DeleteContext("last recorded FEN")
	assert.DeleteContext("current FEN")

	index := numFens - 2
	prevFEN := s.fens[index]
	s.fens = s.fens[:index]

	// this adds prevFEN back to s.fens
	s.LoadFEN(prevFEN)
}

func (s *State) PlayMoves(moves []string) {
	for _, m := range moves {
		s.MakeMove(move.NewMove(m[:2], m[2:]))
	}
}

func (s State) GeneratePossibleMoves() []move.Move {
	moves := []move.Move{}

	for _, m := range generateTmpMoves(s) {
		s.MakeMove(m)

		var capturedKing bool
		for _, nextMove := range generateTmpMoves(s) {
			if s.Piece(nextMove.Target) == piece.King*s.ActiveColor*-1 {
				capturedKing = true
				break
			}
		}

		if !capturedKing {
			moves = append(moves, m)
		}

		s.Undo()

	}

	return moves
}

func (s *State) IsCheck() bool {
	s.ActiveColor *= -1

	var check bool
	for _, m := range s.GeneratePossibleMoves() {
		if s.Piece(m.Target) == piece.King*s.ActiveColor*-1 {
			check = true
			break
		}
	}
	s.ActiveColor *= -1

	return check
}

func (s *State) CheckGameOver() (gameOverState, bool) {
	validMoves := s.GeneratePossibleMoves()
	if len(validMoves) == 0 {
		if s.IsCheck() {
			switch s.ActiveColor {
			case piece.White:
				return blackWin, true
			case piece.Black:
				return whiteWin, true
			default:
				return no, false
			}
		} else {
			return stalemate, true
		}
	} else if s.HalfmoveClock >= 100 {
		return draw, true
	} else {
		return no, false
	}
}
