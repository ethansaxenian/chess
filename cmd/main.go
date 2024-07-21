package main

import (
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/move"
)

func main() {
	b := board.LoadFEN(board.StartingFEN)

	for {
		b.Print()

		src, dest := move.GetMove()
		b.MakeMove(src, dest)
	}
}
