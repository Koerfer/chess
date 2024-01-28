package enginev2

import "chess/pieces"

func (e *Engine) createNewBoardState(option *Option) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	myBoard := make(map[int]*pieces.Piece)
	opponentBoard := make(map[int]*pieces.Piece)

	white := option.Piece.White

	switch white {
	case true:
		for k, v := range e.whiteBoard {
			myBoard[k] = &pieces.Piece{
				Kind:             v.Kind,
				White:            v.White,
				LastPosition:     v.LastPosition,
				Options:          v.Options,
				EnPassantOptions: v.EnPassantOptions,
				HasBeenMoved:     v.HasBeenMoved,
				Checked:          v.Checked,
				CheckingPieces:   v.CheckingPieces,
				PinnedToKing:     v.PinnedToKing,
				PinnedByPosition: v.PinnedByPosition,
				PinnedByPiece:    v.PinnedByPiece,
			}
		}
		for k, v := range e.blackBoard {
			opponentBoard[k] = &pieces.Piece{
				Kind:             v.Kind,
				White:            v.White,
				LastPosition:     v.LastPosition,
				Options:          v.Options,
				EnPassantOptions: v.EnPassantOptions,
				HasBeenMoved:     v.HasBeenMoved,
				Checked:          v.Checked,
				CheckingPieces:   v.CheckingPieces,
				PinnedToKing:     v.PinnedToKing,
				PinnedByPosition: v.PinnedByPosition,
				PinnedByPiece:    v.PinnedByPiece,
			}
		}
	case false:
		for k, v := range e.blackBoard {
			myBoard[k] = &pieces.Piece{
				Kind:             v.Kind,
				White:            v.White,
				LastPosition:     v.LastPosition,
				Options:          v.Options,
				EnPassantOptions: v.EnPassantOptions,
				HasBeenMoved:     v.HasBeenMoved,
				Checked:          v.Checked,
				CheckingPieces:   v.CheckingPieces,
				PinnedToKing:     v.PinnedToKing,
				PinnedByPosition: v.PinnedByPosition,
				PinnedByPiece:    v.PinnedByPiece,
			}
		}
		for k, v := range e.whiteBoard {
			opponentBoard[k] = &pieces.Piece{
				Kind:             v.Kind,
				White:            v.White,
				LastPosition:     v.LastPosition,
				Options:          v.Options,
				EnPassantOptions: v.EnPassantOptions,
				HasBeenMoved:     v.HasBeenMoved,
				Checked:          v.Checked,
				CheckingPieces:   v.CheckingPieces,
				PinnedToKing:     v.PinnedToKing,
				PinnedByPosition: v.PinnedByPosition,
				PinnedByPiece:    v.PinnedByPiece,
			}
		}
	}

	pieceCopy := myBoard[option.Piece.LastPosition]

	if option.EnPassant != 0 {
		return e.enPassant(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, option.EnPassant, pieceCopy, white)
	}

	return e.normal(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, pieceCopy, white)
}

func (e *Engine) enPassant(myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece, position int, moveTo int, take int, selectedPiece *pieces.Piece, whitesTurn bool) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	switch whitesTurn {
	case true:
		delete(opponentBoard, take)
	case false:
		delete(opponentBoard, take)
	}

	myBoard[moveTo] = selectedPiece
	delete(myBoard, position)
	selectedPiece = nil

	return calculateAllPositions(myBoard, opponentBoard)
}

func (e *Engine) normal(myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece, position int, option int, selectedPiece *pieces.Piece, whitesTurn bool) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	if selectedPiece.Kind == pieces.Pawn {
		end := 7
		if selectedPiece.White == true {
			end = 0
		}
		if position/8 == end {
			selectedPiece.Kind = pieces.Queen // todo: add convert to better Piece logic
		}
	}

	switch whitesTurn {
	case true:
		delete(opponentBoard, option)
	case false:
		delete(opponentBoard, option)
	}

	if selectedPiece.Kind == pieces.King && !selectedPiece.HasBeenMoved {
		castled, myBoard, opponentBoard := e.castle(option, myBoard, opponentBoard, selectedPiece)
		if castled {
			return myBoard, opponentBoard
		}
	}

	selectedPiece.HasBeenMoved = true

	myBoard[option] = selectedPiece
	delete(myBoard, selectedPiece.LastPosition)
	selectedPiece = nil

	whitesTurn = !whitesTurn
	return calculateAllPositions(myBoard, opponentBoard)
}

func (e *Engine) castle(option int, myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece, selectedPiece *pieces.Piece) (bool, map[int]*pieces.Piece, map[int]*pieces.Piece) {
	switch option {
	case 2:
		selectedPiece.HasBeenMoved = true
		myBoard[option] = selectedPiece
		myBoard[3] = myBoard[0]
		delete(myBoard, selectedPiece.LastPosition)
		delete(myBoard, 0)
		selectedPiece = nil
		myBoard, opponentBoard := calculateAllPositions(myBoard, opponentBoard)
		return true, myBoard, opponentBoard
	case 6:
		selectedPiece.HasBeenMoved = true
		myBoard[option] = selectedPiece
		myBoard[5] = myBoard[7]
		delete(myBoard, selectedPiece.LastPosition)
		delete(myBoard, 7)
		selectedPiece = nil
		myBoard, opponentBoard := calculateAllPositions(myBoard, opponentBoard)
		return true, myBoard, opponentBoard
	case 58:
		selectedPiece.HasBeenMoved = true
		myBoard[option] = selectedPiece
		myBoard[59] = myBoard[56]
		delete(myBoard, selectedPiece.LastPosition)
		delete(myBoard, 56)
		selectedPiece = nil
		myBoard, opponentBoard := calculateAllPositions(myBoard, opponentBoard)
		return true, myBoard, opponentBoard
	case 62:
		selectedPiece.HasBeenMoved = true
		myBoard[option] = selectedPiece
		myBoard[61] = myBoard[63]
		delete(myBoard, selectedPiece.LastPosition)
		delete(myBoard, 63)
		selectedPiece = nil
		myBoard, opponentBoard := calculateAllPositions(myBoard, opponentBoard)
		return true, myBoard, opponentBoard
	}

	return false, nil, nil
}

func calculateAllPositions(myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range myBoard {
		piece.PinnedToKing = false
	}
	for _, piece := range opponentBoard {
		piece.PinnedToKing = false
	}

	var checkingPieces map[int]*pieces.Piece
	var kingPosition int

	for position, piece := range opponentBoard {
		forbiddenCaptures, tmpCheck := piece.CalculateOptions(myBoard, opponentBoard, position, nil, false)
		check = check || tmpCheck
		for forbidden := range forbiddenCaptures {
			forbiddenSquares[forbidden] = struct{}{}
		}
		if piece.Kind != pieces.Pawn {
			for forbidden := range piece.Options {
				forbiddenSquares[forbidden] = struct{}{}
			}
		}
		if piece.Kind == pieces.King {
			piece.CheckingPieces = make(map[int]*pieces.Piece)
		}
	}

	for position, piece := range myBoard {
		piece.CalculateOptions(myBoard, opponentBoard, position, forbiddenSquares, true)
		if check && piece.Kind == pieces.King {
			checkingPieces = piece.CheckingPieces
			kingPosition = position
		}
	}
	if check {
		for _, piece := range myBoard {
			if piece.Kind != pieces.King {
				piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
			}
		}
	}

	return myBoard, opponentBoard
}
