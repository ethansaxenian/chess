package player

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ethansaxenian/chess/piece"
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

		if len(input) != 4 {
			continue
		}

		if slices.Contains(validMoves, [2]string{input[:2], input[2:]}) {
			return input[:2], input[2:]
		}
	}
}

func (h HumanPlayer) State() map[string]any {
	return map[string]any{"name": h.name}
}

func (h HumanPlayer) ChoosePromotionPiece(square string) piece.Piece {
	fmt.Printf("Promote %c on %s\n", piece.PieceToChar[piece.Pawn], square)
	for i, p := range piece.PossiblePromotions {
		fmt.Printf("%d: %c\n", i, piece.PieceToChar[p])
	}
	for {
		input := getInput()

		i, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		if i >= 0 && i < len(piece.PossiblePromotions) {
			return piece.PossiblePromotions[i]
		}
	}
}

func (h HumanPlayer) String() string {
	return h.name
}
