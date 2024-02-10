package v2

var value struct{}

func CalculateOptions(piece PieceInterface, whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn, PieceKindKnight, PieceKindBishop, PieceKindRook, PieceKindQueen:
		piece.SetOptions(make(map[int]struct{}))
		piece.SetProtectedBy(make(map[int]PieceInterface))
		piece.SetAttackedBy(make(map[int]PieceInterface))
		piece.SetProtecting(make(map[int]PieceInterface))
		return piece.CalculateMoves(whiteBoard, blackBoard, position, make(map[int]struct{}), fixLastPosition)
	case PieceKindKing:
		piece.SetOptions(make(map[int]struct{}))
		piece.SetProtecting(make(map[int]PieceInterface))
		return piece.CalculateMoves(whiteBoard, blackBoard, position, forbiddenSquares, fixLastPosition)
	case PieceKindInvalid:
		panic("invalid piece kind when calculating options")
	}

	return nil
}
