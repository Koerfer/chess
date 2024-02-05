package v2

import (
	"log"
)

var value struct{}

func CalculateOptions(piece any, whiteBoard map[int]any, blackBoard map[int]any, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn:
		pawn := piece.(*Pawn)
		pawn.Options = make(map[int]struct{})
		return pawn.CalculateMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindKnight:
		knight := piece.(*Knight)
		knight.Options = make(map[int]struct{})
		return knight.CalculateMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindBishop:
		bishop := piece.(*Bishop)
		bishop.Options = make(map[int]struct{})
		return bishop.CalculateMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindRook:
		rook := piece.(*Rook)
		rook.Options = make(map[int]struct{})
		return rook.CalculateMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindQueen:
		queen := piece.(*Queen)
		queen.Options = make(map[int]struct{})
		return queen.CalculateMoves(whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindKing:
		king := piece.(*King)
		king.Options = make(map[int]struct{})
		return king.CalculateMoves(whiteBoard, blackBoard, position, forbiddenSquares, fixLastPosition)
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating options")
	}

	return nil
}
