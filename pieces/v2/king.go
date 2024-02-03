package v2

type King struct {
	*Piece
	HasBeenMoved   bool
	CheckingPieces map[int]any
	Checked        bool
}

func CalculateKingMoves(king *King, whiteBoard map[int]any, blackBoard map[int]any, position int, forbiddenSquares map[int]struct{}) map[int]struct{} {
	newForbiddenSquares := make(map[int]struct{})

	myBoard := whiteBoard
	opponentBoard := blackBoard

	if _, ok := forbiddenSquares[position]; ok {
		king.Checked = true
	}

	if king.White == false {
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
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position - 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position + 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position + 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position - 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position + 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position + 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if newPosition/8-colPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	newPosition = position - 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 || newPosition > 63 {
		// nothing
	} else if colPos-newPosition/8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		king.Options[newPosition] = value
	} else {
		king.Options[newPosition] = value
	}

	switch king.White {
	case true:
		if king.HasBeenMoved {
			break
		}
		rook, ok := whiteBoard[63]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[62]; !ok {
					if _, ok := whiteBoard[62]; !ok {
						if _, ok := whiteBoard[61]; !ok {
							king.Options[62] = value
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
								king.Options[58] = value
							}
						}
					}
				}
			}
		}
	case false:
		if king.HasBeenMoved {
			break
		}
		rook, ok := blackBoard[7]
		if ok && CheckPieceKindFromAny(rook) == PieceKindRook {
			rook := rook.(*Rook)
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[6]; !ok {
					if _, ok := blackBoard[6]; !ok {
						if _, ok := blackBoard[5]; !ok {
							king.Options[6] = value
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
								king.Options[2] = value
							}
						}
					}
				}
			}
		}
	}

	return newForbiddenSquares
}
