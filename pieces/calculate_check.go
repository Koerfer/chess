package pieces

func (p *Piece) RemoveOptionsDueToCheck(kingPosition int, checkingPieces map[int]*Piece) {
	if len(checkingPieces) > 1 {
		p.Options = make(map[int]struct{})
	} else {
		var checkingPiecePosition int
		var checkingPieceOptions map[int]struct{}
		for pos, piece := range checkingPieces {
			checkingPiecePosition = pos
			checkingPieceOptions = piece.Options
		}
		blockOptions := calculateBlockOptionPositions(kingPosition, checkingPiecePosition, checkingPieceOptions)

		for option := range p.Options {
			if _, ok := blockOptions[option]; !ok {
				delete(p.Options, option)
			}
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
