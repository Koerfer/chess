package v2

import (
	v2 "chess/pieces/v2"
)

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
		panic("invalid piece kind when resetting pinned")
	}
}

func (a *App) calculateAllPositions(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface) {
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

	switch a.whitesTurn {
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
