package main

import (
	"testing"

	"github.com/ethansaxenian/chess/player"
	"github.com/ethansaxenian/chess/state"
	"github.com/stretchr/testify/assert"
)

func TestMainLoop(t *testing.T) {
	white := player.NewRandoBot(player.WithSeed(10))
	black := player.NewRandoBot(player.WithSeed(10))

	s := state.NewStartingTestState(white, black)

	for i := 0; i < 1; i++ {
		mainLoop(s)
	}

	firstFEN := s.FEN()

	white = player.NewRandoBot(player.WithSeed(10))
	black = player.NewRandoBot(player.WithSeed(10))

	s = state.NewStartingTestState(white, black)

	for i := 0; i < 1; i++ {
		mainLoop(s)
	}

	secondFEN := s.FEN()

	assert.Equal(t, firstFEN, secondFEN)
}
