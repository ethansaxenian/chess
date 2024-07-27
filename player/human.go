package player

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type HumanPlayer struct {
	name string
}

func NewHumanPlayer(name string) HumanPlayer {
	return HumanPlayer{name}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>> ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	return input
}

func (h HumanPlayer) GetMove(validMoves [][2]string) (string, string) {
	for {
		input := getInput()

		if slices.Contains(validMoves, [2]string{input[:2], input[2:]}) {
			return input[:2], input[2:]
		}
	}
}

func (h HumanPlayer) State() map[string]any {
	return map[string]any{"name": h.name}
}
