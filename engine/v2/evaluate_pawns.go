package v2

func (e *Engine) evaluatePawns() {
	whitePawnPositions := make(map[int]struct{})
	for position, piece := range e.whiteBoard {
		if piece.GetValue() != 1 { //pawn
			continue
		}

		for _, protector := range piece.GetProtectedBy() {
			if protector.GetValue() == 1 {
				e.whiteEval += 0.2
			} else if protector.GetValue() == 100 {
				// king protecting no good
			} else {
				if len(protector.GetAttackedBy()) == 0 {
					e.whiteEval += 0.1
				}
			}
		}

		whitePawnPositions[position] = struct{}{}
	}
	e.evaluateWhitePawnPositions(whitePawnPositions)

	blackPawnPositions := make(map[int]struct{})
	for position, piece := range e.blackBoard {
		if piece.GetValue() != 1 { //pawn
			continue
		}

		for _, protector := range piece.GetProtectedBy() {
			if protector.GetValue() == 1 {
				e.blackEval += 0.2
			} else if protector.GetValue() == 100 {
				// king protecting no good
			} else {
				if len(protector.GetAttackedBy()) == 0 {
					e.blackEval += 0.1
				}
			}
		}

		blackPawnPositions[position] = struct{}{}
	}
	e.evaluateBlackPawnPositions(blackPawnPositions)
	e.weakPositions()
}

func (e *Engine) evaluateWhitePawnPositions(positions map[int]struct{}) {
	for i := 8; i < 64; i++ {
		if _, ok := positions[i]; ok {
			if _, ok := e.allWhiteOptions[i+8]; !ok {
				if i%8 == 0 {
					if _, ok := positions[i+1]; ok {
						e.superWeakWhiteSquares[i+8] = struct{}{}
					}
				} else if i%8 == 7 {
					if _, ok := positions[i-1]; ok {
						e.superWeakWhiteSquares[i+8] = struct{}{}
					}
				} else {
					both := true
					if _, ok := positions[i+1]; !ok {
						both = false
					}
					if _, ok := positions[i-1]; !ok {
						both = false
					}
					if both {
						e.superWeakWhiteSquares[i+8] = struct{}{}
					}
				}
			} else {
				if i%8 == 0 {
					if _, ok := positions[i+1]; ok {
						e.weakWhiteSquares[i+8] = struct{}{}
					}
				} else if i%8 == 7 {
					if _, ok := positions[i-1]; ok {
						e.weakWhiteSquares[i+8] = struct{}{}
					}
				} else {
					both := true
					if _, ok := positions[i+1]; !ok {
						both = false
					}
					if _, ok := positions[i-1]; !ok {
						both = false
					}
					if both {
						e.weakWhiteSquares[i+8] = struct{}{}
					}
				}
			}
		}
	}
}

func (e *Engine) evaluateBlackPawnPositions(positions map[int]struct{}) {
	for i := 55; i < 0; i-- {
		if _, ok := positions[i]; ok {
			if _, ok := e.allWhiteOptions[i-8]; !ok {
				if i%8 == 0 {
					if _, ok := positions[i+1]; ok {
						e.superWeakWhiteSquares[i-8] = struct{}{}
					}
				} else if i%8 == 7 {
					if _, ok := positions[i-1]; ok {
						e.superWeakWhiteSquares[i-8] = struct{}{}
					}
				} else {
					both := true
					if _, ok := positions[i+1]; !ok {
						both = false
					}
					if _, ok := positions[i-1]; !ok {
						both = false
					}
					if both {
						e.superWeakWhiteSquares[i-8] = struct{}{}
					}
				}
			} else {
				if i%8 == 0 {
					if _, ok := positions[i+1]; ok {
						e.weakWhiteSquares[i-8] = struct{}{}
					}
				} else if i%8 == 7 {
					if _, ok := positions[i-1]; ok {
						e.weakWhiteSquares[i-8] = struct{}{}
					}
				} else {
					both := true
					if _, ok := positions[i+1]; !ok {
						both = false
					}
					if _, ok := positions[i-1]; !ok {
						both = false
					}
					if both {
						e.weakWhiteSquares[i-8] = struct{}{}
					}
				}
			}
		}
	}
}

func (e *Engine) weakPositions() {
	e.whiteEval -= float64(len(e.superWeakWhiteSquares)) * 0.2
	e.whiteEval -= float64(len(e.weakWhiteSquares)) * 0.1
	e.blackEval -= float64(len(e.superWeakBlackSquares)) * 0.2
	e.blackEval -= float64(len(e.weakBlackSquares)) * 0.1
}
