package v2

import "log"

type Pawn struct {
	*Piece
	EnPassantOptions map[int]int
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
	AttackedBy       map[int]any
}

func (p *Pawn) CalculateMoves(whiteBoard map[int]any, blackBoard map[int]any, position int, fixLastPosition bool) map[int]struct{} {
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
		p.calculateCaptureOption(position, endPosition, captureOption, 1, offsetMultiplier, forbiddenSquares, myBoard, opponentBoard)
	} else if position%8 == 7 { // if right on board
		captureOption := position - (9-offsetAddition)*offsetMultiplier
		p.calculateCaptureOption(position, endPosition, captureOption, -1, offsetMultiplier, forbiddenSquares, myBoard, opponentBoard)
	} else { // if in middle
		captureOption := position - (7+offsetAddition)*offsetMultiplier
		p.calculateCaptureOption(position, endPosition, captureOption, 1, offsetMultiplier, forbiddenSquares, myBoard, opponentBoard)
		captureOption = position - (9-offsetAddition)*offsetMultiplier
		p.calculateCaptureOption(position, endPosition, captureOption, -1, offsetMultiplier, forbiddenSquares, myBoard, opponentBoard)
	}

	oneUp := position - 8*offsetMultiplier
	twoUp := position - 16*offsetMultiplier
	if position/8 == (endPosition + 6*offsetMultiplier) {
		if _, ok := opponentBoard[position-16*offsetMultiplier]; ok {
			p.Options[oneUp] = value
			p.calculatePinnedOptions(position)
			p.deleteOptions(myBoard)
			return forbiddenSquares
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
		return forbiddenSquares
	}
	if _, ok := opponentBoard[oneUp]; ok {
		p.calculatePinnedOptions(position)
		return forbiddenSquares
	}
	p.Options[position-8*offsetMultiplier] = value
	p.calculatePinnedOptions(position)
	p.deleteOptions(myBoard)
	return forbiddenSquares
}

func (p *Pawn) calculateCaptureOption(position int, endPosition int, captureOption int, direction int, offsetMultiplier int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) {
	forbiddenSquares[captureOption] = value
	if protectedPiece, ok := myBoard[captureOption]; ok {
		forbiddenSquares[captureOption] = value
		p.Protecting[captureOption] = protectedPiece
	}
	if capturePiece, ok := opponentBoard[captureOption]; ok { // if black Piece up right
		p.Options[captureOption] = value // add capture move
		if !p.PinnedToKing {
			p.addAttackedBy(opponent, position)
		}
		if CheckPieceKindFromAny(capturePiece) == PieceKindKing {
			king := capturePiece.(*King)
			king.Checked = true
			king.CheckingPieces[position] = king
		}
	}
	if position/8 == endPosition+offsetMultiplier*3 {
		piece, ok := opponentBoard[position+direction]
		if !ok {
			// do nothing
		} else {
			if CheckPieceKindFromAny(piece) == PieceKindPawn {
				pawn := piece.(*Pawn)
				if pawn.LastPosition == (endPosition+offsetMultiplier)*8+(position+direction)%8 {
					pawn.EnPassantOptions[captureOption] = position + direction
				}
			}
		}
	}
}

func (p *Pawn) calculatePinnedOptions(position int) {
	if p.PinnedToKing {
		p.Protecting = make(map[int]any)
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

func (p *Pawn) deleteOptions(board map[int]any) {
	var toRemove []int
	for option, _ := range p.Options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(p.Options, toDelete)
	}
}

func (p *Pawn) addAttackedBy(attackedPiece any, position int) {
	switch CheckPieceKindFromAny(attackedPiece) {
	case PieceKindPawn:
		pawn := attackedPiece.(*Pawn)
		pawn.AttackedBy[position] = p
	case PieceKindKnight:
		knight := attackedPiece.(*Knight)
		knight.AttackedBy[position] = p
	case PieceKindBishop:
		bishop := attackedPiece.(*Bishop)
		bishop.AttackedBy[position] = p
	case PieceKindRook:
		rook := attackedPiece.(*Rook)
		rook.AttackedBy[position] = p
	case PieceKindQueen:
		queen := attackedPiece.(*Queen)
		queen.AttackedBy[position] = p
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating attacked by bishop")
	}
}
