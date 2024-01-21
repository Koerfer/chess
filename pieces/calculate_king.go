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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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
	} else if newPosition < 0 || newPosition > 63 {
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

	switch p.White {
	case true:
		if p.HasBeenMoved {
			break
		}
		if rook, ok := whiteBoard[63]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[62]; !ok {
					if _, ok := whiteBoard[62]; !ok {
						if _, ok := whiteBoard[61]; !ok {
							p.Options[62] = value
						}
					}
				}
			}
		}
		if rook, ok := whiteBoard[56]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[58]; !ok {
					if _, ok := whiteBoard[59]; !ok {
						if _, ok := whiteBoard[58]; !ok {
							if _, ok := whiteBoard[57]; !ok {
								p.Options[58] = value
							}
						}
					}
				}
			}
		}

		if p.HasBeenMoved {
			break
		}
		if rook, ok := blackBoard[63]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[62]; !ok {
					p.Options[62] = value
				}
			}
		}
		if rook, ok := blackBoard[56]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[58]; !ok {
					p.Options[58] = value
				}
			}
		}
	case false:
		if p.HasBeenMoved {
			break
		}
		if rook, ok := blackBoard[7]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[6]; !ok {
					if _, ok := blackBoard[6]; !ok {
						if _, ok := blackBoard[5]; !ok {
							p.Options[6] = value
						}
					}
				}
			}
		}
		if rook, ok := blackBoard[0]; ok {
			if !rook.HasBeenMoved {
				if _, ok := forbiddenSquares[2]; !ok {
					if _, ok := blackBoard[3]; !ok {
						if _, ok := blackBoard[2]; !ok {
							if _, ok := blackBoard[1]; !ok {
								p.Options[2] = value
							}
						}
					}
				}
			}
		}
	}

	return newForbiddenSquares
}
