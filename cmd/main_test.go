package main

import (
	"testing"

	"github.com/ethansaxenian/chess/game"
	"github.com/ethansaxenian/chess/player"
	"github.com/stretchr/testify/assert"
)

func testMainLoop(t *testing.T) {
	white := player.NewRandoBot(player.WithSeed(10))
	black := player.NewRandoBot(player.WithSeed(10))

	state := game.StartingState(white, black)

	for i := 0; i < 20; i++ {
		mainLoop(state)
	}

	firstFEN := state.FEN()

	state = game.StartingState(white, black)
	for i := 0; i < 20; i++ {
		mainLoop(state)
	}

	secondFEN := state.FEN()

	assert.Equal(t, firstFEN, secondFEN)
}
