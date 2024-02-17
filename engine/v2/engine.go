package v2

import (
	v2 "chess/pieces/v2"
	"fmt"
	"github.com/kingledion/go-tools/tree"
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
		myTree := *tree.Empty[*SelectedMove]()
		myTree.Add(0, 0, &SelectedMove{})
		chosen10 := e.getInitial10()

		var nodeId uint
		for _, chosen := range chosen10 {
			if chosen == nil {
				continue
			}
			nodeId++
			myTree.Add(nodeId, 0, chosen)
		}

		for _, child := range myTree.Root().GetChildren() {
			next10 := e.getNext10(child.GetData())
			for _, chosen := range next10 {
				if chosen == nil {
					continue
				}
				nodeId++
				myTree.Add(nodeId, child.GetID(), chosen)
			}
			for _, grandChild := range child.GetChildren() {
				next10 := e.getNext10(grandChild.GetData())
				for _, chosen := range next10 {
					if chosen == nil {
						continue
					}
					nodeId++
					myTree.Add(nodeId, grandChild.GetID(), chosen)
				}
			}
		}

		bestChildValue := 100000.00
		for _, child := range myTree.Root().GetChildren() {
			if child == nil || child.GetData() == nil {
				continue
			}
			fmt.Printf("childValue %v, piece: %v, to: %v\n", child.GetData().valueOfMove, child.GetData().Piece.GetValue(), child.GetData().ToPosition)
			bestGrandChildValue := -100000.00

			for _, grandChild := range child.GetChildren() {
				if grandChild == nil || grandChild.GetData() == nil {
					continue
				}
				bestGrand2ChildValue := 100000.00
				for _, grand2Child := range grandChild.GetChildren() {
					if grand2Child == nil || grand2Child.GetData() == nil {
						continue
					}

					if grand2Child.GetData().valueOfMove < bestGrand2ChildValue {
						bestGrand2ChildValue = grand2Child.GetData().valueOfMove
					}
				}

				fmt.Printf("grandChildValue %v, piece: %v, to: %v\n", grandChild.GetData().valueOfMove, grandChild.GetData().Piece.GetValue(), grandChild.GetData().ToPosition)
				if bestGrand2ChildValue+grandChild.GetData().valueOfMove > bestGrandChildValue {
					bestGrandChildValue = bestGrand2ChildValue + grandChild.GetData().valueOfMove
				}
			}

			fmt.Printf("bestGrandChildValue: %v\n\n", bestGrandChildValue)
			if bestGrandChildValue+child.GetData().valueOfMove < bestChildValue {
				bestChildValue = bestGrandChildValue + child.GetData().valueOfMove
				bestChild = child.GetData()
			}
		}
		fmt.Printf("bestChildValue %v\n", bestChildValue)
		fmt.Println("-----------------------------------------------------")
		fmt.Println()
	}

	return bestChild
}

func (e *Engine) getInitial10() []*SelectedMove {
	chosen10 := make([]*SelectedMove, 10)
	valuePerIndex := make([]float64, 10)
	counter := 0

	for _, piece := range e.blackBoard {
		for option := range piece.GetOptions() {
			move := &SelectedMove{
				Piece:       piece,
				ToPosition:  option,
				valueOfMove: 0,
			}
			simulator := &Simulator{}
			simulator.start(e.whiteBoard, e.blackBoard, move)
			newValue := simulator.newWhiteEval - simulator.newBlackEval
			move.valueOfMove = newValue
			move.simulator = simulator

			index, worstValue := maxOfSlice(valuePerIndex)
			if counter < 10 {
				chosen10[counter] = move
				counter++
			} else if worstValue > newValue {
				chosen10[index] = move
				valuePerIndex[index] = newValue
			}
		}
	}

	return chosen10
}

func (e *Engine) getNext10(selectedMove *SelectedMove) []*SelectedMove {
	chosen10 := make([]*SelectedMove, 10)
	valuePerIndex := make([]float64, 10)
	counter := 0
	white := !selectedMove.Piece.GetWhite()

	whiteBoard := selectedMove.simulator.whiteBoard
	blackBoard := selectedMove.simulator.blackBoard
	var myBoard map[int]v2.PieceInterface
	switch white {
	case true:
		myBoard = whiteBoard
	case false:
		myBoard = blackBoard
	}

	for _, piece := range myBoard {
		for option := range piece.GetOptions() {
			move := &SelectedMove{
				Piece:       piece,
				ToPosition:  option,
				valueOfMove: 0,
			}
			simulator := &Simulator{}
			simulator.start(whiteBoard, blackBoard, move)
			newValue := simulator.newWhiteEval - simulator.newBlackEval
			move.valueOfMove = newValue
			move.simulator = simulator

			index, worstValue := 0, 0.00
			switch piece.GetWhite() {
			case true:
				index, worstValue = minOfSlice(valuePerIndex)
			case false:
				index, worstValue = maxOfSlice(valuePerIndex)
			}

			if counter < 10 {
				chosen10[counter] = move
				valuePerIndex[index] = newValue
				counter++
			} else {
				switch piece.GetWhite() {
				case true:
					if worstValue < newValue {
						chosen10[index] = move
						valuePerIndex[index] = newValue
					}
				case false:
					if worstValue > newValue {
						chosen10[index] = move
						valuePerIndex[index] = newValue
					}
				}
			}
		}
	}

	return chosen10
}

func maxOfSlice(valuePerIndex []float64) (int, float64) {
	var worstIndex int
	worstValue := -1000.00
	for i, value := range valuePerIndex {
		if value > worstValue {
			worstValue = value
			worstIndex = i
		}
	}

	return worstIndex, worstValue
}

func minOfSlice(valuePerIndex []float64) (int, float64) {
	var worstIndex int
	worstValue := 1000.00
	for i, value := range valuePerIndex {
		if value < worstValue {
			worstValue = value
			worstIndex = i
		}
	}

	return worstIndex, worstValue
}
