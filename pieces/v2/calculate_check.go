package v2

import (
	"log"
)

func RemoveOptionsDueToCheck(piece any, kingPosition int, checkingPieces map[int]any) {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn:
		p := piece.(*Pawn)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculatePawnBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindKnight:
		p := piece.(*Knight)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculateKnightBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindBishop:
		p := piece.(*Bishop)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculateBishopBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindRook:
		p := piece.(*Rook)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculateRookBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindQueen:
		p := piece.(*Queen)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculateQueenBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when removing options due to check")
	}
}

func calculatePawnBlockOptionPositions(pawn *Pawn, kingPosition int, checkingPiece map[int]any) {
	blockOptions := calculateBlockPositions(kingPosition, checkingPiece)

	for option := range pawn.Options {
		if _, ok := blockOptions[option]; !ok {
			delete(pawn.Options, option)
		}
	}
	for option := range pawn.EnPassantOptions {
		if _, ok := blockOptions[option]; !ok {
			delete(pawn.Options, option)
		}
	}
}

func calculateKnightBlockOptionPositions(knight *Knight, kingPosition int, checkingPiece map[int]any) {
	blockOptions := calculateBlockPositions(kingPosition, checkingPiece)

	for option := range knight.Options {
		if _, ok := blockOptions[option]; !ok {
			delete(knight.Options, option)
		}
	}
}

func calculateBishopBlockOptionPositions(bishop *Bishop, kingPosition int, checkingPiece map[int]any) {
	blockOptions := calculateBlockPositions(kingPosition, checkingPiece)

	for option := range bishop.Options {
		if _, ok := blockOptions[option]; !ok {
			delete(bishop.Options, option)
		}
	}
}

func calculateRookBlockOptionPositions(rook *Rook, kingPosition int, checkingPiece map[int]any) {
	blockOptions := calculateBlockPositions(kingPosition, checkingPiece)

	for option := range rook.Options {
		if _, ok := blockOptions[option]; !ok {
			delete(rook.Options, option)
		}
	}
}

func calculateQueenBlockOptionPositions(queen *Queen, kingPosition int, checkingPiece map[int]any) {
	blockOptions := calculateBlockPositions(kingPosition, checkingPiece)

	for option := range queen.Options {
		if _, ok := blockOptions[option]; !ok {
			delete(queen.Options, option)
		}
	}
}

func calculateBlockPositions(kingPosition int, checkingPiece map[int]any) map[int]struct{} {
	var checkingPiecePosition int
	var checkingPieceOptions map[int]struct{}
	for pos, piece := range checkingPiece {
		checkingPiecePosition = pos
		switch CheckPieceKindFromAny(piece) {
		case PieceKindPawn:
			p := piece.(*Pawn)
			checkingPieceOptions = p.Options
		case PieceKindKnight:
			p := piece.(*Knight)
			checkingPieceOptions = p.Options
		case PieceKindBishop:
			p := piece.(*Bishop)
			checkingPieceOptions = p.Options
		case PieceKindRook:
			p := piece.(*Rook)
			checkingPieceOptions = p.Options
		case PieceKindQueen:
			p := piece.(*Queen)
			checkingPieceOptions = p.Options
		case PieceKindKing:
			p := piece.(*King)
			checkingPieceOptions = p.Options
		case PieceKindInvalid:
			log.Fatal("invalid piece kind when removing options due to check")
		}
	}
	return calculateBlockOptionPositions(kingPosition, checkingPiecePosition, checkingPieceOptions)
}

func calculateBlockOptionPositions(kingPosition int, checkingPiecePosition int, checkingPieceOptions map[int]struct{}) map[int]struct{} {
	blockOptions := make(map[int]struct{})
	blockOptions[checkingPiecePosition] = value

	for option := range checkingPieceOptions {
		if checkingPiecePosition%8 == kingPosition%8 { // same column
			if checkingPiecePosition/8 < kingPosition/8 {
				if option%8 == kingPosition%8 && option < kingPosition && option > checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
			if checkingPiecePosition/8 > kingPosition/8 {
				if option%8 == kingPosition%8 && option > kingPosition && option < checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
		} else if checkingPiecePosition/8 == kingPosition/8 { // same row
			if checkingPiecePosition%8 < kingPosition%8 {
				if option/8 == kingPosition/8 && option < kingPosition && option > checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
			if checkingPiecePosition%8 > kingPosition%8 {
				if option/8 == kingPosition/8 && option > kingPosition && option < checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
		} else if checkingPiecePosition%9 == kingPosition%9 { // same left->right column
			if checkingPiecePosition < kingPosition {
				if option%9 == kingPosition%9 && option < kingPosition && option > checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
			if checkingPiecePosition > kingPosition {
				if option%9 == kingPosition%9 && option > kingPosition && option < checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
		} else if checkingPiecePosition%7 == kingPosition%7 { // same right->left column
			if checkingPiecePosition < kingPosition {
				if option%7 == kingPosition%7 && option < kingPosition && option > checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
			if checkingPiecePosition > kingPosition {
				if option%7 == kingPosition%7 && option > kingPosition && option < checkingPiecePosition {
					blockOptions[option] = value
					continue
				}
			}
		}
	}

	return blockOptions
}
