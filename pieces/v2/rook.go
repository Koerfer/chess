package v2

import "log"

type Rook struct {
	*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
	HasBeenMoved     bool
}

func CalculateRookMoves(rook *Rook, whiteBoard map[int]any, blackBoard map[int]any, position int) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if rook.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	rowPos := position / 8
	colPos := position % 8

	forbidden := calculateRookHorizontalOptions(rook, position, rowPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateRookHorizontalOptions(rook, position, rowPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateRookVerticalOptions(rook, position, colPos, 1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	forbidden = calculateRookVerticalOptions(rook, position, colPos, -1, forbiddenSquares, myBoard, opponentBoard)
	forbiddenSquares = mergeMaps(forbiddenSquares, forbidden)

	rook.calculatePinnedOptions(position)

	return forbiddenSquares
}

func calculateRookHorizontalOptions(rook *Rook, position int, rowPos int, right int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
	for left := 1; left <= 8; left++ {
		newPosition := position + right*left
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if rowPos-newPosition/8 != 0 {
			return forbiddenSquares
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			rook.Options[newPosition] = value
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = rook
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
		rook.Options[newPosition] = value
	}

	return forbiddenSquares
}

func calculateRookVerticalOptions(rook *Rook, position int, colPos int, down int, forbiddenSquares map[int]struct{}, myBoard map[int]any, opponentBoard map[int]any) map[int]struct{} {
	for up := 1; up <= 8; up++ {
		newPosition := position + down*up*8
		if newPosition < 0 || newPosition > 63 {
			return forbiddenSquares
		}
		if colPos-newPosition%8 != 0 {
			return forbiddenSquares
		}
		if _, ok := myBoard[newPosition]; ok {
			forbiddenSquares[newPosition] = value
			return forbiddenSquares
		}
		opponent, ok := opponentBoard[newPosition]
		if ok {
			rook.Options[newPosition] = value
			if CheckPieceKindFromAny(opponent) == PieceKindKing {
				p := opponent.(*King)
				p.Checked = true
				p.CheckingPieces[position] = rook
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
		rook.Options[newPosition] = value
	}

	return forbiddenSquares
}

func (r *Rook) calculatePinnedOptions(position int) {
	if r.PinnedToKing {
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