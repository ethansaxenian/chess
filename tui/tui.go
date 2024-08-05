package main

import (
	"fmt"
	"slices"

	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/player"
	"github.com/ethansaxenian/chess/state"
)

type botTurnMsg struct{}

func botTurn() tea.Msg {
	return botTurnMsg{}
}

type model struct {
	*state.State
	validMoves []move.Move
	input      textinput.Model
	check      bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 4
	ti.Width = 4

	// s := state.StartingState(player.NewHumanPlayer("human"), player.NewHumanPlayer("foo"))
	white := player.NewRandoBot(player.WithSeed(10))
	black := player.NewRandoBot(player.WithSeed(1722405359723887000))
	s := state.StartingState(white, black)
	return model{
		State: s,
		// State: state.StartingState(player.NewRandoBot(), player.NewHumanPlayer("")),
		input:      ti,
		validMoves: s.GeneratePossibleMoves(),
		check:      false,
	}
}

func (m model) Init() tea.Cmd {
	if m.ActivePlayer().IsBot() {
		return tea.Batch(botTurn, textinput.Blink)
	}

	return textinput.Blink
}

func (m model) View() string {
	view := m.Board.String() + m.FEN() + "\n"

	view += fmt.Sprintf("%s to play\n\n", m.ActivePlayerRepr())

	if m.check {
		view += "check!\n\n"
	}

	view += m.input.View()

	return view
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case botTurnMsg:
		return m.getBotMove()

	case move.Move:
		return m.onMove(msg)

	case tea.KeyMsg:
		return m.onKey(msg)
	}

	return m, nil
}

func (m model) onKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {

	case tea.KeyCtrlC:
		return m, tea.Quit

	case tea.KeyEnter:
		return m.onEnter()

	default:
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd

	}

}

func (m model) getBotMove() (tea.Model, tea.Cmd) {
	mv := m.ActivePlayer().GetMove(m.validMoves)
	return m, func() tea.Msg { return mv }
}

func (m model) onEnter() (tea.Model, tea.Cmd) {
	val := m.input.Value()

	if len(val) == 4 {
		mv := move.NewMove(val[:2], val[2:])

		if slices.Contains(m.validMoves, mv) {
			return m, func() tea.Msg { return mv }
		}
	}

	m.input.Reset()
	return m, nil
}

func (m *model) onMove(mv move.Move) (tea.Model, tea.Cmd) {
	m.MakeMove(mv)
	m.input.Reset()

	if m.HalfmoveClock == 100 {
		return m, tea.Quit
	}

	m.validMoves = m.GeneratePossibleMoves()
	m.check = m.IsCheck()

	if len(m.validMoves) == 0 && m.check {
		return m, tea.Quit
	}

	if m.ActivePlayer().IsBot() {
		return m, botTurn
	}
	return m, nil
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	fmt.Println(m.check)
}
