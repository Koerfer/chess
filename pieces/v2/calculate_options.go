package v2

var value struct{}

func CalculateOptions(piece PieceInterface, whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn, PieceKindKnight, PieceKindBishop, PieceKindRook, PieceKindQueen:
		return piece.CalculateMoves(whiteBoard, blackBoard, position, make(map[int]struct{}), fixLastPosition)
	case PieceKindKing:
		return piece.CalculateMoves(whiteBoard, blackBoard, position, forbiddenSquares, fixLastPosition)
	case PieceKindInvalid:
		panic("invalid piece kind when calculating options")
	}

	return nil
}
