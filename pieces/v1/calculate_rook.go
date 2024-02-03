package v1

func (p *Piece) calculateRookMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) (map[int]struct{}, bool) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if p.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	for left := 1; left <= 8; left++ {
		newPosition := position - left
		if newPosition < 0 || newPosition > 63 {
			break
		}
		if rowPos-newPosition%8 < 0 {
			break
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			break
		}
		if _, ok := opponentBoard[newPosition]; ok {
			p.Options[newPosition] = value
			if opponentBoard[newPosition].Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
				for leftKing := left; leftKing <= 8; leftKing++ {
					newPosition := position - leftKing
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if rowPos-newPosition%8 < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			if opponentBoard[newPosition].Kind != King { // todo: optimisation to only calculate pins when needed
				for leftPin := left + 1; leftPin <= 8; leftPin++ {
					newPositionPin := position - leftPin
					if newPositionPin < 0 {
						break
					}
					if rowPos-newPositionPin%8 < 0 {
						break
					}
					if piece, ok := opponentBoard[newPositionPin]; ok {
						if piece.Kind == King {
							opponentBoard[newPosition].PinnedToKing = true
							opponentBoard[newPosition].PinnedByPosition = position
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for right := 1; right <= 8; right++ {
		newPosition := position + right
		if newPosition < 0 || newPosition > 63 {
			break
		}
		if newPosition%8-rowPos < 0 {
			break
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			break
		}
		if _, ok := opponentBoard[newPosition]; ok {
			p.Options[newPosition] = value
			if opponentBoard[newPosition].Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
				for rightKing := right; rightKing <= 8; rightKing++ {
					newPosition := position + rightKing
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if newPosition%8-rowPos < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			if opponentBoard[newPosition].Kind != King { // todo: optimisation to only calculate pins when needed
				for rightPin := right + 1; rightPin <= 8; rightPin++ {
					newPositionPin := position + rightPin
					if newPositionPin < 0 {
						break
					}
					if newPositionPin%8-rowPos < 0 {
						break
					}
					if piece, ok := opponentBoard[newPositionPin]; ok {
						if piece.Kind == King {
							opponentBoard[newPosition].PinnedToKing = true
							opponentBoard[newPosition].PinnedByPosition = position
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for down := 1; down <= 8; down++ {
		newPosition := position + down*8
		if newPosition < 0 || newPosition > 63 {
			break
		}
		if newPosition/8-colPos < 0 {
			break
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			break
		}
		if _, ok := opponentBoard[newPosition]; ok {
			p.Options[newPosition] = value
			if opponentBoard[newPosition].Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
				for downKing := down; downKing <= 8; downKing++ {
					newPosition := position + downKing*8
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if newPosition/8-colPos < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			if opponentBoard[newPosition].Kind != King { // todo: optimisation to only calculate pins when needed
				for downPin := down + 1; downPin <= 8; downPin++ {
					newPositionPin := position + downPin*8
					if newPositionPin < 0 {
						break
					}
					if newPositionPin/8-colPos < 0 {
						break
					}
					if piece, ok := opponentBoard[newPositionPin]; ok {
						if piece.Kind == King {
							opponentBoard[newPosition].PinnedToKing = true
							opponentBoard[newPosition].PinnedByPosition = position
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for up := 1; up <= 8; up++ {
		newPosition := position - up*8
		if newPosition < 0 || newPosition > 63 {
			break
		}
		if colPos-newPosition/8 < 0 {
			break
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			break
		}
		if _, ok := opponentBoard[newPosition]; ok {
			p.Options[newPosition] = value
			if opponentBoard[newPosition].Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
				for upKing := up; upKing <= 8; upKing++ {
					newPosition := position - upKing*8
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if colPos-newPosition/8 < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			if opponentBoard[newPosition].Kind != King { // todo: optimisation to only calculate pins when needed
				for upPin := up + 1; upPin <= 8; upPin++ {
					newPositionPin := position - upPin*8
					if newPositionPin < 0 {
						break
					}
					if colPos-newPositionPin/8 < 0 {
						break
					}
					if piece, ok := opponentBoard[newPositionPin]; ok {
						if piece.Kind == King {
							opponentBoard[newPosition].PinnedToKing = true
							opponentBoard[newPosition].PinnedByPosition = position
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}

	p.calculatePinnedOptions(position)

	return forbiddenSquares, check
}
