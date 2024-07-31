package main

import (
	"testing"

	"github.com/ethansaxenian/chess/player"
	"github.com/ethansaxenian/chess/state"
	"github.com/stretchr/testify/assert"
)

func testMainLoop(t *testing.T) {
	white := player.NewRandoBot(player.WithSeed(10))
	black := player.NewRandoBot(player.WithSeed(10))

	s := state.StartingState(white, black)

	for i := 0; i < 20; i++ {
		mainLoop(s)
	}

	firstFEN := s.FEN()

	s = state.StartingState(white, black)
	for i := 0; i < 20; i++ {
		mainLoop(s)
	}

	secondFEN := s.FEN()

	assert.Equal(t, firstFEN, secondFEN)
}
