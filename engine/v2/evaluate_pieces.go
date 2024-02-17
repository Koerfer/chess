package v2

func (e *Engine) evaluatePieces() {
	e.countValues()

}

func (e *Engine) countValues() {
	var whiteOptionsCount int
	for _, piece := range e.whiteBoard {
		e.whiteEval += float64(piece.GetValue())
		whiteOptionsCount += len(piece.GetOptions())

		//e.whiteEval += attackerDefenders(piece)
	}
	if whiteOptionsCount == 0 {
		e.whiteEval -= 1000
	}

	var blackOptionsCount int
	for _, piece := range e.blackBoard {
		e.blackEval += float64(piece.GetValue())
		blackOptionsCount += len(piece.GetOptions())

		//e.whiteEval += attackerDefenders(piece)
	}
	if blackOptionsCount == 0 {
		e.blackEval -= 1000
	}
}

//func attackerDefenders(piece v2.PieceInterface) float64 {
//	var sumAttackerValues int
//	var sumDefenderValues int
//	for _, attacker := range piece.GetAttackedBy() {
//		if attacker.GetValue() == 100 {
//			continue
//		}
//		sumAttackerValues += attacker.GetValue()
//	}
//	for _, defender := range piece.GetProtectedBy() {
//		if defender.GetValue() == 100 {
//			continue
//		}
//		sumDefenderValues += defender.GetValue()
//	}
//
//	return float64(sumDefenderValues - sumAttackerValues)
//}
