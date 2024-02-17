package v2

import v2 "chess/pieces/v2"

func (s *Simulator) performMove(move *SelectedMove, board map[int]v2.PieceInterface) {
	switch v2.CheckPieceKindFromAny(move.Piece) {
	case v2.PieceKindPawn:
		p := move.Piece.(*v2.Pawn)
		if stop := s.enPassant(p, move.ToPosition, board, move.Piece); stop {
			s.recalculateBoard()
		}
		if stop := s.normalPawn(p, move.ToPosition, board); stop {
			s.recalculateBoard()
		}
	case v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
		if stop := s.normal(move.Piece, move.ToPosition, board, move.Piece); stop {
			s.recalculateBoard()
		}
	case v2.PieceKindKing:
		p := move.Piece.(*v2.King)
		if stop := s.normalKing(p, move.ToPosition, board, move.Piece); stop {
			s.recalculateBoard()
		}
	case v2.PieceKindInvalid:
		panic("invalid Piece kind iin simulator")
	}
}

func (s *Simulator) normalKing(king *v2.King, position int, board map[int]v2.PieceInterface, move v2.PieceInterface) bool {
	for option := range king.Options {
		if position != option {
			continue
		}

		switch s.whiteTurn {
		case true:
			delete(s.blackBoard, position)
		case false:
			delete(s.whiteBoard, position)
		}

		if !king.HasBeenMoved {
			castled := s.castle(king, option, board, move)
			if castled {
				king.HasCastled = true
				king.HasBeenMoved = true
				return true
			}
		}

		king.HasBeenMoved = true

		board[option] = move
		delete(board, king.LastPosition)
		return true
	}

	return false
}

func (s *Simulator) castle(king *v2.King, option int, board map[int]v2.PieceInterface, move v2.PieceInterface) bool {
	switch option {
	case 2:
		king.HasBeenMoved = true
		board[option] = move
		board[3] = board[0]
		delete(board, king.LastPosition)
		delete(board, 0)
		return true
	case 6:
		king.HasBeenMoved = true
		board[option] = move
		board[5] = board[7]
		delete(board, king.LastPosition)
		delete(board, 7)
		return true
	case 58:
		king.HasBeenMoved = true
		board[option] = move
		board[59] = board[56]
		delete(board, king.LastPosition)
		delete(board, 56)
		return true
	case 62:
		king.HasBeenMoved = true
		board[option] = move
		board[61] = board[63]
		delete(board, king.LastPosition)
		delete(board, 63)
		return true
	}

	return false
}

func (s *Simulator) enPassant(pawn *v2.Pawn, position int, board map[int]v2.PieceInterface, move v2.PieceInterface) bool {
	for option, take := range pawn.EnPassantOptions {
		if position != option {
			continue
		}

		switch s.whiteTurn {
		case true:
			delete(s.blackBoard, take)
		case false:
			delete(s.whiteBoard, take)
		}

		board[option] = move
		delete(board, pawn.LastPosition)
		return true
	}

	return false
}

func (s *Simulator) normalPawn(pawn *v2.Pawn, position int, board map[int]v2.PieceInterface) bool {
	for option := range pawn.Options {
		if position != option {
			continue
		}

		end := 7
		if pawn.White == true {
			end = 0
		}
		if position/8 == end {
			board[position] = &v2.Queen{
				White:        pawn.White,
				LastPosition: pawn.LastPosition,
				Options:      make(map[int]struct{}),

				PinnedToKing:     false,
				PinnedByPosition: 0,
				PinnedByPiece:    nil,
				Protecting:       make(map[int]v2.PieceInterface),
				Value:            9,
				AttackedBy:       make(map[int]v2.PieceInterface),
				ProtectedBy:      make(map[int]v2.PieceInterface),
			}
		} else {
			board[option] = pawn
		}

		switch s.whiteTurn {
		case true:
			delete(s.blackBoard, position)
		case false:
			delete(s.whiteBoard, position)
		}

		delete(board, pawn.LastPosition)
		return true
	}

	return false
}

