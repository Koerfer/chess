package v2

type Knight struct {
	*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    any
}

func CalculateKnightMoves(knight *Knight, whiteBoard map[int]any, blackBoard map[int]any, position int, fixLastPosition bool) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})
	if fixLastPosition {
		knight.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if knight.White == false {
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
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 { // down 1 ok
			newPosition := position + left*2 - up
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if right+2 <= 7 {
		if down-1 >= 0 {
			newPosition := position - left*2 + up
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 {
			newPosition := position - left*2 - up
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down+2 <= 7 {
		if right-1 >= 0 {
			newPosition := position - up*2 + left
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position - up*2 - left
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down-2 >= 0 {
		if right-1 >= 0 {
			newPosition := position + up*2 + left
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position + up*2 - left
			calculateKnightOptions(knight, position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	knightDelete(knight, myBoard)
	if knight.PinnedToKing {
		knight.Options = make(map[int]struct{})
	}

	return forbiddenSquares
}

func calculateKnightOptions(knight *Knight, position int, newPosition int, opponentBoard map[int]any, forbiddenSquares map[int]struct{}) {
	opponent, ok := opponentBoard[newPosition]
	if ok && CheckPieceKindFromAny(opponent) == PieceKindKing {
		p := opponent.(*King)
		p.Checked = true
		p.CheckingPieces[position] = knight
	}
	knight.Options[newPosition] = value
	forbiddenSquares[newPosition] = value
}

func knightDelete(knight *Knight, board map[int]any) {
	var toRemove []int
	for option, _ := range knight.Options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(knight.Options, toDelete)
	}
}
