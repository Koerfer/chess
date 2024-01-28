package enginev2

import (
	"chess/pieces"
)

type Engine struct {
	allBlackMoves map[int][]*pieces.Piece
	allWhiteMoves map[int][]*pieces.Piece

	blackBoard map[int]*pieces.Piece
	whiteBoard map[int]*pieces.Piece

	options map[float32][]*Option

	depth int
}

type Option struct {
	Piece     *pieces.Piece
	MoveTo    int
	EnPassant int
	value     float32
}

func (e *Engine) Init(depth int) {
	e.options = make(map[float32][]*Option)
	e.depth = depth
}

func (e *Engine) Start(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece, white bool) *Option {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard

	e.options = make(map[float32][]*Option)

	e.createAllMoves(white)
	e.AssignValues(white)

	bestValue := float32(-10000)
	var bestOption *Option

	if e.depth == 0 {
		for value := range e.options {
			if value > bestValue {
				bestValue = value
			}
		}
		for _, option := range e.options[bestValue] {
			return option
		}
	}

	for value, options := range e.options {
		for _, option := range options {
			newWhiteBoard, newBlackBoard := e.createNewBoardState(option)
			newEngine := &Engine{}
			newEngine.Init(e.depth - 1)
			newOption := newEngine.Start(newWhiteBoard, newBlackBoard, white)
			if value-newOption.value >= bestValue {
				bestValue = value - newOption.value
				bestOption = option
			}
		}
	}

	//var bestValue float32
	//
	//for value := range e.options {
	//	if value > bestValue {
	//		bestValue = value
	//	}
	//}
	//for _, option := range e.options[bestValue] {
	//	return option
	//}

	return bestOption
	//return nil
}

func (e *Engine) createAllMoves(white bool) {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	if white {
		for position, piece := range e.whiteBoard {
			piece.LastPosition = position
			for option := range piece.Options {
				if p, ok := e.allWhiteMoves[option]; ok {
					e.allWhiteMoves[option] = append(p, piece)
					continue
				}
				e.allWhiteMoves[option] = []*pieces.Piece{piece}
			}
			if piece.Kind == pieces.Pawn {
				for option := range piece.EnPassantOptions {
					if p, ok := e.allWhiteMoves[option]; ok {
						e.allWhiteMoves[option] = append(p, piece)
						continue
					}
					e.allWhiteMoves[option] = []*pieces.Piece{piece}
				}
			}
		}
		return
	}

	for position, piece := range e.blackBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := e.allBlackMoves[option]; ok {
				e.allBlackMoves[option] = append(p, piece)
				continue
			}
			e.allBlackMoves[option] = []*pieces.Piece{piece}
		}
		if piece.Kind == pieces.Pawn {
			for option := range piece.EnPassantOptions {
				if p, ok := e.allBlackMoves[option]; ok {
					e.allBlackMoves[option] = append(p, piece)
					continue
				}
				e.allBlackMoves[option] = []*pieces.Piece{piece}
			}
		}
	}
}

func (e *Engine) AssignValues(white bool) {
	var myMoves map[int][]*pieces.Piece
	var opponentBoard map[int]*pieces.Piece
	switch white {
	case true:
		myMoves = e.allWhiteMoves
		opponentBoard = e.blackBoard
	case false:
		myMoves = e.allBlackMoves
		opponentBoard = e.whiteBoard
	}
	for move, ps := range myMoves {
		for _, piece := range ps {
			var value float32
			if piece.Kind == pieces.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1
					e.options[value] = append(e.options[value], &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					})
					continue
				}
			}

			if capturePiece, ok := opponentBoard[move]; ok {
				value += float32(capturePiece.Kind)
			}
			if piece.HasBeenMoved == false {
				value += 0.3
			}
			e.options[value] = append(e.options[value], &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			})
		}
	}
}
