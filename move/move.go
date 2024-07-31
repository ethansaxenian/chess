package move

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ethansaxenian/chess/assert"
)

type Move struct {
	Source, Target string
}

func (m Move) SourceRank() int {
	srcRank, err := strconv.Atoi(string(m.Source[1]))
	assert.ErrIsNil(err, fmt.Sprintf("move: invalid src: %s", m.Source))
	return srcRank
}

func (m Move) TargetRank() int {
	targetRank, err := strconv.Atoi(string(m.Target[1]))
	assert.ErrIsNil(err, fmt.Sprintf("move: invalid target: %s", m.Target))
	return targetRank
}

func (m Move) SourceFile() byte {
	return m.Source[0]
}

func (m Move) TargetFile() byte {
	return m.Target[0]
}

func NewMove(source, target string) Move {
	return Move{Source: source, Target: target}
}

func (m Move) String() string {
	return m.Source + m.Target
}

func SortMoves(moves []Move) {
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].String() < moves[j].String()
	})
}
