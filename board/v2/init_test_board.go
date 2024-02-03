package v2

import v2 "chess/pieces/v2"

func (a *App) initTestWhiteBoard() { // todo remove: only for testing purpose
	a.whitesTurn = true
	a.whiteBoard = make(map[int]*v2.Piece)

	a.addTestPiece(48, v2.Pawn, true)
	a.addTestPiece(49, v2.Pawn, true)
	a.addTestPiece(42, v2.Pawn, true)
	a.addTestPiece(35, v2.Pawn, true)
	a.addTestPiece(52, v2.Pawn, true)
	a.addTestPiece(37, v2.Pawn, true)
	a.addTestPiece(30, v2.Pawn, true)
	a.addTestPiece(47, v2.Pawn, true)
	a.addTestPiece(32, v2.Queen, true)
	a.addTestPiece(58, v2.Bishop, true)
	a.addTestPiece(61, v2.Bishop, true)
	a.addTestPiece(56, v2.Rook, true)
	a.addTestPiece(63, v2.Rook, true)
	a.addTestPiece(57, v2.Knight, true)
	a.addTestPiece(28, v2.Knight, true)
	a.addTestPiece(53, v2.King, true)
}

func (a *App) initTestBlackBoard() { // todo remove: only for testing purpose
	a.blackBoard = make(map[int]*v2.Piece)

	//a.addTestPiece(10, v2.King, false)
	a.addTestPiece(8, v2.Pawn, false)
	a.addTestPiece(9, v2.Pawn, false)
	a.addTestPiece(18, v2.Pawn, false)
	a.addTestPiece(27, v2.Pawn, false)
	a.addTestPiece(12, v2.Pawn, false)
	a.addTestPiece(29, v2.Pawn, false)
	//a.addTestPiece(14, v2.Pawn, false)
	a.addTestPiece(39, v2.Pawn, false)
	a.addTestPiece(0, v2.Rook, false)
	a.addTestPiece(1, v2.Knight, false)
	a.addTestPiece(2, v2.Bishop, false)
	a.addTestPiece(3, v2.Queen, false)
	a.addTestPiece(4, v2.King, false)
	a.addTestPiece(5, v2.Bishop, false)
	a.addTestPiece(6, v2.Knight, false)
	a.addTestPiece(7, v2.Rook, false)
}

func (a *App) addTestPiece(pos int, kind v2.PieceKind, white bool) { // todo remove: only for testing purpose
	switch white {
	case true:
		a.whiteBoard[pos] = &v2.Piece{
			Kind:         kind,
			White:        white,
			Options:      make(map[int]struct{}),
			LastPosition: pos,
		}
		if kind == v2.King {
			a.whiteBoard[pos].CheckingPieces = make(map[int]*v2.Piece)
			a.whiteBoard[pos].HasBeenMoved = true
		}
		if kind == v2.Pawn {
			a.whiteBoard[pos].EnPassantOptions = make(map[int]int)
		}
	case false:
		a.blackBoard[pos] = &v2.Piece{
			Kind:         kind,
			White:        white,
			Options:      make(map[int]struct{}),
			LastPosition: pos,
		}
		if kind == v2.King {
			a.blackBoard[pos].CheckingPieces = make(map[int]*v2.Piece)
			a.blackBoard[pos].HasBeenMoved = true
		}
		if kind == v2.Pawn {
			a.blackBoard[pos].EnPassantOptions = make(map[int]int)
		}
	}
}
