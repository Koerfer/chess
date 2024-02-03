package v2

import (
	"fmt"
	"github.com/kingledion/go-tools/tree"
)

type Engine struct {
	original bool

	allBlackMoves map[int][]*v2.Piece
	allWhiteMoves map[int][]*v2.Piece

	blackBoard map[int]*v2.Piece
	whiteBoard map[int]*v2.Piece

	depth      int
	tree       *tree.Tree[*Option]
	lastNodeId uint

	bestValue     float32
	nodesPerDepth map[int][]uint

	discardedChildren map[uint]tree.Node[*Option]
}

type Option struct {
	Piece       *v2.Piece
	MoveTo      int
	EnPassant   int
	value       float32
	aggValue    float32
	chosenChild uint
}

func (e *Engine) Init(depth int, original bool) {
	e.depth = depth
	e.nodesPerDepth = make(map[int][]uint)
	e.original = original
	e.discardedChildren = make(map[uint]tree.Node[*Option])
}

func (e *Engine) Start(whiteBoard map[int]*v2.Piece, blackBoard map[int]*v2.Piece, parentID uint, white bool) tree.Node[*Option] {
	e.allWhiteMoves = make(map[int][]*v2.Piece)
	e.allBlackMoves = make(map[int][]*v2.Piece)
	e.whiteBoard = whiteBoard
	e.blackBoard = blackBoard
	e.discardedChildren = make(map[uint]tree.Node[*Option])

	e.tree = tree.Empty[*Option]()
	e.tree.Add(0, 0, &Option{
		Piece: &v2.Piece{
			White: true,
		},
		MoveTo:    0,
		EnPassant: 0,
		value:     0,
	})
	e.nodesPerDepth = make(map[int][]uint)

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

	e.generateValues()
	var node tree.Node[*Option]

	//if !e.original {
	//	for k, v := range e.nodesPerDepth {
	//		fmt.Printf("\ndepth: %v, size: %v", k, len(v))
	//	}
	//	for _, child := range e.tree.Root().GetChildren() {
	//		fmt.Printf("\npiece: %v, moveto: %v, value: %v, white: %v", child.GetData().Piece.Kind, child.GetData().MoveTo, child.GetData().value, child.GetData().Piece.White)
	//	}
	//} else {
	//	for k, v := range e.nodesPerDepth {
	//		fmt.Printf("\noriginal depth: %v, size: %v", k, len(v))
	//	}
	//}

	if e.original {
		if len(e.tree.Root().GetChildren()) == 0 {
			return nil
		}
		for {
			e.bestValue = -10000
			for _, child := range e.tree.Root().GetChildren() {
				if _, ok := e.discardedChildren[child.GetID()]; ok {
					continue
				}

				if child.GetData().value+child.GetData().aggValue > e.bestValue {
					e.bestValue = child.GetData().value + child.GetData().aggValue
					node = child
				}
			}

			fmt.Println(len(e.discardedChildren))
			if len(e.discardedChildren) == len(e.tree.Root().GetChildren()) {
				e.bestValue = -10000
				for _, v := range e.discardedChildren {
					if v.GetData().value+v.GetData().aggValue > e.bestValue {
						e.bestValue = v.GetData().value + v.GetData().aggValue
						node = v
					}
				}

				fmt.Println(e.bestValue)
				return node
			}

			if e.checkIfChosenNodeIsGood(node) {
				break
			}
		}

		fmt.Println(e.bestValue)

		return node
	} else {
		e.bestValue = -10000
		for _, child := range e.tree.Root().GetChildren() {
			if child.GetData().value+child.GetData().aggValue > e.bestValue {
				e.bestValue = child.GetData().value + child.GetData().aggValue
				node = child
			}
		}

		return node
	}
}

func (e *Engine) checkIfChosenNodeIsGood(node tree.Node[*Option]) bool {
	option := node.GetData()
	chosenNode, _ := e.tree.Find(option.chosenChild)
	//fmt.Printf("\nchosen piece: %v, moveto: %v, value: %v, white: %v", chosenNode.GetData().Piece.Kind, chosenNode.GetData().MoveTo, chosenNode.GetData().value, chosenNode.GetData().Piece.White)
	white := chosenNode.GetData().Piece.White
	if white {
		newWhiteBoard, newBlackBoard := e.createNewBoardState(chosenNode)
		newEngine := Engine{}
		newEngine.Init(e.depth, false)
		newNode := newEngine.Start(newWhiteBoard, newBlackBoard, 0, white)
		newOption := newNode.GetData()
		if option.value+option.aggValue < newOption.aggValue+newOption.value {
			node.GetData().aggValue = newNode.GetData().aggValue
			node.GetData().value = newNode.GetData().value
			e.discardedChildren[node.GetID()] = node
			return false
		}
	} else {
		newBlackBoard, newWhiteBoard := e.createNewBoardState(chosenNode)
		newEngine := Engine{}
		newEngine.Init(e.depth, false)
		newNode := newEngine.Start(newWhiteBoard, newBlackBoard, 0, white)
		if newNode == nil {
			return false
		}
		newOption := newNode.GetData()
		if option.value+option.aggValue < newOption.aggValue+newOption.value {
			node.GetData().aggValue = newNode.GetData().aggValue
			node.GetData().value = newNode.GetData().value
			e.discardedChildren[node.GetID()] = node
			return false
		}
	}

	return true
}

