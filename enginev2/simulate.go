package enginev2

import "chess/pieces"

func (e *Engine) createNewBoardState(option *Option) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	myBoard := make(map[int]*pieces.Piece)
	opponentBoard := make(map[int]*pieces.Piece)

	white := !option.Piece.White

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

	pieceCopy := opponentBoard[option.Piece.LastPosition]

	if option.EnPassant != 0 {
		return e.enPassant(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, option.EnPassant, pieceCopy)
	}

	return e.normal(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, pieceCopy)
}

func (e *Engine) enPassant(myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece, position int, moveTo int, take int, selectedPiece *pieces.Piece) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	delete(myBoard, take)

	opponentBoard[moveTo] = selectedPiece
	delete(opponentBoard, position)
	selectedPiece = nil

	return calculateAllPositions(myBoard, opponentBoard)
}

func (e *Engine) normal(myBoard map[int]*pieces.Piece, opponentBoard map[int]*pieces.Piece, position int, option int, selectedPiece *pieces.Piece) (map[int]*pieces.Piece, map[int]*pieces.Piece) {
	if selectedPiece.Kind == pieces.Pawn {
		end := 7
		if selectedPiece.White == true {
			end = 0
		}
		if position/8 == end {
			selectedPiece.Kind = pieces.Queen // todo: add convert to better Piece logic
		}
	}

	delete(myBoard, option)

	if selectedPiece.Kind == pieces.King && !selectedPiece.HasBeenMoved {
		castled, newOpponentBoard := e.castle(option, opponentBoard, selectedPiece)
		if castled {
			return calculateAllPositions(myBoard, newOpponentBoard)
		}
	}

	selectedPiece.HasBeenMoved = true

	opponentBoard[option] = selectedPiece
	delete(opponentBoard, selectedPiece.LastPosition)
	selectedPiece = nil

	return calculateAllPositions(myBoard, opponentBoard)
}

func (e *Engine) castle(option int, board map[int]*pieces.Piece, selectedPiece *pieces.Piece) (bool, map[int]*pieces.Piece) {
	switch option {
	case 2:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[3] = board[0]
		delete(board, selectedPiece.LastPosition)
		delete(board, 0)
		selectedPiece = nil
		return true, board
	case 6:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[5] = board[7]
		delete(board, selectedPiece.LastPosition)
		delete(board, 7)
		selectedPiece = nil
		return true, board
	case 58:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[59] = board[56]
		delete(board, selectedPiece.LastPosition)
		delete(board, 56)
		selectedPiece = nil
		return true, board
	case 62:
		selectedPiece.HasBeenMoved = true
		board[option] = selectedPiece
		board[61] = board[63]
		delete(board, selectedPiece.LastPosition)
		delete(board, 63)
		selectedPiece = nil
		return true, board
	}

	return false, nil
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
