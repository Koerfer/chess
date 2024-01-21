package engine

import (
	"chess/pieces"
	"github.com/kingledion/go-tools/tree"
)

type Engine struct {
	allBlackMoves map[int][]*pieces.Piece
	allWhiteMoves map[int][]*pieces.Piece

	blackBoard map[int]*pieces.Piece
	whiteBoard map[int]*pieces.Piece

	tree *tree.Tree[*Option]
}

type Option struct {
	Piece  *pieces.Piece
	MoveTo int
	value  int
}

func (e *Engine) Init() {
	e.tree = tree.Empty[*Option]()
	e.tree.Add(0, 0, &Option{
		Piece:  nil,
		MoveTo: 0,
		value:  0,
	})
}

func (e *Engine) Start(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece) *Option {
	e.tree = tree.Empty[*Option]()
	e.tree.Add(0, 0, &Option{
		Piece:  nil,
		MoveTo: 0,
		value:  0,
	})
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard

	for position, piece := range blackBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := e.allBlackMoves[option]; ok {
				p = append(p, piece)
				continue
			}
			e.allBlackMoves[option] = []*pieces.Piece{piece}
		}
	}
	for position, piece := range whiteBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := e.allWhiteMoves[option]; ok {
				p = append(p, piece)
				continue
			}
			e.allWhiteMoves[option] = []*pieces.Piece{piece}
		}

	}

	node, _ := e.tree.Find(e.AssignValues(e.tree.Root().GetID()))

	return node.GetData()
}

func (e *Engine) AssignValues(parent uint) uint {
	var bestValue int
	var bestNode uint
	var nodeId uint
	for move, ps := range e.allBlackMoves {
		for _, piece := range ps {
			nodeId++
			var value int
			if capturePiece, ok := e.whiteBoard[move]; ok {
				value += int(capturePiece.Kind)
			}
			e.tree.Add(nodeId, parent, &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			})
			if value >= bestValue {
				bestValue = value
				bestNode = nodeId
			}
		}
	}
	return bestNode
}
