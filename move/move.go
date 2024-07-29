package move

import (
	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/board"
	"github.com/ethansaxenian/chess/game"
	"github.com/ethansaxenian/chess/piece"
)

var precomputedPieceMoves = map[piece.Piece]map[string][]string{}

func init() {
	for _, p := range piece.AllPieces {
		for _, c := range piece.AllColors {
			pieceMap := map[string][]string{}
			for _, sf := range board.Files {
				for _, sr := range board.Ranks {
					src := string(sf) + string(sr)
					assert.Assert(src != "e0", "foobar")
					for _, tf := range board.Files {
						for _, tr := range board.Ranks {
							target := string(tf) + string(tr)
							if validatePieceMove(p*c, src, target) {
								pieceMap[src] = append(pieceMap[src], target)
							}
						}
					}
				}
			}
			precomputedPieceMoves[p*c] = pieceMap
		}
	}
}

func GeneratePossibleMoves(state game.State) [][2]string {
	moves := [][2]string{}

	for src, p := range state.Board.Squares() {
		if p == piece.Empty {
			continue
		}

		for _, target := range precomputedPieceMoves[p][src] {
			if validateMove(state, src, target) {
				moves = append(moves, [2]string{src, target})
			}
		}
	}

	return moves
}
