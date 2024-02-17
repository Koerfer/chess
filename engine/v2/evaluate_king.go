package v2

import v2 "chess/pieces/v2"

func (e *Engine) evaluateKingPosition() {
	protectedByPawnValue := 0.2
	protectedByOtherValue := 0.1
	for position, piece := range e.whiteBoard {
		if v2.CheckPieceKindFromAny(piece) != v2.PieceKindKing {
			continue
		}

		king := piece.(*v2.King)
		if king.HasCastled {
			e.whiteEval += 0.5
		}
		if !king.HasCastled && king.HasBeenMoved {
			e.whiteEval -= 2
		}

		if position%8 == 7 {
			if p, ok := e.whiteBoard[position-8]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-9]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-16]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-17]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
		} else if position%8 == 0 {
			if p, ok := e.whiteBoard[position-8]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-7]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-16]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-15]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
		} else {
			if p, ok := e.whiteBoard[position-7]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-8]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-9]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-15]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-16]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
			if p, ok := e.whiteBoard[position-17]; ok {
				if p.GetValue() == 1 { //pawn
					e.whiteEval += protectedByPawnValue
				} else {
					e.whiteEval += protectedByOtherValue
				}
			}
		}

		if position/8 == 7 {
			if _, ok := e.whiteBoard[position-8]; ok {
				if _, ok := e.whiteBoard[position-7]; ok {
					if _, ok := e.whiteBoard[position-9]; ok {
						e.whiteEval -= 1
					}
				}
			}
		}

		e.whiteEval += float64(position/8-7) * protectedByOtherValue

		break
	}

	for position, piece := range e.blackBoard {
		if v2.CheckPieceKindFromAny(piece) != v2.PieceKindKing {
			continue
		}

		king := piece.(*v2.King)
		if king.HasCastled {
			e.blackEval += 0.5
		}
		if !king.HasCastled && king.HasBeenMoved {
			e.blackEval -= 2
		}

		if position%8 == 7 {
			if p, ok := e.blackBoard[position+8]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+7]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+16]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+15]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
		} else if position%8 == 0 {
			if p, ok := e.blackBoard[position+8]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+9]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+16]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+17]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
		} else {
			if p, ok := e.blackBoard[position+7]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+8]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+9]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+15]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+16]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
			if p, ok := e.blackBoard[position+17]; ok {
				if p.GetValue() == 1 { //pawn
					e.blackEval += protectedByPawnValue
				} else {
					e.blackEval += protectedByOtherValue
				}
			}
		}

		if position/8 == 0 {
			if _, ok := e.blackBoard[position+8]; ok {
				if _, ok := e.blackBoard[position+7]; ok {
					if _, ok := e.blackBoard[position+9]; ok {
						e.blackEval -= 1
					}
				}
			}
		}

		e.blackEval -= float64(position/8) * protectedByOtherValue

		break
	}
}
