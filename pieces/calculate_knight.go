package pieces

func (p *Piece) calculateKnightMoves(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int) map[int]struct{} {
	forbiddenSquares := make(map[int]struct{})

	right := position % 8
	down := position / 8
	up := -8
	left := -1
	if right-2 >= 0 { // left 2 ok
		if down-1 >= 0 { // up 1 ok
			p.Options[position+left*2+up] = value
		}
		if down+1 <= 7 { // down 1 ok
			p.Options[position+left*2-up] = value
		}
	}

	if right+2 <= 7 {
		if down-1 >= 0 {
			p.Options[position-left*2+up] = value
		}
		if down+1 <= 7 {
			p.Options[position-left*2-up] = value
		}
	}

	if down+2 <= 7 {
		if right-1 >= 0 {
			p.Options[position-up*2+left] = value
		}
		if right+1 <= 7 {
			p.Options[position-up*2-left] = value
		}
	}

	if down-2 >= 0 {
		if right-1 >= 0 {
			p.Options[position+up*2+left] = value
		}
		if right+1 <= 7 {
			p.Options[position+up*2-left] = value
		}
	}
	forbiddenSquares = p.Options

	switch p.White {
	case true:
		p.simpleDelete(whiteBoard)
		p.calculatePinnedOptions(position)
	case false:
		p.simpleDelete(blackBoard)
		p.calculatePinnedOptions(position)
	}

	return forbiddenSquares
}
