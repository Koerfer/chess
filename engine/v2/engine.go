package v2

import (
	v2 "chess/pieces/v2"
)

type Engine struct {
	whiteBoard map[int]v2.PieceInterface
	blackBoard map[int]v2.PieceInterface

	weakWhiteSquares      map[int]struct{}
	superWeakWhiteSquares map[int]struct{}
	weakBlackSquares      map[int]struct{}
	superWeakBlackSquares map[int]struct{}

	allWhiteOptions map[int]struct{}
	allBlackOptions map[int]struct{}

	selectedMove *SelectedMove

	whiteEval float64
	blackEval float64

	optionsSaved int
	nodeId       uint
}

type SelectedMove struct {
	Piece       v2.PieceInterface
	ToPosition  int
	valueOfMove float64
	simulator   *Simulator
}

func (e *Engine) Init(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface, white bool, goDeeper bool) *SelectedMove {
	e.whiteEval = 0
	e.blackEval = 0
	e.optionsSaved = 10
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard
	e.weakWhiteSquares = make(map[int]struct{})
	e.superWeakWhiteSquares = make(map[int]struct{})
	e.weakBlackSquares = make(map[int]struct{})
	e.superWeakBlackSquares = make(map[int]struct{})
	e.allWhiteOptions = make(map[int]struct{})
	e.allBlackOptions = make(map[int]struct{})

	e.getAllOptions()
	e.evaluateKingPosition()
	e.evaluatePawns()
	e.evaluatePieces()

	var bestChild *SelectedMove
	if goDeeper {
		bestChild = e.getBestMove()
	}

	return bestChild
}
