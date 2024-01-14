package board

var value struct{}

func (p Piece) calculateOptions(whiteBoard map[*Piece]struct{}, blackBoard map[*Piece]struct{}) {
	p.options = make(map[int]struct{})
	switch p.white {
	case true:
		switch p.kind {
		case Pawn:
			if p.currentPosition/8 == 0 { // top of board
				p.kind = Queen
				return // add convert to better Piece logic
			}
			if p.currentPosition%8 == 0 { // if left on board
				if blackBoard[p.currentPosition-7] != 0 { // if black Piece up right
					m.options[p.currentPosition-7] = value // add capture move
				}

			} else if p.currentPosition%8 == 7 { // if right on board
				if blackBoard[p.currentPosition-9] != 0 { // if black Piece up left
					m.options[p.currentPosition-9] = value // add capture move
				}
			} else { /// if in middle
				if blackBoard[p.currentPosition-7] != 0 { // if black Piece up right
					m.options[p.currentPosition-7] = value // add capture move
				}
				if blackBoard[p.currentPosition-9] != 0 { // if black Piece up left
					m.options[p.currentPosition-9] = value // add capture move
				}
			}

			if p.currentPosition/8 == 6 {
				if blackBoard[p.currentPosition-16] != 0 {
					m.options[p.currentPosition-8] = value
					m.simpleDelete()
					return
				}
				m.options[p.currentPosition-8] = value
				m.options[p.currentPosition-16] = value
				m.goFarDelete(p.kind)
				return
			}
			if blackBoard[p.currentPosition-8] != 0 {
				return
			}
			m.options[p.currentPosition-8] = value
			m.simpleDelete()
			return
		case Knight:
			right := p.currentPosition % 8
			down := p.currentPosition / 8
			up := -8
			left := -1
			if right-2 >= 0 { // left 2 ok
				if down-1 >= 0 { // up 1 ok
					m.options[p.currentPosition+left*2+up] = value
				}
				if down+1 <= 7 { // down 1 ok
					m.options[p.currentPosition+left*2-up] = value
				}
			}

			if right+2 <= 7 {
				if down-1 >= 0 {
					m.options[p.currentPosition-left*2+up] = value
				}
				if down+1 <= 7 {
					m.options[p.currentPosition-left*2-up] = value
				}
			}

			if down+2 <= 7 {
				if right-1 >= 0 {
					m.options[p.currentPosition-up*2+left] = value
				}
				if right+1 <= 7 {
					m.options[p.currentPosition-up*2-left] = value
				}
			}

			if down-2 >= 0 {
				if right-1 >= 0 {
					m.options[p.currentPosition+up*2+left] = value
				}
				if right+1 <= 7 {
					m.options[p.currentPosition+up*2-left] = value
				}
			}

			m.simpleDelete()
		case Bishop:
			rowPos := p.currentPosition % 8
			colPos := p.currentPosition / 8

			for leftUp := 1; leftUp <= colPos; leftUp++ {
				newPosition := p.currentPosition - leftUp*9
				if newPosition < 0 {
					break
				}
				if rowPos-newPosition%8 < 0 {
					break
				}
				if m.whiteBoard[newPosition] != 0 {
					break
				}
				if blackBoard[newPosition] != 0 {
					m.options[newPosition] = value
					break
				}

				m.options[newPosition] = value
			}
			for rightUp := 1; rightUp <= colPos; rightUp++ {
				newPosition := p.currentPosition - rightUp*7
				if newPosition%8-rowPos < 0 {
					break
				}
				if m.whiteBoard[newPosition] != 0 {
					break
				}
				if blackBoard[newPosition] != 0 {
					m.options[newPosition] = value
					break
				}

				m.options[newPosition] = value
			}
			for leftDown := 1; leftDown <= 7-colPos; leftDown++ {
				newPosition := p.currentPosition + leftDown*7
				if rowPos-newPosition%8 < 0 {
					break
				}
				if m.whiteBoard[newPosition] != 0 {
					break
				}
				if blackBoard[newPosition] != 0 {
					m.options[newPosition] = value
					break
				}

				m.options[newPosition] = value
			}
			for rightDown := 1; rightDown <= 7-colPos; rightDown++ {
				newPosition := p.currentPosition + rightDown*9
				if newPosition%8-rowPos < 0 {
					break
				}
				if m.whiteBoard[newPosition] != 0 {
					break
				}
				if blackBoard[newPosition] != 0 {
					m.options[newPosition] = value
					break
				}

				m.options[newPosition] = value
			}

		case Rook:
		case Queen:
		case King:
		}

	case false:
		switch p.kind {
		case Pawn:
		case Knight:
		case Bishop:
		case Rook:
		case Queen:
		case King:
		}
	}
}

func (a *App) simpleDelete() {
	var toRemove []int
	for option, _ := range a.options {
		if a.whiteBoard[option] != 0 {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(a.options, toDelete)
	}
}
func (a *App) goFarDelete(piece int8) {
	var toRemove []int
	switch piece {
	case 1:
		for option := range a.options {
			if a.whiteBoard[option] != 0 {
				toRemove = append(toRemove, option)
				if _, ok := a.options[option-8]; ok {
					toRemove = append(toRemove, option-8)
				} else {
					break
				}
			}
		}
	}
	for _, toDelete := range toRemove {
		if _, ok := a.options[toDelete]; ok {
			delete(a.options, toDelete)
		}
	}
}