func (s *Simulator) normal(piece v2.PieceInterface, position int, board map[int]v2.PieceInterface, move v2.PieceInterface) bool {
	for option := range piece.GetOptions() {
		if position != option {
			continue
		}

		switch s.whiteTurn {
		case true:
			delete(s.blackBoard, position)
		case false:
			delete(s.whiteBoard, position)
		}

		board[option] = move
		delete(board, piece.GetLastPosition())
		return true
	}

	return false
}

func (s *Simulator) recalculateBoard() {
	for _, piece := range s.whiteBoard {
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn, v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtectedBy(make(map[int]v2.PieceInterface))
			piece.SetAttackedBy(make(map[int]v2.PieceInterface))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindKing:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindInvalid:
			panic("invalid Piece kind when calculating options")
		}
	}
	for _, piece := range s.blackBoard {
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn, v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtectedBy(make(map[int]v2.PieceInterface))
			piece.SetAttackedBy(make(map[int]v2.PieceInterface))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindKing:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindInvalid:
			panic("invalid Piece kind when calculating options")
		}
	}
	s.calculateAllPositions(s.whiteBoard, s.blackBoard)
}

func (s *Simulator) calculateAllPositions(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range whiteBoard {
		resetPinned(piece)
	}
	for _, piece := range blackBoard {
		resetPinned(piece)
	}

	var checkingPieces map[int]v2.PieceInterface
	var kingPosition int

	switch !s.whiteTurn {
	case true:
		for position, piece := range blackBoard {
			forbiddenCaptures := v2.CalculateOptions(piece, whiteBoard, blackBoard, position, nil, false)
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				p.CheckingPieces = make(map[int]v2.PieceInterface)
			}
		}

		for _, piece := range whiteBoard {
			if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				check = p.Checked
				break
			}
		}

		for position, piece := range whiteBoard {
			v2.CalculateOptions(piece, whiteBoard, blackBoard, position, forbiddenSquares, true)
			if check && v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				checkingPieces = p.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range whiteBoard {
				if v2.CheckPieceKindFromAny(piece) != v2.PieceKindKing {
					v2.RemoveOptionsDueToCheck(piece, kingPosition, checkingPieces)
				}
			}
		}
	case false:
		for position, piece := range whiteBoard {
			forbiddenCaptures := v2.CalculateOptions(piece, whiteBoard, blackBoard, position, nil, false)
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				p.CheckingPieces = make(map[int]v2.PieceInterface)
			}
		}

		for _, piece := range blackBoard {
			if v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				check = p.Checked
				break
			}
		}

		for position, piece := range blackBoard {
			v2.CalculateOptions(piece, whiteBoard, blackBoard, position, forbiddenSquares, true)
			if check && v2.CheckPieceKindFromAny(piece) == v2.PieceKindKing {
				p := piece.(*v2.King)
				checkingPieces = p.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range blackBoard {
				if v2.CheckPieceKindFromAny(piece) != v2.PieceKindKing {
					v2.RemoveOptionsDueToCheck(piece, kingPosition, checkingPieces)
				}
			}
		}
	}
}

func resetPinned(piece v2.PieceInterface) {
	switch v2.CheckPieceKindFromAny(piece) {
	case v2.PieceKindPawn:
		p := piece.(*v2.Pawn)
		p.PinnedToKing = false
	case v2.PieceKindKnight:
		p := piece.(*v2.Knight)
		p.PinnedToKing = false
	case v2.PieceKindBishop:
		p := piece.(*v2.Bishop)
		p.PinnedToKing = false
	case v2.PieceKindRook:
		p := piece.(*v2.Rook)
		p.PinnedToKing = false
	case v2.PieceKindQueen:
		p := piece.(*v2.Queen)
		p.PinnedToKing = false
	case v2.PieceKindKing:
		p := piece.(*v2.King)
		p.Checked = false
	case v2.PieceKindInvalid:
		panic("invalid Piece kind when resetting pinned")
	}
}
