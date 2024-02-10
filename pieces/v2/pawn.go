package v2

type Pawn struct {
	Value          int
	EvaluatedValue int
	White          bool
	LastPosition   int
	Options        map[int]struct{}
	Protecting     map[int]PieceInterface
	AttackedBy     map[int]PieceInterface
	ProtectedBy    map[int]PieceInterface

	EnPassantOptions map[int]int
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    PieceInterface
}

func (p *Pawn) GetValue() int {
	return p.Value
}
func (p *Pawn) GetEvaluatedValue() int {
	return p.EvaluatedValue
}
func (p *Pawn) GetWhite() bool {
	return p.White
}
func (p *Pawn) GetLastPosition() int {
	return p.LastPosition
}
func (p *Pawn) GetOptions() map[int]struct{} {
	return p.Options
}
func (p *Pawn) GetProtecting() map[int]PieceInterface {
	return p.Protecting
}
func (p *Pawn) GetAttackedBy() map[int]PieceInterface {
	return p.AttackedBy
}
func (p *Pawn) GetProtectedBy() map[int]PieceInterface {
	return p.ProtectedBy
}
func (p *Pawn) SetValue(value int) {
	p.Value = value
}
func (p *Pawn) SetEvaluatedValue(evaluatedValue int) {
	p.EvaluatedValue = evaluatedValue
}
func (p *Pawn) SetWhite(white bool) {
	p.White = white
}
func (p *Pawn) SetLastPosition(lastPosition int) {
	p.LastPosition = lastPosition
}
func (p *Pawn) SetOptions(options map[int]struct{}) {
	p.Options = options
}
func (p *Pawn) SetProtecting(protecting map[int]PieceInterface) {
	p.Protecting = protecting
}
func (p *Pawn) SetAttackedBy(attackedBy map[int]PieceInterface) {
	p.AttackedBy = attackedBy
}
func (p *Pawn) SetProtectedBy(protectedBy map[int]PieceInterface) {
	p.ProtectedBy = protectedBy
}

func (p *Pawn) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
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

func (p *Pawn) calculateCaptureOption(position int, endPosition int, captureOption int, direction int, offsetMultiplier int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) {
	forbiddenSquares[captureOption] = value
	if protectedPiece, ok := myBoard[captureOption]; ok {
		forbiddenSquares[captureOption] = value
		p.Protecting[captureOption] = protectedPiece
		addProtectedBy(p, protectedPiece, position)
	}
	if capturePiece, ok := opponentBoard[captureOption]; ok { // if black Piece up right
		p.Options[captureOption] = value // add capture move
		if !p.PinnedToKing {
			addAttackedBy(p, capturePiece, position)
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
		p.Protecting = make(map[int]PieceInterface)
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

func (p *Pawn) deleteOptions(board map[int]PieceInterface) {
	var toRemove []int
	for option := range p.Options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(p.Options, toDelete)
	}
}

func (p *Pawn) Copy(deep bool) PieceInterface {
	if p == nil {
		return nil
	}
	copyCat := &Pawn{
		Value:            p.Value,
		EvaluatedValue:   p.EvaluatedValue,
		White:            p.White,
		LastPosition:     p.LastPosition,
		Options:          p.Options,
		EnPassantOptions: p.EnPassantOptions,
		PinnedToKing:     p.PinnedToKing,
		PinnedByPosition: p.PinnedByPosition,
	}
	if deep {
		copyCat.PinnedByPiece = p.PinnedByPiece.Copy(false)
		copyCat.Protecting, copyCat.ProtectedBy, copyCat.AttackedBy = copyProtectingAndAttacking(p.Protecting, p.ProtectedBy, p.AttackedBy)
	}
	return copyCat
}
