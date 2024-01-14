package pieces

func (p *Piece) calculateKingMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int, forbiddenSquares map[int]struct{}) map[int]struct{} {
	newForbiddenSquares := make(map[int]struct{})

	myBoard := whiteBoard
	opponentBoard := blackBoard

	if _, ok := forbiddenSquares[position]; ok {
		p.Checked = true
	}

	if p.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	newPosition := position - 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position - 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position + 7
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position + 9
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position - 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition < 0 {
		// nothing
	} else if rowPos-newPosition%8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position + 1
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition%8-rowPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position + 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if newPosition/8-colPos < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	newPosition = position - 8
	if _, ok := forbiddenSquares[newPosition]; ok {
		// nothing
	} else if colPos-newPosition/8 < 0 {
		// nothing
	} else if _, ok := myBoard[newPosition]; ok {
		newForbiddenSquares[newPosition] = value
	} else if _, ok := opponentBoard[newPosition]; ok {
		p.Options[newPosition] = value
	} else {
		p.Options[newPosition] = value
	}

	return newForbiddenSquares
}
