package board

import (
	"fmt"
	"log"

	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

type State struct {
	Players     map[piece.Piece]player.Player
	Moves       []string
	Board       Chessboard
	ActiveColor piece.Piece
}

func StartingState(white, black player.Player) State {
	return State{
		Board:       LoadFEN(StartingFEN),
		Players:     map[piece.Piece]player.Player{piece.White: white, piece.Black: black},
		ActiveColor: piece.White,
		Moves:       []string{},
	}
}

func (s *State) NextTurn() {
	s.ActiveColor *= -1
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
	// slog.Info("white", playerStateMsg(s.Players[piece.White])...)
	// slog.Info("black", playerStateMsg(s.Players[piece.Black])...)
	// log.Println(s.Moves)
}

func (s *State) MakeMove(src, target string) {
	srcPiece := piece.Value(s.Board.Square(src))
	targetPiece := s.Board.Square(target)
	repr := piece.PieceToRepr[srcPiece]
	if targetPiece != piece.None {
		repr += "x"
	}
	repr += target
	s.Moves = append(s.Moves, repr)
	s.Board.MakeMove(src, target)
}
