package v2

type Bishop struct {
	Value          int
	EvaluatedValue int
	White          bool
	LastPosition   int
	Options        map[int]struct{}
	Protecting     map[int]PieceInterface
	AttackedBy     map[int]PieceInterface
	ProtectedBy    map[int]PieceInterface

	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    PieceInterface
}

func (b *Bishop) GetValue() int {
	return b.Value
}
func (b *Bishop) GetEvaluatedValue() int {
	return b.EvaluatedValue
}
func (b *Bishop) GetWhite() bool {
	return b.White
}
func (b *Bishop) GetLastPosition() int {
	return b.LastPosition
}
func (b *Bishop) GetOptions() map[int]struct{} {
	return b.Options
}
func (b *Bishop) GetProtecting() map[int]PieceInterface {
	return b.Protecting
}
func (b *Bishop) GetAttackedBy() map[int]PieceInterface {
	return b.AttackedBy
}
func (b *Bishop) GetProtectedBy() map[int]PieceInterface {
	return b.ProtectedBy
}
func (b *Bishop) SetValue(value int) {
	b.Value = value
}
func (b *Bishop) SetEvaluatedValue(evaluatedValue int) {
	b.EvaluatedValue = evaluatedValue
}
func (b *Bishop) SetWhite(white bool) {
	b.White = white
}
func (b *Bishop) SetLastPosition(lastPosition int) {
	b.LastPosition = lastPosition
}
func (b *Bishop) SetOptions(options map[int]struct{}) {
	b.Options = options
}
func (b *Bishop) SetProtecting(protecting map[int]PieceInterface) {
	b.Protecting = protecting
}
func (b *Bishop) SetAttackedBy(attackedBy map[int]PieceInterface) {
	b.AttackedBy = attackedBy
}
func (b *Bishop) SetProtectedBy(protectedBy map[int]PieceInterface) {
	b.ProtectedBy = protectedBy
}

func (b *Bishop) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	if fixLastPosition {
		b.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if b.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	forbidden := b.calculateOptions(position, rowPos, -1, 9, 1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = b.calculateOptions(position, rowPos, -1, 7, -1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = b.calculateOptions(position, rowPos, 1, 7, 1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = b.calculateOptions(position, rowPos, 1, 9, -1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	b.calculatePinnedOptions(position)

	return forbiddenSquares
}

func (b *Bishop) calculateOptions(position int, rowPos int, down int, sideways int, beyondBoardMultiplier int, until int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
	for up := 1; up <= until; up++ {
		newPosition := position + down*up*sideways
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if beyondBoardMultiplier*(rowPos-newPosition%8) < 0 {
			return forbiddenSquares
		}
		if protectedPiece, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			b.Protecting[newPosition] = protectedPiece
			addProtectedBy(b, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			b.Options[newPosition] = value
			if !b.PinnedToKing {
				addAttackedBy(b, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = b
				for upKing := up; upKing <= 8; upKing++ {
					newPosition := position + down*upKing*sideways
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if beyondBoardMultiplier*(rowPos-newPosition%8) < 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			} else {
				for upPin := up + 1; upPin <= 8; upPin++ {
					newPositionPin := position + down*upPin*sideways
					if newPositionPin < 0 {
						break
					}
					if beyondBoardMultiplier*(rowPos-newPositionPin%8) < 0 {
						break
					}
					piece, ok := opponentBoard[newPositionPin]
					if ok {
						if CheckPieceKindFromAny(piece) == PieceKindKing {
							switch CheckPieceKindFromAny(opponent) {
							case PieceKindPawn:
								p := opponent.(*Pawn)
								p.PinnedToKing = true
								p.PinnedByPosition = position
							case PieceKindKnight:
								p := opponent.(*Knight)
								p.PinnedToKing = true
								p.PinnedByPosition = position
							case PieceKindBishop:
								p := opponent.(*Bishop)
								p.PinnedToKing = true
								p.PinnedByPosition = position
							case PieceKindRook:
								p := opponent.(*Rook)
								p.PinnedToKing = true
								p.PinnedByPosition = position
							case PieceKindQueen:
								p := opponent.(*Queen)
								p.PinnedToKing = true
								p.PinnedByPosition = position
							case PieceKindKing:
								// do nothing
							case PieceKindInvalid:
								panic("invalid piece kind during queen pinning")
							}
						}
						return forbiddenSquares
					}
				}
			}

			return forbiddenSquares
		}

		forbiddenSquares[newPosition] = value
		b.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (b *Bishop) calculatePinnedOptions(position int) {
	if b.PinnedToKing {
		b.Protecting = make(map[int]PieceInterface)
		for option := range b.Options {
			if option == b.PinnedByPosition {
				continue
			}
			if b.PinnedByPosition%8 == position%8 { // same column
				if b.PinnedByPosition/8 < position/8 {
					if option%8 == position%8 && option < position && option > b.PinnedByPosition {
						continue
					}
				}
				if b.PinnedByPosition/8 > position/8 {
					if option%8 == position%8 && option > position && option < b.PinnedByPosition {
						continue
					}
				}
			} else if b.PinnedByPosition/8 == position/8 { // same row
				if b.PinnedByPosition%8 < position%8 {
					if option/8 == position/8 && option < position && option > b.PinnedByPosition {
						continue
					}
				}
				if b.PinnedByPosition%8 > position%8 {
					if option/8 == position/8 && option > position && option < b.PinnedByPosition {
						continue
					}
				}
			} else if b.PinnedByPosition%9 == position%9 { // same left->right column
				if b.PinnedByPosition < position {
					if option%9 == position%9 && option < position && option > b.PinnedByPosition {
						continue
					}
				}
				if b.PinnedByPosition > position {
					if option%9 == position%9 && option > position && option < b.PinnedByPosition {
						continue
					}
				}
			} else if b.PinnedByPosition%7 == position%7 { // same right->left column
				if b.PinnedByPosition < position {
					if option%7 == position%7 && option < position && option > b.PinnedByPosition {
						continue
					}
				}
				if b.PinnedByPosition > position {
					if option%7 == position%7 && option > position && option < b.PinnedByPosition {
						continue
					}
				}
			}
			delete(b.Options, option)
		}
	}
}

func (b *Bishop) Copy(deep bool) PieceInterface {
	if b == nil {
		return nil
	}
	copyCat := &Bishop{
		Value:            b.Value,
		EvaluatedValue:   b.EvaluatedValue,
		White:            b.White,
		LastPosition:     b.LastPosition,
		Options:          b.Options,
		PinnedToKing:     b.PinnedToKing,
		PinnedByPosition: b.PinnedByPosition,
	}
	if deep {
		copyCat.PinnedByPiece = b.PinnedByPiece.Copy(false)
		copyCat.Protecting, copyCat.ProtectedBy, copyCat.AttackedBy = copyProtectingAndAttacking(b.Protecting, b.ProtectedBy, b.AttackedBy)
	}
	return copyCat
}
