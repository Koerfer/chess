package v2

import (
	v2 "chess/pieces/v2"
	"sync"
)

var wg sync.WaitGroup

func (e *Engine) getBestMove() *SelectedMove {
	input := make(chan *SelectedMove)
	numberOfWorkers := 30

	chosen10 := e.getInitial10()
	results := make(chan *SelectedMove, len(chosen10))
	bestValue := 100000.00
	bestOption := chosen10[0]

	for i := 0; i < numberOfWorkers; i++ {
		go e.bestNextMove(input, 2, results)
	}

	for _, chosen := range chosen10 {
		if chosen == nil {
			continue
		}
		wg.Add(1)
		input <- chosen
	}

	wg.Wait()
	close(input)

	for result := range results {
		if result.valueOfMove < bestValue {
			bestValue = result.valueOfMove
			bestOption = result
		}

		if len(results) == 0 {
			break
		}
	}
	close(results)

	return bestOption
}

func (e *Engine) bestNextMove(input chan *SelectedMove, depth int, result chan *SelectedMove) {
	for previousChosen := range input {
		white := !previousChosen.Piece.GetWhite()
		var nextBestValue float64
		switch white {
		case true:
			previousChosen.valueOfMove = -100000.00
		case false:
			previousChosen.valueOfMove = 100000.00
		}

		Chosen10 := e.getNext10(previousChosen)
		for _, chosen := range Chosen10 {
			if chosen == nil {
				continue
			}

			if depth != 0 {
				nextBestValue = e.bestNextMoveSingleThread(chosen, depth-1)
			}

			switch white {
			case true:
				if chosen.valueOfMove+nextBestValue > previousChosen.valueOfMove {
					previousChosen.valueOfMove = chosen.valueOfMove + nextBestValue
				}
			case false:
				if chosen.valueOfMove < previousChosen.valueOfMove {
					previousChosen.valueOfMove = chosen.valueOfMove + nextBestValue
				}
			}
		}

		wg.Done()
		result <- previousChosen
	}
}

func (e *Engine) bestNextMoveSingleThread(previousChosen *SelectedMove, depth int) float64 {
	white := !previousChosen.Piece.GetWhite()
	var chosenValue float64
	var nextBestValue float64
	switch white {
	case true:
		chosenValue = -100000.00
	case false:
		chosenValue = 100000.00
	}

	Chosen10 := e.getNext10(previousChosen)
	for _, chosen := range Chosen10 {
		if chosen == nil {
			continue
		}

		if depth != 0 {
			nextBestValue = e.bestNextMoveSingleThread(chosen, depth-1)
		}

		switch white {
		case true:
			if chosen.valueOfMove+nextBestValue > chosenValue {
				chosenValue = chosen.valueOfMove + nextBestValue
			}
		case false:
			if chosen.valueOfMove < chosenValue {
				chosenValue = chosen.valueOfMove + nextBestValue
			}
		}
	}

	return chosenValue
}

func (e *Engine) getInitial10() []*SelectedMove {
	var chosen10 []*SelectedMove
	var valuePerIndex []float64
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
			if counter < e.optionsSaved*2 {
				chosen10 = append(chosen10, move)
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
	chosen10 := make([]*SelectedMove, e.optionsSaved)
	valuePerIndex := make([]float64, e.optionsSaved)
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
			move.valueOfMove += newValue
			move.simulator = simulator

			if counter < e.optionsSaved {
				chosen10[counter] = move
				valuePerIndex[counter] = newValue
				counter++
			} else {
				index, worstValue := 0, 0.00
				switch piece.GetWhite() {
				case true:
					index, worstValue = minOfSlice(valuePerIndex)
					if worstValue < newValue {
						chosen10[index] = move
						valuePerIndex[index] = newValue
					}
				case false:
					index, worstValue = maxOfSlice(valuePerIndex)
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
	worstValue := -100000.00
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
	worstValue := 100000.00
	for i, value := range valuePerIndex {
		if value < worstValue {
			worstValue = value
			worstIndex = i
		}
	}

	return worstIndex, worstValue
}
