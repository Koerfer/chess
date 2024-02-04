package v2

import "reflect"

type Piece struct {
	value          int
	evaluatedValue int
	White          bool
	LastPosition   int
	Options        map[int]struct{}
	Protecting     map[int]any
}

type PieceKind int8

const (
	PieceKindInvalid PieceKind = iota
	PieceKindPawn
	PieceKindKnight
	PieceKindBishop
	PieceKindRook
	PieceKindQueen
	PieceKindKing
)

func CheckPieceKindFromAny(piece any) PieceKind {
	switch reflect.TypeOf(piece) {
	case reflect.TypeOf(&Pawn{}):
		return PieceKindPawn
	case reflect.TypeOf(&Bishop{}):
		return PieceKindBishop
	case reflect.TypeOf(&Rook{}):
		return PieceKindRook
	case reflect.TypeOf(&Queen{}):
		return PieceKindQueen
	case reflect.TypeOf(&Knight{}):
		return PieceKindKnight
	case reflect.TypeOf(&King{}):
		return PieceKindKing
	default:
		return PieceKindInvalid
	}
}

func mergeMaps(m1 map[int]struct{}, m2 map[int]struct{}) map[int]struct{} {
	for k, _ := range m2 {
		m1[k] = value
	}
	return m1
}
