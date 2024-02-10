package v2

func RemoveOptionsDueToCheck(piece PieceInterface, kingPosition int, checkingPieces map[int]PieceInterface) {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn:
		p := piece.(*Pawn)
		if len(checkingPieces) > 1 {
			p.Options = make(map[int]struct{})
		}
		calculatePawnBlockOptionPositions(p, kingPosition, checkingPieces)
	case PieceKindKnight, PieceKindBishop, PieceKindRook, PieceKindQueen:
		if len(checkingPieces) > 1 {
			piece.SetOptions(make(map[int]struct{}))
		}
		calculateBlockOptions(piece, kingPosition, checkingPieces)
	case PieceKindKing:
		// do nothing
	case PieceKindInvalid:
		panic("invalid piece kind when removing options due to check")
	}
}

func calculatePawnBlockOptionPositions(pawn *Pawn, kingPosition int, checkingPiece map[int]PieceInterface) {
	var checkingPiecePosition int
	var checkingPieceOptions map[int]struct{}
	for pos, p := range checkingPiece {
		checkingPiecePosition = pos
		checkingPieceOptions = p.GetOptions()
	}
	blockOptions := calculateBlockOptionPositions(kingPosition, checkingPiecePosition, checkingPieceOptions)

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

func calculateBlockOptions(piece PieceInterface, kingPosition int, checkingPiece map[int]PieceInterface) {
	var checkingPiecePosition int
	var checkingPieceOptions map[int]struct{}
	for pos, p := range checkingPiece {
		checkingPiecePosition = pos
		checkingPieceOptions = p.GetOptions()
	}
	blockOptions := calculateBlockOptionPositions(kingPosition, checkingPiecePosition, checkingPieceOptions)

	for option := range piece.GetOptions() {
		if _, ok := blockOptions[option]; !ok {
			options := piece.GetOptions()
			delete(options, option)
			piece.SetOptions(options)
		}
	}
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
