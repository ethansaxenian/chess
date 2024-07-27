package board

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

type State struct {
	Players   map[piece.Piece]player.Player
	Board     chessboard
	CurrColor piece.Piece
}

func StartingState(white, black player.Player) State {
	return State{
		Board:     LoadFEN(StartingFEN),
		Players:   map[piece.Piece]player.Player{piece.White: white, piece.Black: black},
		CurrColor: piece.White,
	}
}

func (s *State) NextTurn() {
	s.CurrColor *= -1
}

func (s State) CurrPlayer() player.Player {
	return s.Players[s.CurrColor]
}

func (s State) PlayerRepr() string {
	switch s.CurrColor {
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
	slog.Info("white", playerStateMsg(s.Players[piece.White])...)
	slog.Info("black", playerStateMsg(s.Players[piece.Black])...)
}
