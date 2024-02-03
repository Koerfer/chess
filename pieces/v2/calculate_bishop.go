package v2

func (p *Piece) calculateBishopMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) (map[int]struct{}, bool) {
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

	for leftUp := 1; leftUp <= colPos; leftUp++ {
		newPosition := position - leftUp*9
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
				for leftUpKing := leftUp; leftUpKing <= 8; leftUpKing++ {
					newPosition := position - leftUpKing*9
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
				for leftUpPin := leftUp + 1; leftUpPin <= 8; leftUpPin++ {
					newPositionPin := position - leftUpPin*9
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
							opponentBoard[newPosition].PinnedByPiece = p
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for rightUp := 1; rightUp <= colPos; rightUp++ {
		newPosition := position - rightUp*7
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
				for rightUpKing := rightUp; rightUpKing <= 8; rightUpKing++ {
					newPosition := position - rightUpKing*7
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
				for rightUpPin := rightUp + 1; rightUpPin <= 8; rightUpPin++ {
					newPositionPin := position - rightUpPin*7
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
							opponentBoard[newPosition].PinnedByPiece = p
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for leftDown := 1; leftDown <= 7-colPos; leftDown++ {
		newPosition := position + leftDown*7
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
				for leftDownKing := leftDown; leftDownKing <= 8; leftDownKing++ {
					newPosition := position + leftDownKing*7
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
				for leftDownPin := leftDown + 1; leftDownPin <= 8; leftDownPin++ {
					newPositionPin := position + leftDownPin*7
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
							opponentBoard[newPosition].PinnedByPiece = p
						}
						break
					}
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for rightDown := 1; rightDown <= 7-colPos; rightDown++ {
		newPosition := position + rightDown*9
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
				for rightDownKing := rightDown; rightDownKing <= 8; rightDownKing++ {
					newPosition := position + rightDownKing*9
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
				for rightDownPin := rightDown + 1; rightDownPin <= 8; rightDownPin++ {
					newPositionPin := position + rightDownPin*9
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
							opponentBoard[newPosition].PinnedByPiece = p
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
