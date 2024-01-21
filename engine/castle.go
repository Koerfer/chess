package engine

import (
	"chess/pieces"
)

func (e *Engine) castle(option int, board map[int]*pieces.Piece, selectedPiece *pieces.Piece, whitesTurn bool) bool {
	switch option {
	case 2:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[3] = board[0]
		delete(board, selectedPiece.LastPosition)
		delete(board, 0)
		selectedPiece = nil
		whitesTurn = !whitesTurn
		e.calculateAllPositionsNew(whitesTurn)
		return true
	case 6:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[5] = board[7]
		delete(board, selectedPiece.LastPosition)
		delete(board, 7)
		selectedPiece = nil
		whitesTurn = !whitesTurn
		e.calculateAllPositionsNew(whitesTurn)
		return true
	case 58:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[59] = board[56]
		delete(board, selectedPiece.LastPosition)
		delete(board, 56)
		selectedPiece = nil
		whitesTurn = !whitesTurn
		e.calculateAllPositionsNew(whitesTurn)
		return true
	case 62:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[61] = board[63]
		delete(board, selectedPiece.LastPosition)
		delete(board, 63)
		selectedPiece = nil
		whitesTurn = !whitesTurn
		e.calculateAllPositionsNew(whitesTurn)
		return true
	}

	return false
}
