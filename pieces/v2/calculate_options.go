package v2

import (
	"log"
)

var value struct{}

func CalculateOptions(piece any, whiteBoard map[int]any, blackBoard map[int]any, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	switch CheckPieceKindFromAny(piece) {
	case PieceKindPawn:
		p := piece.(*Pawn)
		p.Options = make(map[int]struct{})
		return CalculatePawnMoves(p, whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindKnight:
		p := piece.(*Knight)
		p.Options = make(map[int]struct{})
		return CalculateKnightMoves(p, whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindBishop:
		p := piece.(*Bishop)
		p.Options = make(map[int]struct{})
		return CalculateBishopMoves(p, whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindRook:
		p := piece.(*Rook)
		p.Options = make(map[int]struct{})
		return CalculateRookMoves(p, whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindQueen:
		p := piece.(*Queen)
		p.Options = make(map[int]struct{})
		return CalculateQueenMoves(p, whiteBoard, blackBoard, position, fixLastPosition)
	case PieceKindKing:
		p := piece.(*King)
		p.Options = make(map[int]struct{})
		return CalculateKingMoves(p, whiteBoard, blackBoard, position, forbiddenSquares, fixLastPosition)
	case PieceKindInvalid:
		log.Fatal("invalid piece kind when calculating options")
	}

	return nil
}
