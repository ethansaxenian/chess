package main

import (
	"fmt"
	"slices"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/player"
)

func main() {
	white := player.NewHumanPlayer("human")
	black := player.NewRandoBot(player.WithSeed(10))
	state := board.StartingState(white, black)

	for {
		state.Print()
		possibleMoves := move.GeneratePossibleMoves(state)

		src, target := state.CurrPlayer().GetMove(possibleMoves)
		assert.Assert(slices.Contains(possibleMoves, [2]string{src, target}), fmt.Sprintf("%s,%s not in possibleMoves", src, target))
		state.Board.MakeMove(src, target)

		state.NextTurn()
	}
}
