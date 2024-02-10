package v2

import (
	v2 "chess/pieces/v2"
	"fmt"
	"sort"
)

type Engine struct {
	whiteBoard map[int]v2.PieceInterface
	blackBoard map[int]v2.PieceInterface

	whiteEval int
	blackEval int
}

func (e *Engine) Init(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface) {
	e.copyBoard(whiteBoard, blackBoard)
	e.evaluate()
	whiteEval, blackEval := e.getTotalValue()

	fmt.Printf("white: %d\nblack: %d\n\n", whiteEval, blackEval)
}

func (e *Engine) evaluate() {
	for _, piece := range e.blackBoard {
		if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {

		} else {
			ifCaptured := calculateBestOutcomeIfCaptured(piece)
			ifCapturing := e.calculateBestOutcomeIfCapturing(piece)
			piece.SetEvaluatedValue(maxInt(ifCaptured, ifCapturing) + piece.GetValue())
		}
	}

	for _, piece := range e.whiteBoard {
		if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {

		} else {
			ifCaptured := calculateBestOutcomeIfCaptured(piece)
			ifCapturing := e.calculateBestOutcomeIfCapturing(piece)
			piece.SetEvaluatedValue(maxInt(ifCaptured, ifCapturing) + piece.GetValue())
		}
	}
}

func maxInt(a int, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return a
	}
	return b
}

func calculateBestOutcomeIfCaptured(piece v2.PieceInterface) int {
	if len(piece.GetAttackedBy()) == 0 {
		return 0
	}
	var attackerValues []int
	var defenderValues []int
	for _, attacker := range piece.GetAttackedBy() {
		attackerValues = append(attackerValues, attacker.GetValue())
	}
	for _, defender := range piece.GetProtectedBy() {
		defenderValues = append(defenderValues, defender.GetValue())
	}

	sort.Ints(attackerValues)
	sort.Ints(defenderValues)

	var bestValue int
	if len(attackerValues) == len(defenderValues) {
		bestValue -= piece.GetValue()
		currentValue := bestValue

		for n := range attackerValues {
			if n == len(defenderValues)-1 {
				currentValue += attackerValues[n]
			} else {
				currentValue += attackerValues[n] - defenderValues[n]
			}
			if currentValue > bestValue {
				bestValue = currentValue
			}
		}
	} else if len(attackerValues) > len(defenderValues) {
		bestValue -= piece.GetValue()
		currentValue := bestValue

		for n := range defenderValues {
			if n == len(defenderValues)-1 {
				currentValue += -defenderValues[n]
			} else {
				currentValue += attackerValues[n] - defenderValues[n]
			}
			if currentValue > bestValue {
				bestValue = currentValue
			}
		}
	} else if len(attackerValues) < len(defenderValues) {
		bestValue -= piece.GetValue()
		currentValue := bestValue

		for n := range attackerValues {
			if n == len(attackerValues)-1 {
				currentValue += attackerValues[n]
			} else {
				currentValue += attackerValues[n] - defenderValues[n]
			}
			if currentValue > bestValue {
				bestValue = currentValue
			}
		}
	}

	return bestValue
}

func (e *Engine) calculateBestOutcomeIfCapturing(piece v2.PieceInterface) int {
	var bestValue int
	white := piece.GetWhite()

	for option := range piece.GetOptions() {
		switch white {
		case true:
			if p, ok := e.blackBoard[option]; ok {
				if v2.CheckPieceKindFromAny(p) == v2.PieceKindKing {

				} else {
					if len(p.GetProtectedBy()) == 0 {
						if p.GetValue() > bestValue {
							bestValue = p.GetValue()
						}
					} else {
						if p.GetValue()-piece.GetValue() > bestValue {
							bestValue = p.GetValue()
						}
					}
				}
			}
		case false:
			if p, ok := e.whiteBoard[option]; ok {
				if v2.CheckPieceKindFromAny(p) == v2.PieceKindKing {

				} else {
					if len(p.GetProtectedBy()) == 0 {
						if p.GetValue() > bestValue {
							bestValue = p.GetValue()
						}
					} else {
						if p.GetValue()-piece.GetValue() > bestValue {
							bestValue = p.GetValue()
						}
					}
				}
			}
		}
	}

	return bestValue
}

func (e *Engine) getTotalValue() (int, int) {
	var white int
	var black int
	for _, piece := range e.whiteBoard {
		white += piece.GetEvaluatedValue()
	}
	for _, piece := range e.blackBoard {
		black += piece.GetEvaluatedValue()
	}

	return white, black
}
