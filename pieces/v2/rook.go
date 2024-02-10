package v2

type Rook struct {
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

	HasBeenMoved bool
}

func (r *Rook) GetValue() int {
	return r.Value
}
func (r *Rook) GetEvaluatedValue() int {
	return r.EvaluatedValue
}
func (r *Rook) GetWhite() bool {
	return r.White
}
func (r *Rook) GetLastPosition() int {
	return r.LastPosition
}
func (r *Rook) GetOptions() map[int]struct{} {
	return r.Options
}
func (r *Rook) GetProtecting() map[int]PieceInterface {
	return r.Protecting
}
func (r *Rook) GetAttackedBy() map[int]PieceInterface {
	return r.AttackedBy
}
func (r *Rook) GetProtectedBy() map[int]PieceInterface {
	return r.ProtectedBy
}
func (r *Rook) SetValue(value int) {
	r.Value = value
}
func (r *Rook) SetEvaluatedValue(evaluatedValue int) {
	r.EvaluatedValue = evaluatedValue
}
func (r *Rook) SetWhite(white bool) {
	r.White = white
}
func (r *Rook) SetLastPosition(lastPosition int) {
	r.LastPosition = lastPosition
}
func (r *Rook) SetOptions(options map[int]struct{}) {
	r.Options = options
}
func (r *Rook) SetProtecting(protecting map[int]PieceInterface) {
	r.Protecting = protecting
}
func (r *Rook) SetAttackedBy(attackedBy map[int]PieceInterface) {
	r.AttackedBy = attackedBy
}
func (r *Rook) SetProtectedBy(protectedBy map[int]PieceInterface) {
	r.ProtectedBy = protectedBy
}

func (r *Rook) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	if fixLastPosition {
		r.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if r.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position / 8
	colPos := position % 8

	forbidden := r.calculateHorizontalOptions(position, rowPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = r.calculateHorizontalOptions(position, rowPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = r.calculateVerticalOptions(position, colPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = r.calculateVerticalOptions(position, colPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	r.calculatePinnedOptions(position)

	return forbiddenSquares
}

func (r *Rook) calculateHorizontalOptions(position int, rowPos int, right int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
	for left := 1; left <= 8; left++ {
		newPosition := position + right*left
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if rowPos-newPosition/8 != 0 {
			return forbiddenSquares
		}
		if protectedPiece, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			r.Protecting[newPosition] = protectedPiece
			addProtectedBy(r, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			r.Options[newPosition] = value
			if !r.PinnedToKing {
				addAttackedBy(r, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = r
				for leftKing := left; leftKing <= 8; leftKing++ {
					newPosition := position + right*leftKing
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if rowPos-newPosition/8 != 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			} else {
				for leftPin := left + 1; leftPin <= 8; leftPin++ {
					newPositionPin := position + right*leftPin
					if newPositionPin < 0 {
						break
					}
					if rowPos-newPositionPin/8 != 0 {
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
						break
					}
				}
			}
			return forbiddenSquares
		}

		forbiddenSquares[newPosition] = value
		r.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (r *Rook) calculateVerticalOptions(position int, colPos int, down int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
	for up := 1; up <= 8; up++ {
		newPosition := position + down*up*8
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if colPos-newPosition%8 != 0 {
			return forbiddenSquares
		}
		if protectedPiece, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			r.Protecting[newPosition] = protectedPiece
			addProtectedBy(r, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			r.Options[newPosition] = value
			if !r.PinnedToKing {
				addAttackedBy(r, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = r
				for upKing := up; upKing <= 8; upKing++ {
					newPosition := position + down*upKing*8
					if newPosition < 0 || newPosition > 63 {
						break
					}
					if colPos-newPosition%8 != 0 {
						break
					}
					forbiddenSquares[newPosition] = value
				}
			} else {
				for upPin := up + 1; upPin <= 8; upPin++ {
					newPositionPin := position + down*upPin*8
					if newPositionPin < 0 {
						break
					}
					if colPos-newPositionPin%8 != 0 {
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
						break
					}
				}
			}
			return forbiddenSquares
		}

		forbiddenSquares[newPosition] = value
		r.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (r *Rook) calculatePinnedOptions(position int) {
	if r.PinnedToKing {
		r.Protecting = make(map[int]PieceInterface)
		for option := range r.Options {
			if option == r.PinnedByPosition {
				continue
			}
			if r.PinnedByPosition%8 == position%8 { // same column
				if r.PinnedByPosition/8 < position/8 {
					if option%8 == position%8 && option < position && option > r.PinnedByPosition {
						continue
					}
				}
				if r.PinnedByPosition/8 > position/8 {
					if option%8 == position%8 && option > position && option < r.PinnedByPosition {
						continue
					}
				}
			} else if r.PinnedByPosition/8 == position/8 { // same row
				if r.PinnedByPosition%8 < position%8 {
					if option/8 == position/8 && option < position && option > r.PinnedByPosition {
						continue
					}
				}
				if r.PinnedByPosition%8 > position%8 {
					if option/8 == position/8 && option > position && option < r.PinnedByPosition {
						continue
					}
				}
			} else if r.PinnedByPosition%9 == position%9 { // same left->right column
				if r.PinnedByPosition < position {
					if option%9 == position%9 && option < position && option > r.PinnedByPosition {
						continue
					}
				}
				if r.PinnedByPosition > position {
					if option%9 == position%9 && option > position && option < r.PinnedByPosition {
						continue
					}
				}
			} else if r.PinnedByPosition%7 == position%7 { // same right->left column
				if r.PinnedByPosition < position {
					if option%7 == position%7 && option < position && option > r.PinnedByPosition {
						continue
					}
				}
				if r.PinnedByPosition > position {
					if option%7 == position%7 && option > position && option < r.PinnedByPosition {
						continue
					}
				}
			}
			delete(r.Options, option)
		}
	}
}

func (r *Rook) Copy(deep bool) PieceInterface {
	if r == nil {
		return nil
	}
	copyCat := &Rook{
		Value:            r.Value,
		EvaluatedValue:   r.EvaluatedValue,
		White:            r.White,
		LastPosition:     r.LastPosition,
		Options:          r.Options,
		PinnedToKing:     r.PinnedToKing,
		PinnedByPosition: r.PinnedByPosition,
		HasBeenMoved:     r.HasBeenMoved,
	}
	if deep {
		copyCat.PinnedByPiece = r.PinnedByPiece.Copy(false)
		copyCat.Protecting, copyCat.ProtectedBy, copyCat.AttackedBy = copyProtectingAndAttacking(r.Protecting, r.ProtectedBy, r.AttackedBy)
	}
	return copyCat
}
