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

	newBlackBoard map[int]*pieces.Piece
	newWhiteBoard map[int]*pieces.Piece

	tree *tree.Tree[*Option]
}

type Option struct {
	Piece     *pieces.Piece
	MoveTo    int
	EnPassant int
	value     int
}

func (e *Engine) Init() {
	e.tree = tree.Empty[*Option]()
	e.tree.Add(0, 0, &Option{
		Piece:     nil,
		MoveTo:    0,
		EnPassant: 0,
		value:     0,
	})
}

func (e *Engine) Start(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece) *Option {
	e.tree = tree.Empty[*Option]()
	e.tree.Add(0, 0, &Option{
		Piece:     nil,
		MoveTo:    0,
		EnPassant: 0,
		value:     0,
	})
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard

	e.createAllMoves()
	e.tree.Find(e.AssignValues(e.tree.Root().GetID()))
	start := uint(len(e.tree.Root().GetChildren()))
	var bestFound uint
	for _, child := range e.tree.Root().GetChildren() {
		e.newWhiteBoard = e.whiteBoard
		e.newBlackBoard = e.blackBoard
		childData := child.GetData()
		selectedPiece := childData.Piece
		takeOrPromote(childData.MoveTo, selectedPiece)
		if childData.EnPassant != 0 {
			delete(e.newWhiteBoard, childData.EnPassant)
		}

		if selectedPiece.Kind == pieces.King && !selectedPiece.HasBeenMoved {
			e.castle(childData.MoveTo, e.newBlackBoard, selectedPiece, false)
			selectedPiece.HasBeenMoved = true
		} else {
			if selectedPiece.Kind == pieces.King || selectedPiece.Kind == pieces.Rook {
				selectedPiece.HasBeenMoved = true
			}

			e.newWhiteBoard[childData.MoveTo] = selectedPiece
			delete(e.newWhiteBoard, selectedPiece.LastPosition)
			selectedPiece = nil
		}

		e.calculateAllPositionsNew(false)
		e.createAllMovesNew()
		bestNode, next := e.AssignSecondValues(child.GetID(), start, child.GetData().value)
		start = next
		if bestNode >= bestNode {
			bestFound = bestNode
		}
	}
	node, _ := e.tree.Find(bestFound)
	bestMove := node.GetParent().GetData()
	return bestMove
}

func takeOrPromote(position int, selectedPiece *pieces.Piece) {
	if selectedPiece.Kind == pieces.Pawn {
		end := 7
		if selectedPiece.White == true {
			end = 0
		}
		if position/8 == end {
			selectedPiece.Kind = pieces.Queen // todo: add convert to better Piece logic
		}
	}

	//switch a.whitesTurn {
	//case true:
	//	delete(a.blackBoard, position)
	//case false:
	//	delete(a.whiteBoard, position)
	//}
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

func (e *Engine) createAllMovesNew() {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	for position, piece := range e.newBlackBoard {
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
	for position, piece := range e.newWhiteBoard {
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

func (e *Engine) AssignValues(parent uint) uint {
	var bestValue int
	var bestNode uint
	var nodeId uint
	for move, ps := range e.allBlackMoves {
		for _, piece := range ps {
			nodeId++
			var value int
			if piece.Kind == pieces.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1
					e.tree.Add(nodeId, parent, &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					})
					if value >= bestValue {
						bestValue = value
						bestNode = nodeId
					}
					continue
				}
			}

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

func (e *Engine) AssignSecondValues(parent uint, nodeId uint, parentValue int) (uint, uint) {
	var bestValue int
	var bestNode uint
	for move, ps := range e.allWhiteMoves {
		for _, piece := range ps {
			nodeId++
			value := parentValue
			if piece.Kind == pieces.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value -= 1
					e.tree.Add(nodeId, parent, &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					})
					if value >= bestValue {
						bestValue = value
						bestNode = nodeId
					}
					continue
				}
			}

			if capturePiece, ok := e.blackBoard[move]; ok {
				value -= int(capturePiece.Kind)
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
	return bestNode, nodeId
}
