package pieces

func (p *Piece) calculateRookMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})

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
				for leftKing := left; leftKing <= 8; leftKing++ {
					newPosition := position - leftKing
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
	for right := 1; right <= 8; right++ {
		newPosition := position + right
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
				for rightKing := right; rightKing <= 8; rightKing++ {
					newPosition := position + rightKing
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
	for down := 1; down <= 8; down++ {
		newPosition := position + down*8
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
				for downKing := down; downKing <= 8; downKing++ {
					newPosition := position + downKing*8
					if newPosition < 0 {
						break
					}
					if newPosition/8-colPos < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}
	for up := 1; up <= 8; up++ {
		newPosition := position - up*8
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
				for upKing := up; upKing <= 8; upKing++ {
					newPosition := position - upKing*8
					if newPosition < 0 {
						break
					}
					if colPos-newPosition/8 < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			}
			break
		}

		p.Options[newPosition] = value
	}

	return forbiddenSquares
}
