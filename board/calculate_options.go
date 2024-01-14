package board

var value struct{}

func (p *Piece) calculateOptions(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) {
	p.options = make(map[int]struct{})
	switch p.white {
	case true:
		switch p.kind {
		case Pawn:
			if position/8 == 0 { // top of board
				p.kind = Queen
				return // add convert to better Piece logic
			}
			if position%8 == 0 { // if left on board
				if _, ok := blackBoard[position-7]; ok { // if black Piece up right
					p.options[position-7] = value // add capture move
				}

			} else if position%8 == 7 { // if right on board
				if _, ok := blackBoard[position-9]; ok { // if black Piece up left
					p.options[position-9] = value // add capture move
				}
			} else { /// if in middle
				if _, ok := blackBoard[position-7]; ok { // if black Piece up right
					p.options[position-7] = value // add capture move
				}
				if _, ok := blackBoard[position-9]; ok { // if black Piece up left
					p.options[position-9] = value // add capture move
				}
			}

			if position/8 == 6 {
				if _, ok := blackBoard[position-16]; ok {
					p.options[position-8] = value
					p.simpleDelete(whiteBoard)
					return
				}
				p.options[position-8] = value
				p.options[position-16] = value
				p.goFarDelete(whiteBoard)
				return
			}
			if _, ok := blackBoard[position-8]; ok {
				return
			}
			p.options[position-8] = value
			p.simpleDelete(whiteBoard)
			return
		case Knight:
			right := position % 8
			down := position / 8
			up := -8
			left := -1
			if right-2 >= 0 { // left 2 ok
				if down-1 >= 0 { // up 1 ok
					p.options[position+left*2+up] = value
				}
				if down+1 <= 7 { // down 1 ok
					p.options[position+left*2-up] = value
				}
			}

			if right+2 <= 7 {
				if down-1 >= 0 {
					p.options[position-left*2+up] = value
				}
				if down+1 <= 7 {
					p.options[position-left*2-up] = value
				}
			}

			if down+2 <= 7 {
				if right-1 >= 0 {
					p.options[position-up*2+left] = value
				}
				if right+1 <= 7 {
					p.options[position-up*2-left] = value
				}
			}

			if down-2 >= 0 {
				if right-1 >= 0 {
					p.options[position+up*2+left] = value
				}
				if right+1 <= 7 {
					p.options[position+up*2-left] = value
				}
			}

			p.simpleDelete(whiteBoard)
		case Bishop:
			rowPos := position % 8
			colPos := position / 8

			for leftUp := 1; leftUp <= colPos; leftUp++ {
				newPosition := position - leftUp*9
				if newPosition < 0 {
					break
				}
				if rowPos-newPosition%8 < 0 {
					break
				}
				if _, ok := whiteBoard[newPosition]; ok {
					break
				}
				if _, ok := blackBoard[newPosition]; ok {
					p.options[newPosition] = value
					break
				}

				p.options[newPosition] = value
			}
			for rightUp := 1; rightUp <= colPos; rightUp++ {
				newPosition := position - rightUp*7
				if newPosition%8-rowPos < 0 {
					break
				}
				if _, ok := whiteBoard[newPosition]; ok {
					break
				}
				if _, ok := blackBoard[newPosition]; ok {
					p.options[newPosition] = value
					break
				}

				p.options[newPosition] = value
			}
			for leftDown := 1; leftDown <= 7-colPos; leftDown++ {
				newPosition := position + leftDown*7
				if rowPos-newPosition%8 < 0 {
					break
				}
				if _, ok := whiteBoard[newPosition]; ok {
					break
				}
				if _, ok := blackBoard[newPosition]; ok {
					p.options[newPosition] = value
					break
				}

				p.options[newPosition] = value
			}
			for rightDown := 1; rightDown <= 7-colPos; rightDown++ {
				newPosition := position + rightDown*9
				if newPosition%8-rowPos < 0 {
					break
				}
				if _, ok := whiteBoard[newPosition]; ok {
					break
				}
				if _, ok := blackBoard[newPosition]; ok {
					p.options[newPosition] = value
					break
				}

				p.options[newPosition] = value
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

func (p *Piece) simpleDelete(board map[int]*Piece) {
	var toRemove []int
	for option, _ := range p.options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(p.options, toDelete)
	}
}
func (p *Piece) goFarDelete(board map[int]*Piece) {
	var toRemove []int
	switch p.kind {
	case Pawn:
		for option := range p.options {
			if _, ok := board[option]; ok {
				toRemove = append(toRemove, option)
				if _, ok := p.options[option-8]; ok {
					toRemove = append(toRemove, option-8)
				} else {
					break
				}
			}
		}
	}
	for _, toDelete := range toRemove {
		if _, ok := p.options[toDelete]; ok {
			delete(p.options, toDelete)
		}
	}
}
