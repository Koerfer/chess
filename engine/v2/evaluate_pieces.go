package v2

import (
	v2 "chess/pieces/v2"
)

func (e *Engine) evaluatePieces() {
	var whiteOptionsCount int
	whiteChecked := false
	for pos, piece := range e.whiteBoard {
		if checkIfBackRank(piece, pos) {
			e.whiteEval -= 0.2
		}
		e.whiteEval += float64(piece.GetValue())
		whiteOptionsCount += len(piece.GetOptions())
		if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
			king := piece.(*v2.King)
			if king.Checked {
				whiteChecked = true
			}
		}

		if value, ok := e.safeWhiteFork(piece); ok {
			e.whiteEval += float64(value)
		}
	}
	if whiteOptionsCount == 0 && !whiteChecked {
		e.whiteEval += 1000
	} else if whiteOptionsCount == 0 {
		e.whiteEval -= 1000
	}

	var blackOptionsCount int
	blackChecked := false
	for pos, piece := range e.blackBoard {
		if checkIfBackRank(piece, pos) {
			e.blackEval -= 0.2
		}
		e.blackEval += float64(piece.GetValue())
		blackOptionsCount += len(piece.GetOptions())
		if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
			king := piece.(*v2.King)
			if king.Checked {
				blackChecked = true
			}
		}

		if value, ok := e.safeBlackFork(piece); ok {
			e.blackEval += float64(value)
		}
	}
	if blackOptionsCount == 0 && !blackChecked {
		e.blackEval += 1000
	} else if blackOptionsCount == 0 {
		e.blackEval -= 1000
	}
}

func (e *Engine) safeWhiteFork(piece v2.PieceInterface) (int, bool) {
	count := 0
	value := 0
	if len(piece.GetAttackedBy()) == 0 {
		for attacked := range piece.GetOptions() {
			if piece, ok := e.blackBoard[attacked]; ok {
				if len(piece.GetProtectedBy()) == 0 {
					count++
					if piece.GetValue() > value {
						value = piece.GetValue()
					}
				}
			}
		}
	}

	if count >= 2 {
		return value, true
	}

	return 0, false
}

func (e *Engine) safeBlackFork(piece v2.PieceInterface) (int, bool) {
	count := 0
	value := 0
	if len(piece.GetAttackedBy()) == 0 {
		for attacked := range piece.GetOptions() {
			if piece, ok := e.whiteBoard[attacked]; ok {
				if len(piece.GetProtectedBy()) == 0 {
					count++
					if piece.GetValue() > value {
						value = piece.GetValue()
					}
				}
			}
		}
	}

	if count >= 2 {
		return value, true
	}

	return 0, false
}

func checkIfBackRank(piece v2.PieceInterface, pos int) bool {
	if piece.GetValue() == 5 {
		rook := piece.(*v2.Rook)
		if !rook.HasBeenMoved {
			return false
		}
	}
	switch piece.GetWhite() {
	case true:
		if pos/8 == 7 {
			return true
		}
	case false:
		if pos/8 == 0 {
			return true
		}
	}
	return false
}
