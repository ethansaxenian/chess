package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime/pprof"
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
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var useTUI = flag.Bool("tui", false, "use the bubbletea tui")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	initLogger(*logLevel)

	// white := player.NewHumanPlayer("human")
	// black := player.NewHumanPlayer("human")
	white := player.NewRandoBot()
	black := player.NewRandoBot()

	if *useTUI {
		tui.RunTUI(white, black)
	} else {
		s := state.StartingState(white, black)
		for {
			mainLoop(s)
		}
	}
}
