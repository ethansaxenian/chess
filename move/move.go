package move

import (
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/piece"
)

var precomputedPieceMoves = map[piece.Piece]map[string]map[string]bool{}

func init() {
	for _, p := range piece.AllPieces {
		for _, c := range piece.AllColors {
			pieceMap := map[string]map[string]bool{}
			for _, sf := range board.Files {
				for _, sr := range board.Ranks {
					src := string(sf) + string(sr)
					pieceMap[src] = map[string]bool{}
					for _, tf := range board.Files {
						for _, tr := range board.Ranks {
							target := string(tf) + string(tr)
							pieceMap[src][target] = validatePieceMovement(p*c, src, target)
						}
					}
				}
			}
			precomputedPieceMoves[p*c] = pieceMap
		}
	}
}

func GeneratePossibleMoves(state board.State) [][2]string {
	moves := [][2]string{}

	for src, p := range state.Board.Squares() {
		if p == piece.None {
			continue
		}

		for target, valid := range precomputedPieceMoves[p][src] {
			if valid && piece.IsColor(p, state.CurrColor) {
				moves = append(moves, [2]string{src, target})
			}
		}
	}

	return moves
}
