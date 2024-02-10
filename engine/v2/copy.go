package v2

import (
	v2 "chess/pieces/v2"
)

func (e *Engine) copyBoard(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface) {
	e.whiteBoard = make(map[int]v2.PieceInterface)
	e.blackBoard = make(map[int]v2.PieceInterface)
	for k, v := range whiteBoard {
		e.whiteBoard[k] = v.Copy(true)
	}

	for k, v := range blackBoard {
		e.blackBoard[k] = v.Copy(true)
	}
}
