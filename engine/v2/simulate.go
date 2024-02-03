package v2

import (
	"github.com/kingledion/go-tools/tree"
	"slices"
)

func (e *Engine) createNewBoardState(node tree.Node[*Option]) (map[int]*v2.Piece, map[int]*v2.Piece) {
	tmpWhiteBoard := make(map[int]*v2.Piece)
	tmpBlackBoard := make(map[int]*v2.Piece)
	var parents []tree.Node[*Option]

	tmpNode := node
	for {
		if tmpNode.GetParent().GetID() == e.tree.Root().GetID() {
			break
		}
		parents = append(parents, tmpNode.GetParent())
		tmpNode = tmpNode.GetParent()
	}
	slices.Reverse(parents)

	tmpWhiteBoard = e.whiteBoard
	tmpBlackBoard = e.blackBoard

	for _, parent := range parents {
		myBoard := make(map[int]*v2.Piece)
		opponentBoard := make(map[int]*v2.Piece)
		option := parent.GetData()
		white := option.Piece.White

		switch white {
		case true:
			for k, v := range tmpWhiteBoard {
				myBoard[k] = newPieceFromPointer(v)
			}
			for k, v := range tmpBlackBoard {
				opponentBoard[k] = newPieceFromPointer(v)
			}

			pieceCopy := myBoard[option.Piece.LastPosition]
			if option.EnPassant != 0 {
				tmpWhiteBoard, tmpBlackBoard = e.enPassant(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, option.EnPassant, pieceCopy)
			}

			tmpWhiteBoard, tmpBlackBoard = e.normal(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, pieceCopy)
		case false:
			for k, v := range tmpBlackBoard {
				myBoard[k] = newPieceFromPointer(v)
			}
			for k, v := range tmpWhiteBoard {
				opponentBoard[k] = newPieceFromPointer(v)
			}

			pieceCopy := myBoard[option.Piece.LastPosition]
			if option.EnPassant != 0 {
				tmpBlackBoard, tmpWhiteBoard = e.enPassant(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, option.EnPassant, pieceCopy)
			}

			tmpBlackBoard, tmpWhiteBoard = e.normal(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, pieceCopy)
		}
	}

	myBoard := make(map[int]*v2.Piece)
	opponentBoard := make(map[int]*v2.Piece)
	option := node.GetData()
	white := option.Piece.White

	switch white {
	case true:
		for k, v := range tmpWhiteBoard {
			myBoard[k] = newPieceFromPointer(v)
		}
		for k, v := range tmpBlackBoard {
			opponentBoard[k] = newPieceFromPointer(v)
		}
	case false:
		for k, v := range tmpBlackBoard {
			myBoard[k] = newPieceFromPointer(v)
		}
		for k, v := range tmpWhiteBoard {
			opponentBoard[k] = newPieceFromPointer(v)
		}
	}

	pieceCopy := myBoard[option.Piece.LastPosition]

	if option.EnPassant != 0 {
		return e.enPassant(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, option.EnPassant, pieceCopy)
	}

	return e.normal(myBoard, opponentBoard, option.Piece.LastPosition, option.MoveTo, pieceCopy)
}

func newPieceFromPointer(piece *v2.Piece) *v2.Piece {
	if piece == nil {
		return nil
	}
	checkingPieces := make(map[int]*v2.Piece)
	if piece.Kind == v2.King {
		for checkingPosition, checkingPiece := range piece.CheckingPieces {
			checkingPieces[checkingPosition] = &v2.Piece{
				Kind:             checkingPiece.Kind,
				White:            checkingPiece.White,
				LastPosition:     checkingPiece.LastPosition,
				Options:          checkingPiece.Options,
				EnPassantOptions: checkingPiece.EnPassantOptions,
				HasBeenMoved:     checkingPiece.HasBeenMoved,
				Checked:          checkingPiece.Checked,
				CheckingPieces:   checkingPiece.CheckingPieces,
				PinnedToKing:     checkingPiece.PinnedToKing,
				PinnedByPosition: checkingPiece.PinnedByPosition,
			}
		}
	}
	return &v2.Piece{
		Kind:             piece.Kind,
		White:            piece.White,
		LastPosition:     piece.LastPosition,
		Options:          piece.Options,
		EnPassantOptions: piece.EnPassantOptions,
		HasBeenMoved:     piece.HasBeenMoved,
		Checked:          piece.Checked,
		CheckingPieces:   checkingPieces,
		PinnedToKing:     piece.PinnedToKing,
		PinnedByPosition: piece.PinnedByPosition,
	}
}

