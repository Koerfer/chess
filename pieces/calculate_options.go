package pieces

var value struct{}

func (p *Piece) CalculateOptions(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool, check bool) (map[int]struct{}, bool) {
	p.Options = make(map[int]struct{})
	if check {

	}

	switch p.Kind {
	case Pawn:
		return p.calculatePawnMoves(whiteBoard, blackBoard, position, fixLastPosition), false
	case Knight:
		return p.calculateKnightMoves(whiteBoard, blackBoard, position), false
	case Bishop:
		return p.calculateBishopMoves(whiteBoard, blackBoard, position)
	case Rook:
		return p.calculateRookMoves(whiteBoard, blackBoard, position)
	case Queen:
		forbiddenDiagonal, check1 := p.calculateBishopMoves(whiteBoard, blackBoard, position)
		forbidden, check2 := p.calculateRookMoves(whiteBoard, blackBoard, position)
		return mergeMaps(forbiddenDiagonal, forbidden), check1 || check2
	case King:
		return p.calculateKingMoves(whiteBoard, blackBoard, position, forbiddenSquares), false
	}

	return nil, false
}

func mergeMaps(m1 map[int]struct{}, m2 map[int]struct{}) map[int]struct{} {
	for k, _ := range m2 {
		m1[k] = value
	}
	return m1
}

func (p *Piece) simpleDelete(board map[int]*Piece) {
	var toRemove []int
	for option, _ := range p.Options {
		if _, ok := board[option]; ok {
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(p.Options, toDelete)
	}
}
func (p *Piece) goFarDelete(board map[int]*Piece) {
	var toRemove []int
	switch p.Kind {
	case Pawn:
		for option := range p.Options {
			if _, ok := board[option]; ok {
				toRemove = append(toRemove, option)
				if _, ok := p.Options[option-8]; ok {
					toRemove = append(toRemove, option-8)
				} else {
					break
				}
			}
		}
	}
	for _, toDelete := range toRemove {
		if _, ok := p.Options[toDelete]; ok {
			delete(p.Options, toDelete)
		}
	}
}
