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
	"github.com/ethansaxenian/chess/player"
	"github.com/ethansaxenian/chess/state"
	"github.com/ethansaxenian/chess/tui"
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
	if res, over := state.CheckGameOver(); over {
		fmt.Println(res)
		os.Exit(0)
	}

	state.Print()
	possibleMoves := state.GeneratePossibleMoves()
	assert.AddContext("possible moves", possibleMoves)
	assert.AddContext("FEN", state.FEN())
	assert.AddContext("moves", state.Moves)

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
	white := player.NewRandoBot()
	black := player.NewRandoBot()

	tui.RunTUI(white, black)
}
