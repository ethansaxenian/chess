package move

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMoves(t *testing.T) {
	moves := []Move{NewMove("h8", "a1"), NewMove("a1", "h8")}
	SortMoves(moves)
	expected := []Move{NewMove("a1", "h8"), NewMove("h8", "a1")}
	assert.Equal(t, expected, moves)
}
