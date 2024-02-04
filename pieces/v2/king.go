package v2

import "log"

type King struct {
	*Piece
	HasBeenMoved   bool
	CheckingPieces map[int]any
	Checked        bool
}

func (k *King) CalculateMoves(whiteBoard map[int]any, blackBoard map[int]any, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	newForbiddenSquares := make(map[int]struct{})
	if fixLastPosition {
		k.LastPosition = position
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
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position - 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position + 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position + 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position - 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position + 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position + 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition/8-colPos < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
	}

	newPosition = position - 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if colPos-newPosition/8 < 0 {
		// nothing
	} else if protectedPiece, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
		k.Protecting[newPosition] = protectedPiece
	} else if attackedPiece, ok := opponentBoard[newPosition]; ok {
		k.Options[newPosition] = value
		k.addAttackedBy(attackedPiece, position)
	} else {
		k.Options[newPosition] = value
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

func (k *King) addAttackedBy(attackedPiece any, position int) {
	switch CheckPieceKindFromAny(attackedPiece) {
	case PieceKindPawn:
		pawn := attackedPiece.(*Pawn)
		pawn.AttackedBy[position] = k
	case PieceKindKnight:
		knight := attackedPiece.(*Knight)
		knight.AttackedBy[position] = k
	case PieceKindBishop:
		bishop := attackedPiece.(*Bishop)
		bishop.AttackedBy[position] = k
	case PieceKindRook:
		rook := attackedPiece.(*Rook)
		rook.AttackedBy[position] = k
	case PieceKindQueen:
		queen := attackedPiece.(*Queen)
		queen.AttackedBy[position] = k
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating attacked by king")
	}
}
