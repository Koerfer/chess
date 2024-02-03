package v1

var value struct{}

func (p *Piece) CalculateOptions(whiteBoard map[int]*Piece, blackBoard map[int]*Piece, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) (map[int]struct{}, bool) {
	p.Options = make(map[int]struct{})

	switch p.Kind {
	case Pawn:
		return p.calculatePawnMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case Knight:
		return p.calculateKnightMoves(whiteBoard, blackBoard, position)
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
