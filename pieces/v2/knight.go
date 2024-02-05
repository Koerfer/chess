package v2

import "log"

type Knight struct {
	*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
	AttackedBy       map[int]any
}

func (k *Knight) CalculateMoves(whiteBoard map[int]any, blackBoard map[int]any, position int, fixLastPosition bool) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})
	if fixLastPosition {
		k.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if k.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	right := position % 8
	down := position / 8
	up := -8
	left := -1
	if right-2 >= 0 { // left 2 ok
		if down-1 >= 0 { // up 1 ok
			newPosition := position + left*2 + up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 { // down 1 ok
			newPosition := position + left*2 - up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if right+2 <= 7 {
		if down-1 >= 0 {
			newPosition := position - left*2 + up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 {
			newPosition := position - left*2 - up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down+2 <= 7 {
		if right-1 >= 0 {
			newPosition := position - up*2 + left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position - up*2 - left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down-2 >= 0 {
		if right-1 >= 0 {
			newPosition := position + up*2 + left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position + up*2 - left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	k.deleteOptions(myBoard)
	if k.PinnedToKing {
		k.Options = make(map[int]struct{})
		k.Protecting = make(map[int]any)
	}

	return forbiddenSquares
}

func (k *Knight) calculateOptions(position int, newPosition int, opponentBoard map[int]any, forbiddenSquares map[int]struct{}) {
	opponent, ok := opponentBoard[newPosition]
	if ok && !k.PinnedToKing {
		k.addAttackedBy(opponent, position)
	}
	if ok && CheckPieceKindFromAny(opponent) == PieceKindKing {
		p := opponent.(*King)
		p.Checked = true
		p.CheckingPieces[position] = k
	}
	k.Options[newPosition] = value
	forbiddenSquares[newPosition] = value
}

func (k *Knight) deleteOptions(board map[int]any) {
	var toRemove []int
	for option := range k.Options {
		if protectedPiece, ok := board[option]; ok {
			k.Protecting[option] = protectedPiece
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(k.Options, toDelete)
	}
}

func (k *Knight) addAttackedBy(attackedPiece any, position int) {
	switch CheckPieceKindFromAny(attackedPiece) {
	case PieceKindPawn:
		pawn := attackedPiece.(*Pawn)
		pawn.AttackedBy[position] = k
	case PieceKindKnight:
		knight := attackedPiece.(*Knight)
		knight.AttackedBy[position] = k
	case PieceKindBishop:
		bishop := attackedPiece.(*Bishop)
		bishop.AttackedBy[position] = k
	case PieceKindRook:
		rook := attackedPiece.(*Rook)
		rook.AttackedBy[position] = k
	case PieceKindQueen:
		queen := attackedPiece.(*Queen)
		queen.AttackedBy[position] = k
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating attacked by knight")
	}
}
