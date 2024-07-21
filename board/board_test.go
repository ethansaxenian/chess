package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareToIndex(t *testing.T) {
	assert.Equal(t, SquareToIndex("a1"), 0)
	assert.Equal(t, SquareToIndex("b1"), 1)
	assert.Equal(t, SquareToIndex("a2"), 8)
	assert.Equal(t, SquareToIndex("h8"), 63)
}
