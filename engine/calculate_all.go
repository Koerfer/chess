package engine

import "chess/pieces"

func (e *Engine) calculateAllPositions(whitesTurn bool) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range e.whiteBoard {
		piece.PinnedToKing = false
	}
	for _, piece := range e.blackBoard {
		piece.PinnedToKing = false
	}

	var checkingPieces map[int]*pieces.Piece
	var kingPosition int

	switch whitesTurn {
	case true:
		for position, piece := range e.blackBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(e.whiteBoard, e.blackBoard, position, nil, false)
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

		for position, piece := range e.whiteBoard {
			piece.CalculateOptions(e.whiteBoard, e.blackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == pieces.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range e.whiteBoard {
				if piece.Kind != pieces.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	case false:
		for position, piece := range e.whiteBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(e.whiteBoard, e.blackBoard, position, nil, false)
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

		for position, piece := range e.blackBoard {
			piece.CalculateOptions(e.whiteBoard, e.blackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == pieces.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range e.blackBoard {
				if piece.Kind != pieces.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	}
}

func (e *Engine) calculateAllPositionsNew(whitesTurn bool) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range e.newWhiteBoard {
		piece.PinnedToKing = false
	}
	for _, piece := range e.newBlackBoard {
		piece.PinnedToKing = false
	}

	var checkingPieces map[int]*pieces.Piece
	var kingPosition int

	switch whitesTurn {
	case true:
		for position, piece := range e.newBlackBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(e.newWhiteBoard, e.newBlackBoard, position, nil, false)
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

		for position, piece := range e.newWhiteBoard {
			piece.CalculateOptions(e.newWhiteBoard, e.newBlackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == pieces.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range e.newWhiteBoard {
				if piece.Kind != pieces.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	case false:
		for position, piece := range e.newWhiteBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(e.newWhiteBoard, e.newBlackBoard, position, nil, false)
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

		for position, piece := range e.newBlackBoard {
			piece.CalculateOptions(e.newWhiteBoard, e.newBlackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == pieces.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range e.newBlackBoard {
				if piece.Kind != pieces.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	}
}
