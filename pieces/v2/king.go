package v2

type King struct {
	Value          int
	EvaluatedValue int
	White          bool
	LastPosition   int
	Options        map[int]struct{}
	Protecting     map[int]PieceInterface

	HasBeenMoved   bool
	CheckingPieces map[int]PieceInterface
	Checked        bool
}

func (k *King) GetValue() int {
	return k.Value
}
func (k *King) GetEvaluatedValue() int {
	return k.EvaluatedValue
}
func (k *King) GetWhite() bool {
	return k.White
}
func (k *King) GetLastPosition() int {
	return k.LastPosition
}
func (k *King) GetOptions() map[int]struct{} {
	return k.Options
}
func (k *King) GetProtecting() map[int]PieceInterface {
	return k.Protecting
}
func (k *King) GetAttackedBy() map[int]PieceInterface {
	panic("wrong use of GetAttackedBy for king")
	return nil
}
func (k *King) GetProtectedBy() map[int]PieceInterface {
	panic("wrong use of GetProtectedBy for king")
	return nil
}
func (k *King) SetValue(value int) {
	k.Value = value
}
func (k *King) SetEvaluatedValue(evaluatedValue int) {
	k.EvaluatedValue = evaluatedValue
}
func (k *King) SetWhite(white bool) {
	k.White = white
}
func (k *King) SetLastPosition(lastPosition int) {
	k.LastPosition = lastPosition
}
func (k *King) SetOptions(options map[int]struct{}) {
	k.Options = options
}
func (k *King) SetProtecting(protecting map[int]PieceInterface) {
	k.Protecting = protecting
}
func (k *King) SetAttackedBy(attackedBy map[int]PieceInterface) {
	panic("wrong use of SetAttackedBy for king")
	return
}
func (k *King) SetProtectedBy(protectedBy map[int]PieceInterface) {
	panic("wrong use of SetProtectedBy for king")
	return
}

func (k *King) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	newForbiddenSquares := make(map[int]struct{})
	if fixLastPosition {
		k.LastPosition = position
	} else {
		k.CheckingPieces = make(map[int]PieceInterface)
		k.Checked = false
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard

	if _, ok := forbiddenSquares[position]; ok {
		k.Checked = true
	}

	if k.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	newPosition := position - 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position - 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position + 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position + 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position - 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position + 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position + 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition/8-colPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	newPosition = position - 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if colPos-newPosition/8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		addAttackedBy(k, attackedPiece, position)
		newForbiddenSquares[newPosition] = value
	} else {
		k.Options[newPosition] = value
		newForbiddenSquares[newPosition] = value
	}

	switch k.White {
	case true:
		if k.HasBeenMoved {
			break
		}
		rook, ok := whiteBoard[63]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[62]; !ok {
					if _, ok := whiteBoard[62]; !ok {
						if _, ok := whiteBoard[61]; !ok {
							k.Options[62] = value
						}
					}
				}
			}
		}
		rook, ok = whiteBoard[56]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[58]; !ok {
					if _, ok := whiteBoard[59]; !ok {
						if _, ok := whiteBoard[58]; !ok {
							if _, ok := whiteBoard[57]; !ok {
								k.Options[58] = value
							}
						}
					}
				}
			}
		}
	case false:
		if k.HasBeenMoved {
			break
		}
		rook, ok := blackBoard[7]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[6]; !ok {
					if _, ok := blackBoard[6]; !ok {
						if _, ok := blackBoard[5]; !ok {
							k.Options[6] = value
						}
					}
				}
			}
		}
		rook, ok = blackBoard[0]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[2]; !ok {
					if _, ok := blackBoard[3]; !ok {
						if _, ok := blackBoard[2]; !ok {
							if _, ok := blackBoard[1]; !ok {
								k.Options[2] = value
							}
						}
					}
				}
			}
		}
	}

	return newForbiddenSquares
}

func (k *King) Copy(deep bool) PieceInterface {
	if k == nil {
		return nil
	}
	copyCat := &King{
		Value:          k.Value,
		EvaluatedValue: k.EvaluatedValue,
		White:          k.White,
		LastPosition:   k.LastPosition,
		Options:        k.Options,
		HasBeenMoved:   k.HasBeenMoved,
		Checked:        k.Checked,
	}
	if deep {
		copyCat.Protecting, copyCat.CheckingPieces, _ = copyProtectingAndAttacking(k.Protecting, k.CheckingPieces, nil)
	}
	return copyCat
}
