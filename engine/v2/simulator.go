package v2

import (
	v2 "chess/pieces/v2"
)

type Simulator struct {
	whiteBoard map[int]v2.PieceInterface
	blackBoard map[int]v2.PieceInterface

	whiteTurn bool

	newWhiteEval float64
	newBlackEval float64
}

func (s *Simulator) start(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface, moveToMake *SelectedMove) {
	s.copyBoard(whiteBoard, blackBoard)

	copyOfMove := &SelectedMove{
		Piece:       moveToMake.Piece.Copy(false),
		ToPosition:  moveToMake.ToPosition,
		valueOfMove: moveToMake.valueOfMove,
	}

	s.whiteTurn = moveToMake.Piece.GetWhite()
	var board map[int]v2.PieceInterface
	switch s.whiteTurn {
	case true:
		board = s.whiteBoard
	case false:
		board = s.blackBoard
	}
	s.performMove(copyOfMove, board)

	newEngine := &Engine{}
	newEngine.Init(s.whiteBoard, s.blackBoard, !s.whiteTurn, false)

	s.newWhiteEval = newEngine.whiteEval
	s.newBlackEval = newEngine.blackEval
}
