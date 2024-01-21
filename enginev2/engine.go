package enginev2

import (
	"chess/pieces"
)

type Engine struct {
	allBlackMoves map[int][]*pieces.Piece
	allWhiteMoves map[int][]*pieces.Piece

	blackBoard map[int]*pieces.Piece
	whiteBoard map[int]*pieces.Piece

	options map[int][]*Option
}

type Option struct {
	Piece     *pieces.Piece
	MoveTo    int
	EnPassant int
	value     int
}

func (e *Engine) Init() {
	e.options = make(map[int][]*Option)
}

func (e *Engine) Start(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece) *Option {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard
	e.options = make(map[int][]*Option)

	e.createAllMoves()
	e.AssignValues()

	var bestValue int
	for value := range e.options {
		if value > bestValue {
			bestValue = value
		}
	}
	for _, option := range e.options[bestValue] {
		return option
	}

	return nil
}

func (e *Engine) createAllMoves() {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	for position, piece := range e.blackBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := e.allBlackMoves[option]; ok {
				p = append(p, piece)
				continue
			}
			e.allBlackMoves[option] = []*pieces.Piece{piece}
		}
		if piece.Kind == pieces.Pawn {
			for option := range piece.EnPassantOptions {
				if p, ok := e.allBlackMoves[option]; ok {
					p = append(p, piece)
					continue
				}
				e.allBlackMoves[option] = []*pieces.Piece{piece}
			}
		}
	}
	for position, piece := range e.whiteBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := e.allWhiteMoves[option]; ok {
				p = append(p, piece)
				continue
			}
			e.allWhiteMoves[option] = []*pieces.Piece{piece}
		}
		if piece.Kind == pieces.Pawn {
			for option := range piece.EnPassantOptions {
				if p, ok := e.allWhiteMoves[option]; ok {
					p = append(p, piece)
					continue
				}
				e.allWhiteMoves[option] = []*pieces.Piece{piece}
			}
		}
	}
}

func (e *Engine) AssignValues() {
	for move, ps := range e.allBlackMoves {
		for _, piece := range ps {
			var value int
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

			if capturePiece, ok := e.whiteBoard[move]; ok {
				value += int(capturePiece.Kind)
			}
			e.options[value] = append(e.options[value], &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			})
		}
	}
}
