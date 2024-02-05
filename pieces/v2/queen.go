package v2

import "log"

type Queen struct {
	*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
	AttackedBy       map[int]any
}

func (q *Queen) CalculateMoves(whiteBoard map[int]any, blackBoard map[int]any, position int, fixLastPosition bool) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})
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

func (q *Queen) calculateDiagonalOptions(position int, rowPos int, down int, sideways int, beyondBoardMultiplier int, until int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
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
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				q.addAttackedBy(opponent, position)
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
								log.Fatal("invalid piece kind during queen pinning")
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

func (q *Queen) calculateHorizontalOptions(position int, rowPos int, right int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
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
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				q.addAttackedBy(opponent, position)
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
								log.Fatal("invalid piece kind during queen pinning")
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

func (q *Queen) calculateVerticalOptions(position int, colPos int, down int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
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
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			q.Options[newPosition] = value
			if !q.PinnedToKing {
				q.addAttackedBy(opponent, position)
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
								log.Fatal("invalid piece kind during queen pinning")
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
		q.Protecting = make(map[int]any)
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

func (q *Queen) addAttackedBy(attackedPiece any, position int) {
	switch CheckPieceKindFromAny(attackedPiece) {
	case PieceKindPawn:
		pawn := attackedPiece.(*Pawn)
		pawn.AttackedBy[position] = q
	case PieceKindKnight:
		knight := attackedPiece.(*Knight)
		knight.AttackedBy[position] = q
	case PieceKindBishop:
		bishop := attackedPiece.(*Bishop)
		bishop.AttackedBy[position] = q
	case PieceKindRook:
		rook := attackedPiece.(*Rook)
		rook.AttackedBy[position] = q
	case PieceKindQueen:
		queen := attackedPiece.(*Queen)
		queen.AttackedBy[position] = q
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating attacked by queen")
	}
}
