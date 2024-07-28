package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareToIndex(t *testing.T) {
	assert.Equal(t, 0, squareToIndex("a1"))
	assert.Equal(t, 1, squareToIndex("b1"))
	assert.Equal(t, 8, squareToIndex("a2"))
	assert.Equal(t, 63, squareToIndex("h8"))
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
