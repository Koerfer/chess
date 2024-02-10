package v2

type Queen struct {
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

func (q *Queen) GetValue() int {
	return q.Value
}
func (q *Queen) GetEvaluatedValue() int {
	return q.EvaluatedValue
}
func (q *Queen) GetWhite() bool {
	return q.White
}
func (q *Queen) GetLastPosition() int {
	return q.LastPosition
}
func (q *Queen) GetOptions() map[int]struct{} {
	return q.Options
}
func (q *Queen) GetProtecting() map[int]PieceInterface {
	return q.Protecting
}
func (q *Queen) GetAttackedBy() map[int]PieceInterface {
	return q.AttackedBy
}
func (q *Queen) GetProtectedBy() map[int]PieceInterface {
	return q.ProtectedBy
}
func (q *Queen) SetValue(value int) {
	q.Value = value
}
func (q *Queen) SetEvaluatedValue(evaluatedValue int) {
	q.EvaluatedValue = evaluatedValue
}
func (q *Queen) SetWhite(white bool) {
	q.White = white
}
func (q *Queen) SetLastPosition(lastPosition int) {
	q.LastPosition = lastPosition
}
func (q *Queen) SetOptions(options map[int]struct{}) {
	q.Options = options
}
func (q *Queen) SetProtecting(protecting map[int]PieceInterface) {
	q.Protecting = protecting
}
func (q *Queen) SetAttackedBy(attackedBy map[int]PieceInterface) {
	q.AttackedBy = attackedBy
}
func (q *Queen) SetProtectedBy(protectedBy map[int]PieceInterface) {
	q.ProtectedBy = protectedBy
}

func (q *Queen) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	if fixLastPosition {
		q.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if q.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	forbidden := q.calculateDiagonalOptions(position, rowPos, -1, 9, 1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateDiagonalOptions(position, rowPos, -1, 7, -1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateDiagonalOptions(position, rowPos, 1, 7, 1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateDiagonalOptions(position, rowPos, 1, 9, -1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	rowPos = position / 8
	colPos = position % 8

	forbidden = q.calculateHorizontalOptions(position, rowPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateHorizontalOptions(position, rowPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateVerticalOptions(position, colPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = q.calculateVerticalOptions(position, colPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	q.calculatePinnedOptions(position)

	return forbiddenSquares
}

func (q *Queen) calculateDiagonalOptions(position int, rowPos int, down int, sideways int, beyondBoardMultiplier int, until int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
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
			q.Protecting[newPosition] = protectedPiece
			addProtectedBy(q, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				addAttackedBy(q, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = q
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
		q.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (q *Queen) calculateHorizontalOptions(position int, rowPos int, right int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
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
			q.Protecting[newPosition] = protectedPiece
			addProtectedBy(q, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				addAttackedBy(q, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = q
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
					if rowPos-newPositionPin/8 < 0 {
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
		q.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (q *Queen) calculateVerticalOptions(position int, colPos int, down int, forbiddenSquares map[int]struct{}, myBoard map[int]PieceInterface, opponentBoard map[int]PieceInterface) map[int]struct{} {
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
			q.Protecting[newPosition] = protectedPiece
			addProtectedBy(q, protectedPiece, position)
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				addAttackedBy(q, opponent, position)
			}
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = q
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
		q.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (q *Queen) calculatePinnedOptions(position int) {
	if q.PinnedToKing {
		q.Protecting = make(map[int]PieceInterface)
		for option := range q.Options {
			if option == q.PinnedByPosition {
				continue
			}
			if q.PinnedByPosition%8 == position%8 { // same column
				if q.PinnedByPosition/8 < position/8 {
					if option%8 == position%8 && option < position && option > q.PinnedByPosition {
						continue
					}
				}
				if q.PinnedByPosition/8 > position/8 {
					if option%8 == position%8 && option > position && option < q.PinnedByPosition {
						continue
					}
				}
			} else if q.PinnedByPosition/8 == position/8 { // same row
				if q.PinnedByPosition%8 < position%8 {
					if option/8 == position/8 && option < position && option > q.PinnedByPosition {
						continue
					}
				}
				if q.PinnedByPosition%8 > position%8 {
					if option/8 == position/8 && option > position && option < q.PinnedByPosition {
						continue
					}
				}
			} else if q.PinnedByPosition%9 == position%9 { // same left->right column
				if q.PinnedByPosition < position {
					if option%9 == position%9 && option < position && option > q.PinnedByPosition {
						continue
					}
				}
				if q.PinnedByPosition > position {
					if option%9 == position%9 && option > position && option < q.PinnedByPosition {
						continue
					}
				}
			} else if q.PinnedByPosition%7 == position%7 { // same right->left column
				if q.PinnedByPosition < position {
					if option%7 == position%7 && option < position && option > q.PinnedByPosition {
						continue
					}
				}
				if q.PinnedByPosition > position {
					if option%7 == position%7 && option > position && option < q.PinnedByPosition {
						continue
					}
				}
			}
			delete(q.Options, option)
		}
	}
}

func (q *Queen) Copy(deep bool) PieceInterface {
	if q == nil {
		return nil
	}
	copyCat := &Queen{
		Value:            q.Value,
		EvaluatedValue:   q.EvaluatedValue,
		White:            q.White,
		LastPosition:     q.LastPosition,
		Options:          q.Options,
		PinnedToKing:     q.PinnedToKing,
		PinnedByPosition: q.PinnedByPosition,
	}
	if deep {
		copyCat.PinnedByPiece = q.PinnedByPiece.Copy(false)
		copyCat.Protecting, copyCat.ProtectedBy, copyCat.AttackedBy = copyProtectingAndAttacking(q.Protecting, q.ProtectedBy, q.AttackedBy)
	}
	return copyCat
}
