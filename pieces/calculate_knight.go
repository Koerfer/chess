package pieces

func (p *Piece) calculateKnightMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) (map[int]struct{}, bool) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if p.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	right := position % 8
	down := position / 8
	up := -8
	left := -1
	if right-2 >= 0 { // left 2 ok
		if down-1 >= 0 { // up 1 ok
			newPosition := position + left*2 + up
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
		if down+1 <= 7 { // down 1 ok
			newPosition := position + left*2 - up
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
	}

	if right+2 <= 7 {
		if down-1 >= 0 {
			newPosition := position - left*2 + up
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
		if down+1 <= 7 {
			newPosition := position - left*2 - up
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
	}

	if down+2 <= 7 {
		if right-1 >= 0 {
			newPosition := position - up*2 + left
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
		if right+1 <= 7 {
			newPosition := position - up*2 - left
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
	}

	if down-2 >= 0 {
		if right-1 >= 0 {
			newPosition := position + up*2 + left
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
		if right+1 <= 7 {
			newPosition := position + up*2 - left
			opponent, ok := opponentBoard[newPosition]
			if ok && opponent.Kind == King {
				check = true
				opponentBoard[newPosition].CheckingPieces[position] = p
			}
			p.Options[newPosition] = value
		}
	}
	for option := range p.Options {
		forbiddenSquares[option] = value
	}

	p.simpleDelete(myBoard)
	if p.PinnedToKing {
		p.Options = make(map[int]struct{})
	}

	return forbiddenSquares, check
}
