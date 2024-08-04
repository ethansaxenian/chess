package player

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ethansaxenian/chess/move"
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

func (h HumanPlayer) GetMove(validMoves []move.Move) move.Move {
	for {
		input := getInput()

		if len(input) != 4 {
			continue
		}

		m := move.NewMove(input[:2], input[2:])
		if slices.Contains(validMoves, m) {
			return m
		}
	}
}

func (h HumanPlayer) State() map[string]any {
	return map[string]any{"name": h.name}
}

func (h HumanPlayer) ChoosePromotionPiece(square string) piece.Piece {
	// TODO: promotions
	return piece.Queen
	// fmt.Printf("Promote %s on %s\n", piece.Pawn, square)
	// for i, p := range piece.PossiblePromotions {
	// 	fmt.Printf("%d: %s\n", i, p)
	// }
	// for {
	// 	input := getInput()
	//
	// 	i, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		continue
	// 	}
	//
	// 	if i >= 0 && i < len(piece.PossiblePromotions) {
	// 		return piece.PossiblePromotions[i]
	// 	}
	// }
}

func (h HumanPlayer) String() string {
	return h.name
}
