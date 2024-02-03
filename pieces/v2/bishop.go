package v2

import "log"

type Bishop struct {
	*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
}

func CalculateBishopMoves(bishop *Bishop, whiteBoard map[int]any, blackBoard map[int]any, position int) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if bishop.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position % 8
	colPos := position / 8

	forbidden := calculateBishopOptions(bishop, position, rowPos, -1, 9, 1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateBishopOptions(bishop, position, rowPos, -1, 7, -1, colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateBishopOptions(bishop, position, rowPos, 1, 7, 1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateBishopOptions(bishop, position, rowPos, 1, 9, -1, 7-colPos, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	bishop.calculatePinnedOptions(position)

	return forbiddenSquares
}

func calculateBishopOptions(bishop *Bishop, position int, rowPos int, down int, sideways int, beyondBoardMultiplier int, until int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
	for up := 1; up <= until; up++ {
		newPosition := position + down*up*sideways
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if beyondBoardMultiplier*(rowPos-newPosition%8) < 0 {
			return forbiddenSquares
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			bishop.Options[newPosition] = value
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = bishop
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
		bishop.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (b *Bishop) calculatePinnedOptions(position int) {
	if b.PinnedToKing {
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
