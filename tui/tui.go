package tui

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
	input textinput.Model
}

func initialModel(white, black player.Player) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 4
	ti.Width = 4

	return model{
		State: state.StartingState(white, black),
		input: ti,
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

	if m.IsCheck() && !m.ActivePlayer().IsBot() {
		view += "check!\n\n"
	}

	if !m.ActivePlayer().IsBot() {
		view += m.input.View()
	}

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
	validMoves := m.GeneratePossibleMoves()
	mv := m.ActivePlayer().GetMove(validMoves)
	return m, func() tea.Msg { return mv }
}

func (m model) onEnter() (tea.Model, tea.Cmd) {
	val := m.input.Value()

	if len(val) == 4 {
		mv := move.NewMove(val[:2], val[2:])

		validMoves := m.GeneratePossibleMoves()
		if slices.Contains(validMoves, mv) {
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

	validMoves := m.GeneratePossibleMoves()

	if len(validMoves) == 0 && m.IsCheck() {
		return m, tea.Quit
	}

	if m.ActivePlayer().IsBot() {
		return m, botTurn
	}
	return m, nil
}

func RunTUI(white, black player.Player) {
	m := initialModel(white, black)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if res, over := m.CheckGameOver(); over {
		fmt.Println(res)
	}
}
