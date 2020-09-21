package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChessBoard(t *testing.T) {
	board1 := NewChessBoard()
	board1.Move(1, 1, 2)
	board2 := NewChessBoard()
	board2.Move(2, 2, 3)

	assert.Equal(t, board1.chessPieces[1].Unit, board2.chessPieces[1].Unit)
	assert.Equal(t, board1.chessPieces[2].Unit, board2.chessPieces[2].Unit)
}
