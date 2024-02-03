package v2

type Pawn struct {
	*Piece
	EnPassantOptions map[int]int
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
}

func CalculatePawnMoves(pawn *Pawn, whiteBoard map[int]any, blackBoard map[int]any, position int, fixLastPosition bool) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})
	pawn.EnPassantOptions = make(map[int]int)
	if fixLastPosition {
		pawn.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	endPosition := 0
	offsetMultiplier := 1
	offsetAddition := 0
	if pawn.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
		endPosition = 7
		offsetMultiplier = -1
		offsetAddition = 2
	}

	// comments are from white perspective, flipped for black
	if position%8 == 0 { // if left on board
		captureOption := position - (7+offsetAddition)*offsetMultiplier
		calculateCaptureOption(pawn, position, endPosition, captureOption, 1, offsetMultiplier, forbiddenSquares, opponentBoard)
	} else if position%8 == 7 { // if right on board
		captureOption := position - (9-offsetAddition)*offsetMultiplier
		calculateCaptureOption(pawn, position, endPosition, captureOption, -1, offsetMultiplier, forbiddenSquares, opponentBoard)
	} else { // if in middle
		captureOption := position - (7+offsetAddition)*offsetMultiplier
		calculateCaptureOption(pawn, position, endPosition, captureOption, 1, offsetMultiplier, forbiddenSquares, opponentBoard)
		captureOption = position - (9-offsetAddition)*offsetMultiplier
		calculateCaptureOption(pawn, position, endPosition, captureOption, -1, offsetMultiplier, forbiddenSquares, opponentBoard)
	}

	oneUp := position - 8*offsetMultiplier
	twoUp := position - 16*offsetMultiplier
	if position/8 == (endPosition + 6*offsetMultiplier) {
		if _, ok := opponentBoard[position-16*offsetMultiplier]; ok {
			pawn.Options[oneUp] = value
			pawn.calculatePinnedOptions(position)
			pawnDelete(pawn, myBoard)
			return forbiddenSquares
		}
		if _, ok := myBoard[oneUp]; !ok {
			if _, ok := opponentBoard[oneUp]; !ok {
				pawn.Options[oneUp] = value
				if _, ok := myBoard[twoUp]; !ok {
					if _, ok := opponentBoard[oneUp]; !ok {
						pawn.Options[twoUp] = value
					}
				}
			}

		}
		pawn.calculatePinnedOptions(position)
		return forbiddenSquares
	}
	if _, ok := opponentBoard[oneUp]; ok {
		pawn.calculatePinnedOptions(position)
		return forbiddenSquares
	}
	pawn.Options[position-8*offsetMultiplier] = value
	pawn.calculatePinnedOptions(position)
	pawnDelete(pawn, myBoard)
	return forbiddenSquares
}

func calculateCaptureOption(pawn *Pawn, position int, endPosition int, captureOption int, direction int, offsetMultiplier int, forbiddenSquares map[int]struct{}, opponentBoard map[int]any) {
	forbiddenSquares[captureOption] = value
	if capturePiece, ok := opponentBoard[captureOption]; ok { // if black Piece up right
		pawn.Options[captureOption] = value // add capture move
		if CheckPieceKindFromAny(capturePiece) == PieceKindKing {
			p := capturePiece.(*King)
			p.Checked = true
			p.CheckingPieces[position] = pawn
		}
	}
	if position/8 == endPosition+offsetMultiplier*3 {
		piece, ok := opponentBoard[position+direction]
		if !ok {
			// do nothing
		} else {
			if CheckPieceKindFromAny(piece) == PieceKindPawn {
				p := piece.(*Pawn)
				if p.LastPosition == (endPosition+offsetMultiplier)*8+(position+direction)%8 {
					pawn.EnPassantOptions[captureOption] = position + direction
				}
			}
		}
	}
}

func (p *Pawn) calculatePinnedOptions(position int) {
	if p.PinnedToKing {
		for option := range p.Options {
			if option == p.PinnedByPosition {
				continue
			}
			if p.PinnedByPosition%8 == position%8 { // same column
				if p.PinnedByPosition/8 < position/8 {
					if option%8 == position%8 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition/8 > position/8 {
					if option%8 == position%8 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition/8 == position/8 { // same row
				if p.PinnedByPosition%8 < position%8 {
					if option/8 == position/8 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition%8 > position%8 {
					if option/8 == position/8 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition%9 == position%9 { // same left->right column
				if p.PinnedByPosition < position {
					if option%9 == position%9 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition > position {
					if option%9 == position%9 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition%7 == position%7 { // same right->left column
				if p.PinnedByPosition < position {
					if option%7 == position%7 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition > position {
					if option%7 == position%7 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			}
			delete(p.Options, option)
		}
	}
}

func pawnDelete(pawn *Pawn, board map[int]any) {
	var toRemove []int
	for option, _ := range pawn.Options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(pawn.Options, toDelete)
	}
}