func (e *Engine) enPassant(myBoard map[int]*v2.Piece, opponentBoard map[int]*v2.Piece, position int, moveTo int, take int, selectedPiece *v2.Piece) (map[int]*v2.Piece, map[int]*v2.Piece) {
	delete(opponentBoard, take)
	white := selectedPiece.White

	myBoard[moveTo] = selectedPiece
	delete(myBoard, position)
	selectedPiece = nil

	return calculateAllPositions(myBoard, opponentBoard, white)
}

func (e *Engine) normal(myBoard map[int]*v2.Piece, opponentBoard map[int]*v2.Piece, position int, option int, selectedPiece *v2.Piece) (map[int]*v2.Piece, map[int]*v2.Piece) {
	white := selectedPiece.White
	if selectedPiece.Kind == v2.Pawn {
		end := 7
		if white == true {
			end = 0
		}
		if position/8 == end {
			selectedPiece.Kind = v2.Queen // todo: add convert to better Piece logic
		}
	}

	delete(opponentBoard, option)

	if selectedPiece.Kind == v2.King && !selectedPiece.HasBeenMoved {
		castled, newOpponentBoard := e.castle(option, myBoard, selectedPiece)
		if castled {
			return calculateAllPositions(myBoard, newOpponentBoard, selectedPiece.White)
		}
	}

	selectedPiece.HasBeenMoved = true

	myBoard[position] = selectedPiece
	delete(myBoard, selectedPiece.LastPosition)
	selectedPiece = nil

	return calculateAllPositions(myBoard, opponentBoard, white)
}

func (e *Engine) castle(option int, board map[int]*v2.Piece, selectedPiece *v2.Piece) (bool, map[int]*v2.Piece) {
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

func calculateAllPositions(myBoard map[int]*v2.Piece, opponentBoard map[int]*v2.Piece, white bool) (map[int]*v2.Piece, map[int]*v2.Piece) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range myBoard {
		piece.PinnedToKing = false
	}
	for _, piece := range opponentBoard {
		piece.PinnedToKing = false
	}

	var whiteBoard map[int]*v2.Piece
	var blackBoard map[int]*v2.Piece
	switch white {
	case true:
		whiteBoard = myBoard
		blackBoard = opponentBoard
	case false:
		whiteBoard = opponentBoard
		blackBoard = myBoard
	}

	var checkingPieces map[int]*v2.Piece
	var kingPosition int

	for position, piece := range opponentBoard {
		forbiddenCaptures, tmpCheck := piece.CalculateOptions(whiteBoard, blackBoard, position, nil, false)
		check = check || tmpCheck
		for forbidden := range forbiddenCaptures {
			forbiddenSquares[forbidden] = struct{}{}
		}
		if piece.Kind != v2.Pawn {
			for forbidden := range piece.Options {
				forbiddenSquares[forbidden] = struct{}{}
			}
		}
		if piece.Kind == v2.King {
			piece.CheckingPieces = make(map[int]*v2.Piece)
		}
	}

	for position, piece := range myBoard {
		piece.CalculateOptions(whiteBoard, blackBoard, position, forbiddenSquares, true)
		if check && piece.Kind == v2.King {
			checkingPieces = piece.CheckingPieces
			kingPosition = position
		}
	}
	if check {
		for _, piece := range myBoard {
			if piece.Kind != v2.King {
				piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
			}
		}
	}

	return myBoard, opponentBoard
}
