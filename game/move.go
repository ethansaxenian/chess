package game

import (
	"fmt"
	"strconv"

	"github.com/ethansaxenian/chess/assert"
	"github.com/ethansaxenian/chess/piece"
)

type move struct {
	src, target           string
	srcRank, targetRank   int
	srcFile, targetFile   byte
	srcPiece, targetPiece piece.Piece
	srcValue, targetValue piece.Piece
	srcColor, targetColor piece.Piece
}

func newMove(s State, src, target string) move {
	srcPiece := s.Board.Square(src)
	targetPiece := s.Board.Square(target)

	srcRank, err := strconv.Atoi(string(src[1]))
	assert.ErrIsNil(err, fmt.Sprintf("newMove: invalid src: %s", src))

	targetRank, err := strconv.Atoi(string(target[1]))
	assert.ErrIsNil(err, fmt.Sprintf("newMove: invalid target: %s", target))

	return move{
		src:         src,
		target:      target,
		srcRank:     srcRank,
		targetRank:  targetRank,
		srcFile:     src[0],
		targetFile:  target[0],
		srcPiece:    srcPiece,
		targetPiece: targetPiece,
		srcValue:    piece.Value(srcPiece),
		targetValue: piece.Value(targetPiece),
		srcColor:    piece.Color(srcPiece),
		targetColor: piece.Color(targetPiece),
	}
}

func (m move) String() string {
	// srcPiece := piece.Value(b.Square(src))
	// targetPiece := b.Square(target)
	// repr := piece.PieceToRepr[srcPiece]
	// if targetPiece != piece.None {
	// 	repr += "x"
	// }
	// repr += target
	//
	// return repr
	return m.src + m.target
}