func (e *Engine) generateValues() {
	for i := e.depth - 1; i >= 0; i-- {
		for _, nodeId := range e.nodesPerDepth[i] {
			node, _ := e.tree.Find(nodeId)
			children := node.GetChildren()
			if len(children) == 0 {
				value := float32(100)
				if node.GetData().Piece.White {
					value = -100
				}
				node.GetData().value += value
				continue
			}
			e.calculateAggFromChildren(children, node)
		}
	}
}

func (e *Engine) calculateAggFromChildren(children []tree.Node[*Option], parent tree.Node[*Option]) {
	parentOption := parent.GetData()
	if parentOption.Piece.White {
		parentOption.aggValue = -10000
	} else {
		parentOption.aggValue = 10000
	}
	for _, child := range children {
		option := child.GetData()
		if option.Piece.White { // if white leaf, black parent
			if parentOption.aggValue > option.value+option.aggValue {
				parentOption.aggValue = option.value + option.aggValue
				if option.chosenChild != 0 {
					parentOption.chosenChild = option.chosenChild
					continue
				}
				parentOption.chosenChild = child.GetID()
			}
		} else { // if black leaf, white parent
			if parentOption.aggValue < option.value+option.aggValue {
				parentOption.aggValue = option.value + option.aggValue
				if option.chosenChild != 0 {
					parentOption.chosenChild = option.chosenChild
					continue
				}
				parentOption.chosenChild = child.GetID()
			}
		}
	}
}

func (e *Engine) buildTree(depth int, children []tree.Node[*Option]) {
	if depth == 0 {
		return
	}
	for _, child := range children {
		var newWhiteBoard map[int]*v2.Piece
		var newBlackBoard map[int]*v2.Piece

		white := child.GetData().Piece.White
		if white {
			newWhiteBoard, newBlackBoard = e.createNewBoardState(child)
		} else {
			newBlackBoard, newWhiteBoard = e.createNewBoardState(child)
		}

		moves := createAllMoves(newWhiteBoard, newBlackBoard, white)

		var opponentBoard map[int]*v2.Piece
		switch white {
		case true:
			opponentBoard = newBlackBoard
		case false:
			opponentBoard = newWhiteBoard
		}

		e.lastNodeId = e.assignValues(moves, opponentBoard, child.GetID(), e.lastNodeId, white, depth)
	}

	for _, child := range children {
		e.buildTree(depth-1, child.GetChildren())
	}
}

func createAllMoves(whiteBoard map[int]*v2.Piece, blackBoard map[int]*v2.Piece, white bool) map[int][]*v2.Piece {
	allWhiteMoves := make(map[int][]*v2.Piece)
	allBlackMoves := make(map[int][]*v2.Piece)
	if white {
		for position, piece := range whiteBoard {
			piece.LastPosition = position
			for option := range piece.Options {
				if p, ok := allWhiteMoves[option]; ok {
					allWhiteMoves[option] = append(p, piece)
					continue
				}
				allWhiteMoves[option] = []*v2.Piece{piece}
			}
			if piece.Kind == v2.Pawn {
				for option := range piece.EnPassantOptions {
					if p, ok := allWhiteMoves[option]; ok {
						allWhiteMoves[option] = append(p, piece)
						continue
					}
					allWhiteMoves[option] = []*v2.Piece{piece}
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
			allBlackMoves[option] = []*v2.Piece{piece}
		}
		if piece.Kind == v2.Pawn {
			for option := range piece.EnPassantOptions {
				if p, ok := allBlackMoves[option]; ok {
					allBlackMoves[option] = append(p, piece)
					continue
				}
				allBlackMoves[option] = []*v2.Piece{piece}
			}
		}
	}

	return allBlackMoves
}

func (e *Engine) assignValues(myMoves map[int][]*v2.Piece, opponentBoard map[int]*v2.Piece, parent uint, nodeId uint, white bool, depth int) uint {
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
			if piece.Kind == v2.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1 * valueMultiplier
					option := &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					}
					e.tree.Add(nodeId, parent, option)
					e.nodesPerDepth[e.depth-depth+1] = append(e.nodesPerDepth[e.depth-depth+1], nodeId)
					continue
				}
			}

			if capturePiece, ok := opponentBoard[move]; ok {
				value += float32(capturePiece.Kind) * valueMultiplier
			}
			option := &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			}
			e.tree.Add(nodeId, parent, option)
			e.nodesPerDepth[e.depth-depth+1] = append(e.nodesPerDepth[e.depth-depth+1], nodeId)
		}
	}

	return nodeId
}

func (e *Engine) AssignValues(parent uint, nodeId uint, parentValue float32, white bool) uint {
	var myMoves map[int][]*v2.Piece
	var opponentBoard map[int]*v2.Piece
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
			if piece.Kind == v2.Pawn {
				if take, ok := piece.EnPassantOptions[move]; ok {
					value += 1 * valueMultiplier
					option := &Option{
						Piece:     piece,
						MoveTo:    move,
						EnPassant: take,
						value:     value,
					}
					e.tree.Add(nodeId, parent, option)
					e.nodesPerDepth[0] = append(e.nodesPerDepth[0], nodeId)
					continue
				}
			}

			if capturePiece, ok := opponentBoard[move]; ok {
				value += float32(capturePiece.Kind) * valueMultiplier
			}
			option := &Option{
				Piece:  piece,
				MoveTo: move,
				value:  value,
			}
			e.tree.Add(nodeId, parent, option)
			e.nodesPerDepth[0] = append(e.nodesPerDepth[0], nodeId)
		}
	}

	return nodeId
}
