package v2

import v2 "chess/pieces/v2"

func (a *App) castle(king *v2.King, option int, board map[int]v2.PieceInterface) bool {
	switch option {
	case 2:
		king.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[3] = board[0]
		delete(board, king.LastPosition)
		delete(board, 0)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 6:
		king.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[5] = board[7]
		delete(board, king.LastPosition)
		delete(board, 7)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 58:
		king.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[59] = board[56]
		delete(board, king.LastPosition)
		delete(board, 56)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 62:
		king.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[61] = board[63]
		delete(board, king.LastPosition)
		delete(board, 63)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}
