package move

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethansaxenian/chess/board"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a move (e.g. e2e4): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	return input
}

func GetMove() (string, string) {
	for {
		input := getInput()

		if validateMove(input) {
			return input[:2], input[2:]
		}
	}
}

func validateMove(move string) bool {
	if len(move) != 4 {
		return false
	}

	if valid := validateSquare(move[:2]); !valid {
		return false
	}

	if valid := validateSquare(move[2:]); !valid {
		return false
	}

	return true
}

func validateSquare(square string) bool {
	if len(square) != 2 {
		return false
	}

	rank := strings.IndexByte(board.Ranks, square[1])
	if rank == -1 {
		return false
	}

	file := strings.IndexByte(board.Files, square[0])
	if file == -1 {
		return false
	}

	return true
}
