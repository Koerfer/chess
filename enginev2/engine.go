package enginev2

import (
	"chess/pieces"
	"fmt"
	"github.com/kingledion/go-tools/tree"
)

type Engine struct {
	allBlackMoves map[int][]*pieces.Piece
	allWhiteMoves map[int][]*pieces.Piece

	blackBoard map[int]*pieces.Piece
	whiteBoard map[int]*pieces.Piece

	options map[float32][]*Option

	depth      uint
	tree       *tree.Tree[*Option]
	lastNodeId uint

	bestValue float32
}

type Option struct {
	Piece     *pieces.Piece
	MoveTo    int
	EnPassant int
	value     float32
	aggValue  float32
}

func (e *Engine) Init(depth uint) {
	e.options = make(map[float32][]*Option)
	e.depth = depth
}

func (e *Engine) Start(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece, parentID uint, oldTree *tree.Tree[*Option], white bool) *Option {
	e.allWhiteMoves = make(map[int][]*pieces.Piece)
	e.allBlackMoves = make(map[int][]*pieces.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard

	e.options = make(map[float32][]*Option)
	if oldTree != nil {
		e.tree = oldTree
	} else {
		e.tree = tree.Empty[*Option]()
		e.tree.Add(0, 0, &Option{
			Piece: &pieces.Piece{
				White: true,
			},
			MoveTo:    0,
			EnPassant: 0,
			value:     0,
		})
	}

	switch white {
	case true:
		e.allWhiteMoves = createAllMoves(e.whiteBoard, e.blackBoard, white)
	case false:
		e.allBlackMoves = createAllMoves(e.whiteBoard, e.blackBoard, white)
	}

	parent, ok := e.tree.Find(parentID)
	if !ok {
		fmt.Println("Not found parent")
	}
	e.lastNodeId = e.AssignValues(parentID, e.lastNodeId, parent.GetData().value, white)

	e.buildTree(e.depth, e.tree.Root().GetChildren())

	e.bestValue = -10000
	e.generateAggValues(e.tree.Root().GetChildren(), e.tree.Root())
	var option *Option
	for _, child := range e.tree.Root().GetChildren() {
		if child.GetData().value+child.GetData().aggValue > e.bestValue {
			e.bestValue = child.GetData().value + child.GetData().aggValue
			option = child.GetData()
		}
	}
	fmt.Println(e.bestValue)

	return option
}

func (e *Engine) generateAggValues(children []tree.Node[*Option], parent tree.Node[*Option]) {
	for _, child := range children {
		if len(child.GetChildren()) != 0 {
			e.generateAggValues(child.GetChildren(), child)
			for _, child2 := range child.GetChildren() {
				option := child2.GetData()
				if option.Piece.White { // if white, black parent
					if child.GetData().aggValue > option.value+option.aggValue {
						child.GetData().aggValue = option.value + option.aggValue
					}
				} else { // if black, white parent
					if child.GetData().aggValue < option.value+option.aggValue {
						child.GetData().aggValue = option.value + option.aggValue
					}
				}
			}
		} else {
			option := child.GetData()
			if option.Piece.White { // if white leaf, black parent
				if parent.GetData().aggValue > option.value {
					parent.GetData().aggValue = option.value
				}
			} else { // if black leaf, white parent
				if parent.GetData().aggValue < option.value {
					parent.GetData().aggValue = option.value
				}
			}
		}
	}
}

func (e *Engine) buildTree(depth uint, children []tree.Node[*Option]) {
	if depth == 0 {
		return
	}
	for _, child := range children {
		var newWhiteBoard map[int]*pieces.Piece
		var newBlackBoard map[int]*pieces.Piece

		white := !child.GetData().Piece.White
		if white {
			newWhiteBoard, newBlackBoard = e.createNewBoardState(child.GetData())
		} else {
			newBlackBoard, newWhiteBoard = e.createNewBoardState(child.GetData())
		}

		moves := createAllMoves(newWhiteBoard, newBlackBoard, white)

		var opponentBoard map[int]*pieces.Piece
		switch white {
		case true:
			opponentBoard = newBlackBoard
		case false:
			opponentBoard = newWhiteBoard
		}

		e.lastNodeId = assignValues(moves, opponentBoard, e.tree, child.GetID(), e.lastNodeId, white)
	}

	for _, child := range children {
		e.buildTree(depth-1, child.GetChildren())
	}
}

func createAllMoves(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece, white bool) map[int][]*pieces.Piece {
	allWhiteMoves := make(map[int][]*pieces.Piece)
	allBlackMoves := make(map[int][]*pieces.Piece)
	if white {
		for position, piece := range whiteBoard {
			piece.LastPosition = position
			for option := range piece.Options {
				if p, ok := allWhiteMoves[option]; ok {
					allWhiteMoves[option] = append(p, piece)
					continue
				}
				allWhiteMoves[option] = []*pieces.Piece{piece}
			}
			if piece.Kind == pieces.Pawn {
				for option := range piece.EnPassantOptions {
					if p, ok := allWhiteMoves[option]; ok {
						allWhiteMoves[option] = append(p, piece)
						continue
					}
					allWhiteMoves[option] = []*pieces.Piece{piece}
				}
			}
		}
		return allWhiteMoves
	}

	for position, piece := range blackBoard {
		piece.LastPosition = position
		for option := range piece.Options {
			if p, ok := allBlackMoves[option]; ok {
				allBlackMoves[option] = append(p, piece)
				continue
			}
			allBlackMoves[option] = []*pieces.Piece{piece}
		}
		if piece.Kind == pieces.Pawn {
			for option := range piece.EnPassantOptions {
				if p, ok := allBlackMoves[option]; ok {
					allBlackMoves[option] = append(p, piece)
					continue
				}
				allBlackMoves[option] = []*pieces.Piece{piece}
			}
		}
	}

	return allBlackMoves
}

func assignValues(myMoves map[int][]*pieces.Piece, opponentBoard map[int]*pieces.Piece, tree *tree.Tree[*Option], parent uint, nodeId uint, white bool) uint {
	var valueMultiplier float32

	switch white {
	case true:
		valueMultiplier = -1
	case false:
		valueMultiplier = 1
	}
	for move, ps := range myMoves {
		for _, piece := range ps {
			nodeId++
			var value float32
			if piece.Kind == pieces.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1 * valueMultiplier
					option := &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					}
					tree.Add(nodeId, parent, option)
					continue
				}
			}

			if capturePiece, ok := opponentBoard[move]; ok {
				value += float32(capturePiece.Kind) * valueMultiplier
			}
			if piece.HasBeenMoved == false {
				value += 0.3 * valueMultiplier
			}
			option := &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			}
			tree.Add(nodeId, parent, option)
		}
	}

	return nodeId
}

func (e *Engine) AssignValues(parent uint, nodeId uint, parentValue float32, white bool) uint {
	var myMoves map[int][]*pieces.Piece
	var opponentBoard map[int]*pieces.Piece
	var valueMultiplier float32

	switch white {
	case true:
		myMoves = e.allWhiteMoves
		opponentBoard = e.blackBoard
		valueMultiplier = -1
	case false:
		myMoves = e.allBlackMoves
		opponentBoard = e.whiteBoard
		valueMultiplier = 1
	}
	for move, ps := range myMoves {
		for _, piece := range ps {
			nodeId++
			value := parentValue
			if piece.Kind == pieces.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1 * valueMultiplier
					option := &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					}
					e.options[value] = append(e.options[value], option)
					e.tree.Add(nodeId, parent, option)
					continue
				}
			}

			if capturePiece, ok := opponentBoard[move]; ok {
				value += float32(capturePiece.Kind) * valueMultiplier
			}
			if piece.HasBeenMoved == false {
				value += 0.3 * valueMultiplier
			}
			option := &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			}
			e.options[value] = append(e.options[value], option)
			e.tree.Add(nodeId, parent, option)
		}
	}

	return nodeId
}
