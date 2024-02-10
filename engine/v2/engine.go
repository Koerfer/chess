package v2

import (
	v2 "chess/pieces/v2"
)

type Engine struct {
	whiteBoard map[int]v2.PieceInterface
	blackBoard map[int]v2.PieceInterface

	whiteEval int
	blackEval int
}

func (e *Engine) Init(whiteBoard map[int]v2.PieceInterface, blackBoard map[int]v2.PieceInterface) {
	e.copyBoard(whiteBoard, blackBoard)
}

//func (e *Engine) evaluate() {
//	for pos, piece := range e.whiteBoard {
//		switch v2.CheckPieceKindFromAny(piece) {
//		case v2.PieceKindPawn:
//			pawn := piece.(*v2.Pawn)
//			var protectingValue int
//			for _, protectedPiece := range pawn.Protecting {
//				protectingValue += protectedPiece
//			}
//		case v2.PieceKindKnight:
//			knight := piece.(*v2.Knight)
//
//		case v2.PieceKindBishop:
//			bishop := piece.(*v2.Bishop)
//
//		case v2.PieceKindRook:
//			rook := piece.(*v2.Rook)
//
//		case v2.PieceKindQueen:
//			queen := piece.(*v2.Queen)
//
//		case v2.PieceKindKing:
//			king := piece.(*v2.King)
//
//		case v2.PieceKindInvalid:
//			panic("invalid piece kind when evaluating white board")
//		}
//	}
//}
