package pieces

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
		if newPosition < 0 {
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
				for leftUpKing := leftUp; leftUpKing <= 8; leftUpKing++ {
					newPosition := position - leftUpKing*9
					if newPosition < 0 {
						break
					}
					if rowPos-newPosition%8 < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for rightUp := 1; rightUp <= colPos; rightUp++ {
		newPosition := position - rightUp*7
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
				for rightUpKing := rightUp; rightUpKing <= 8; rightUpKing++ {
					newPosition := position - rightUpKing*7
					if newPosition < 0 {
						break
					}
					if newPosition%8-rowPos < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for leftDown := 1; leftDown <= 7-colPos; leftDown++ {
		newPosition := position + leftDown*7
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
				for leftDownKing := leftDown; leftDownKing <= 8; leftDownKing++ {
					newPosition := position + leftDownKing*7
					if newPosition < 0 {
						break
					}
					if rowPos-newPosition%8 < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for rightDown := 1; rightDown <= 7-colPos; rightDown++ {
		newPosition := position + rightDown*9
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
				for rightDownKing := rightDown; rightDownKing <= 8; rightDownKing++ {
					newPosition := position + rightDownKing*9
					if newPosition < 0 {
						break
					}
					if newPosition%8-rowPos < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}

	return forbiddenSquares, check
}
