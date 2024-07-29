package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/game"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
)

func main() {
	white := player.NewHumanPlayer("human")
	// black := player.NewHumanPlayer("human")
	// white := player.NewRandoBot()
	black := player.NewRandoBot()

	state := game.StartingState(white, black)

	for {
		state.Print()
		possibleMoves := move.GeneratePossibleMoves(*state)
		assert.AddContext("possible moves", possibleMoves)
		assert.AddContext("FEN", state.FEN())
		assert.AddContext("moves", state.Moves)

		for _, m := range possibleMoves {
			assert.Assert(piece.Value(state.Board.Square(m[1])) != piece.King, fmt.Sprintf("wtf: %s", m))
		}
		// fmt.Println(possibleMoves)
		if len(possibleMoves) == 0 {
			fmt.Println(state.ActivePlayerRepr(), "to play")
			state.ActiveColor *= -1

			var checkmate bool
			for _, m := range move.GeneratePossibleMoves(*state) {
				if state.Board.Square(m[1]) == piece.King*state.ActiveColor*-1 {
					checkmate = true
					break
				}
			}
			if checkmate {
				fmt.Println("checkmate!")
			} else {
				fmt.Println("draw!")
			}

			os.Exit(0)
		}

		if state.HalfmoveClock == 100 {
			fmt.Println("draw!")
			os.Exit(0)
		}

		src, target := state.ActivePlayer().GetMove(possibleMoves)
		assert.Assert(slices.Contains(possibleMoves, [2]string{src, target}), fmt.Sprintf("%s,%s not in possibleMoves", src, target))
		state.MakeMove(src, target)

		state.NextTurn()
		// time.Sleep(time.Millisecond * 100)
	}
}
