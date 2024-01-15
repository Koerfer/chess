package pieces

func (p *Piece) calculatePawnMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int, fixLastPosition bool) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})
	p.EnPassantOptions = make(map[int]int)
	if fixLastPosition {
		p.LastPosition = position
	}

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
		}
		if position/8 == endPosition+offsetMultiplier*3 {
			if piece, ok := opponentBoard[position-1]; ok && piece.Kind == Pawn && piece.LastPosition == (endPosition+offsetMultiplier)*8+(position-1)%8 {
				p.EnPassantOptions[captureOption] = position - 1
			}
		}
	}

	if position/8 == (endPosition + 6*offsetMultiplier) {
		if _, ok := opponentBoard[position-16*offsetMultiplier]; ok {
			p.Options[position-8*offsetMultiplier] = value
			p.simpleDelete(myBoard)
			return forbiddenSquares
		}
		p.Options[position-8*offsetMultiplier] = value
		p.Options[position-16*offsetMultiplier] = value
		p.goFarDelete(myBoard)
		return forbiddenSquares
	}
	if _, ok := opponentBoard[position-8*offsetMultiplier]; ok {
		return forbiddenSquares
	}
	p.Options[position-8*offsetMultiplier] = value
	p.simpleDelete(myBoard)
	return forbiddenSquares
}
