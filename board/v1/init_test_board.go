package v1

import v1 "chess/pieces/v1"

func (a *App) initTestWhiteBoard() { // todo remove: only for testing purpose
	a.whitesTurn = true
	a.whiteBoard = make(map[int]*v1.Piece)

	a.addTestPiece(48, v1.Pawn, true)
	a.addTestPiece(49, v1.Pawn, true)
	a.addTestPiece(42, v1.Pawn, true)
	a.addTestPiece(35, v1.Pawn, true)
	a.addTestPiece(52, v1.Pawn, true)
	a.addTestPiece(37, v1.Pawn, true)
	a.addTestPiece(30, v1.Pawn, true)
	a.addTestPiece(47, v1.Pawn, true)
	a.addTestPiece(32, v1.Queen, true)
	a.addTestPiece(58, v1.Bishop, true)
	a.addTestPiece(61, v1.Bishop, true)
	a.addTestPiece(56, v1.Rook, true)
	a.addTestPiece(63, v1.Rook, true)
	a.addTestPiece(57, v1.Knight, true)
	a.addTestPiece(28, v1.Knight, true)
	a.addTestPiece(53, v1.King, true)
}

func (a *App) initTestBlackBoard() { // todo remove: only for testing purpose
	a.blackBoard = make(map[int]*v1.Piece)

	//a.addTestPiece(10, v1.King, false)
	a.addTestPiece(8, v1.Pawn, false)
	a.addTestPiece(9, v1.Pawn, false)
	a.addTestPiece(18, v1.Pawn, false)
	a.addTestPiece(27, v1.Pawn, false)
	a.addTestPiece(12, v1.Pawn, false)
	a.addTestPiece(29, v1.Pawn, false)
	//a.addTestPiece(14, v1.Pawn, false)
	a.addTestPiece(39, v1.Pawn, false)
	a.addTestPiece(0, v1.Rook, false)
	a.addTestPiece(1, v1.Knight, false)
	a.addTestPiece(2, v1.Bishop, false)
	a.addTestPiece(3, v1.Queen, false)
	a.addTestPiece(4, v1.King, false)
	a.addTestPiece(5, v1.Bishop, false)
	a.addTestPiece(6, v1.Knight, false)
	a.addTestPiece(7, v1.Rook, false)
}

func (a *App) addTestPiece(pos int, kind v1.PieceKind, white bool) { // todo remove: only for testing purpose
	switch white {
	case true:
		a.whiteBoard[pos] = &v1.Piece{
			Kind:         kind,
			White:        white,
			Options:      make(map[int]struct{}),
			LastPosition: pos,
		}
		if kind == v1.King {
			a.whiteBoard[pos].CheckingPieces = make(map[int]*v1.Piece)
			a.whiteBoard[pos].HasBeenMoved = true
		}
		if kind == v1.Pawn {
			a.whiteBoard[pos].EnPassantOptions = make(map[int]int)
		}
	case false:
		a.blackBoard[pos] = &v1.Piece{
			Kind:         kind,
			White:        white,
			Options:      make(map[int]struct{}),
			LastPosition: pos,
		}
		if kind == v1.King {
			a.blackBoard[pos].CheckingPieces = make(map[int]*v1.Piece)
			a.blackBoard[pos].HasBeenMoved = true
		}
		if kind == v1.Pawn {
			a.blackBoard[pos].EnPassantOptions = make(map[int]int)
		}
	}
}
