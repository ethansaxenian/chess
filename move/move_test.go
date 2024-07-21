package move

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSquareValid(t *testing.T) {
	assert.True(t, validateSquare("a1"))
	assert.True(t, validateSquare("h8"))
	assert.True(t, validateSquare("d4"))
}

func TestValidateSquareInvalid(t *testing.T) {
	assert.False(t, validateSquare("a0"))
	assert.False(t, validateSquare("h9"))
	assert.False(t, validateSquare("i4"))
	assert.False(t, validateSquare("1a"))
	assert.False(t, validateSquare("a"))
	assert.False(t, validateSquare(""))
	assert.False(t, validateSquare("foo"))
}

func TestValidateMoveValid(t *testing.T) {
	assert.True(t, validateMove("a1f1"))
}

func TestValidateMoveInvalid(t *testing.T) {
	assert.False(t, validateMove("a1"))
	assert.False(t, validateMove("a1f1g1"))
	assert.False(t, validateMove("a1 f1"))
}
