package v1

import (
	"chess/pieces/v1"
)

func (a *App) calculateAllPositions(whiteBoard map[int]*v1.Piece, blackBoard map[int]*v1.Piece) {
	forbiddenSquares := make(map[int]struct{})
	var check bool

	for _, piece := range whiteBoard {
		piece.PinnedToKing = false
	}
	for _, piece := range blackBoard {
		piece.PinnedToKing = false
	}

	var checkingPieces map[int]*v1.Piece
	var kingPosition int

	switch a.whitesTurn {
	case true:
		for position, piece := range blackBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(whiteBoard, blackBoard, position, nil, false)
			check = check || tmpCheck
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if piece.Kind != v1.Pawn {
				for forbidden := range piece.Options {
					forbiddenSquares[forbidden] = struct{}{}
				}
			}
			if piece.Kind == v1.King {
				piece.CheckingPieces = make(map[int]*v1.Piece)
			}
		}

		for position, piece := range whiteBoard {
			piece.CalculateOptions(whiteBoard, blackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == v1.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range whiteBoard {
				if piece.Kind != v1.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	case false:
		for position, piece := range whiteBoard {
			forbiddenCaptures, tmpCheck := piece.CalculateOptions(whiteBoard, blackBoard, position, nil, false)
			check = check || tmpCheck
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if piece.Kind != v1.Pawn {
				for forbidden := range piece.Options {
					forbiddenSquares[forbidden] = struct{}{}
				}
			}
			if piece.Kind == v1.King {
				piece.CheckingPieces = make(map[int]*v1.Piece)
			}
		}

		for position, piece := range blackBoard {
			piece.CalculateOptions(whiteBoard, blackBoard, position, forbiddenSquares, true)
			if check && piece.Kind == v1.King {
				checkingPieces = piece.CheckingPieces
				kingPosition = position
			}
		}
		if check {
			for _, piece := range blackBoard {
				if piece.Kind != v1.King {
					piece.RemoveOptionsDueToCheck(kingPosition, checkingPieces)
				}
			}
		}
	}
}
