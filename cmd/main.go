package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/piece"
	"github.com/ethansaxenian/chess/player"
	"github.com/ethansaxenian/chess/state"
)

func initLogger(value string) {
	var level slog.Level
	switch strings.ToLower(value) {
	case "error":
		level = slog.LevelError
	case "warning":
		level = slog.LevelWarn
	case "info":
		level = slog.LevelInfo
	case "debug":
		level = slog.LevelDebug
	default:
		log.Fatalf("invalid log level: %s\n", value)
	}

	h := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: level},
	)

	slog.SetDefault(slog.New(h))
}

func mainLoop(state *state.State) {
	state.Print()
	possibleMoves := state.GeneratePossibleMoves()
	assert.AddContext("possible moves", possibleMoves)
	assert.AddContext("FEN", state.FEN())
	assert.AddContext("moves", state.Moves)

	for _, m := range possibleMoves {
		assert.Assert(state.Piece(m.Target).Type() != piece.King, fmt.Sprintf("wtf: %s", m))
	}

	if len(possibleMoves) == 0 {
		fmt.Println(state.ActivePlayerRepr(), "to play")
		state.ActiveColor *= -1

		var checkmate bool
		for _, m := range state.GeneratePossibleMoves() {
			if state.Piece(m.Target) == piece.King*state.ActiveColor*-1 {
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

	m := state.ActivePlayer().GetMove(possibleMoves)
	assert.Assert(slices.Contains(possibleMoves, m), fmt.Sprintf("%s not in possibleMoves", m))
	state.MakeMove(m)
}

func main() {
	var logLevel = flag.String("log-level", "info", "set the log level (debug, info, warning, error)")
	flag.Parse()

	initLogger(*logLevel)

	// white := player.NewHumanPlayer("human")
	// black := player.NewHumanPlayer("human")
	white := player.NewRandoBot(player.WithSeed(10))
	black := player.NewRandoBot(player.WithSeed(1722405359723887000))

	state := state.StartingState(white, black)

	for {
		mainLoop(state)
	}
}
