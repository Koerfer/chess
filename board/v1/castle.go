package v1

import (
	"chess/pieces/v1"
)

func (a *App) castle(option int, board map[int]*v1.Piece) bool {
	switch option {
	case 2:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[3] = board[0]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 0)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 6:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[5] = board[7]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 7)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 58:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[59] = board[56]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 56)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 62:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[61] = board[63]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 63)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}
