package v1

func (p *Piece) calculatePawnMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int, fixLastPosition bool) (map[int]struct{}, bool) {
	forbiddenSquares := make(map[int]struct{})
	p.EnPassantOptions = make(map[int]int)
	if fixLastPosition {
		p.LastPosition = position
	}
	var check bool

	myBoard := whiteBoard
	opponentBoard := blackBoard
	endPosition := 0
	offsetMultiplier := 1
	offsetAddition := 0
	if p.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
		endPosition = 7
		offsetMultiplier = -1
		offsetAddition = 2
	}

	// comments are from white perspective, flipped for black
	if position%8 == 0 { // if left on board
		captureOption := position - (7+offsetAddition)*offsetMultiplier
		forbiddenSquares[captureOption] = value
		if _, ok := opponentBoard[captureOption]; ok { // if black Piece up right
			p.Options[captureOption] = value // add capture move
			if opponentBoard[captureOption].Kind == King {
				check = true
				opponentBoard[captureOption].CheckingPieces[position] = p
			}
		}
		if position/8 == endPosition+offsetMultiplier*3 {
			if piece, ok := opponentBoard[position+1]; ok && piece.Kind == Pawn && piece.LastPosition == (endPosition+offsetMultiplier)*8+(position+1)%8 {
				p.EnPassantOptions[captureOption] = position + 1
			}
		}
	} else if position%8 == 7 { // if right on board
		captureOption := position - (9-offsetAddition)*offsetMultiplier
		forbiddenSquares[captureOption] = value
		if _, ok := opponentBoard[captureOption]; ok { // if black Piece up left
			p.Options[captureOption] = value // add capture move
			if opponentBoard[captureOption].Kind == King {
				check = true
				opponentBoard[captureOption].CheckingPieces[position] = p
			}
		}
		if position/8 == endPosition+offsetMultiplier*3 {
			if piece, ok := opponentBoard[position-1]; ok && piece.Kind == Pawn && piece.LastPosition == (endPosition+offsetMultiplier)*8+(position-1)%8 {
				p.EnPassantOptions[captureOption] = position - 1
			}
		}
	} else { // if in middle
		captureOption := position - (7+offsetAddition)*offsetMultiplier
		forbiddenSquares[captureOption] = value
		if _, ok := opponentBoard[captureOption]; ok { // if black Piece up right
			p.Options[captureOption] = value // add capture move
			if opponentBoard[captureOption].Kind == King {
				check = true
				opponentBoard[captureOption].CheckingPieces[position] = p
			}
		}
		if position/8 == endPosition+offsetMultiplier*3 {
			if piece, ok := opponentBoard[position+1]; ok && piece.Kind == Pawn && piece.LastPosition == (endPosition+offsetMultiplier)*8+(position+1)%8 {
				p.EnPassantOptions[captureOption] = position + 1
			}
		}

		captureOption = position - (9-offsetAddition)*offsetMultiplier
		forbiddenSquares[captureOption] = value
		if _, ok := opponentBoard[captureOption]; ok { // if black Piece up left
			p.Options[captureOption] = value // add capture move
			if opponentBoard[captureOption].Kind == King {
				check = true
				opponentBoard[captureOption].CheckingPieces[position] = p
			}
		}
		if position/8 == endPosition+offsetMultiplier*3 {
			if piece, ok := opponentBoard[position-1]; ok && piece.Kind == Pawn && piece.LastPosition == (endPosition+offsetMultiplier)*8+(position-1)%8 {
				p.EnPassantOptions[captureOption] = position - 1
			}
		}
	}

	oneUp := position - 8*offsetMultiplier
	twoUp := position - 16*offsetMultiplier
	if position/8 == (endPosition + 6*offsetMultiplier) {
		if _, ok := opponentBoard[position-16*offsetMultiplier]; ok {
			p.Options[oneUp] = value
			p.calculatePinnedOptions(position)
			p.simpleDelete(myBoard)
			return forbiddenSquares, check
		}
		if _, ok := myBoard[oneUp]; !ok {
			if _, ok := opponentBoard[oneUp]; !ok {
				p.Options[oneUp] = value
				if _, ok := myBoard[twoUp]; !ok {
					if _, ok := opponentBoard[oneUp]; !ok {
						p.Options[twoUp] = value
					}
				}
			}

		}
		p.calculatePinnedOptions(position)
		return forbiddenSquares, check
	}
	if _, ok := opponentBoard[oneUp]; ok {
		p.calculatePinnedOptions(position)
		return forbiddenSquares, check
	}
	p.Options[position-8*offsetMultiplier] = value
	p.calculatePinnedOptions(position)
	p.simpleDelete(myBoard)
	return forbiddenSquares, check
}
