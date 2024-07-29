package board

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareToIndex(t *testing.T) {
	assert.Equal(t, 0, SquareToIndex("a1"))
	assert.Equal(t, 1, SquareToIndex("b1"))
	assert.Equal(t, 8, SquareToIndex("a2"))
	assert.Equal(t, 63, SquareToIndex("h8"))
}

func TestIndexToSquare(t *testing.T) {
	assert.Equal(t, "a1", indexToSquare(0))
	assert.Equal(t, "b1", indexToSquare(1))
	assert.Equal(t, "a2", indexToSquare(8))
	assert.Equal(t, "h8", indexToSquare(63))
}

func TestSquareToCoords(t *testing.T) {
	f, r := SquareToCoords("a1")
	assert.Equal(t, 97, f)
	assert.Equal(t, 1, r)
}

func TestCoordsToSquare(t *testing.T) {
	s := CoordsToSquare(97, 1)
	assert.Equal(t, "a1", s)
}

func TestBoardToFen(t *testing.T) {
	fen := strings.Fields(StartingFEN)[0]
	b := LoadFEN(fen)
	assert.Equal(t, fen, b.FEN())

	b.MakeMove("e2", "e4")
	newFen := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR"
	assert.Equal(t, newFen, b.FEN())
}
