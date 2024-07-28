package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/game"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/player"
)

func main() {
	// white := player.NewHumanPlayer("human")
	// black := player.NewHumanPlayer("human")
	white := player.NewRandoBot()
	black := player.NewRandoBot()

	state := game.StartingState(white, black)
	// state.LoadFEN("r3k2r/8/8/8/8/8/8/R3KB1R w KQkq - 0 1")

	for {
		state.Print()
		possibleMoves := move.GeneratePossibleMoves(*state)

		// fmt.Println(possibleMoves)
		src, target := state.ActivePlayer().GetMove(possibleMoves)
		assert.Assert(slices.Contains(possibleMoves, [2]string{src, target}), fmt.Sprintf("%s,%s not in possibleMoves", src, target))
		state.MakeMove(src, target)

		state.NextTurn()
		time.Sleep(time.Millisecond * 100)
	}
}
