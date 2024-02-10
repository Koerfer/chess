package v2

import "reflect"

type PieceInterface interface {
	GetValue() int
	GetEvaluatedValue() int
	GetWhite() bool
	GetLastPosition() int
	GetOptions() map[int]struct{}
	GetProtecting() map[int]PieceInterface
	GetAttackedBy() map[int]PieceInterface
	GetProtectedBy() map[int]PieceInterface

	SetValue(int)
	SetEvaluatedValue(int)
	SetWhite(bool)
	SetLastPosition(int)
	SetOptions(map[int]struct{})
	SetProtecting(map[int]PieceInterface)
	SetAttackedBy(map[int]PieceInterface)
	SetProtectedBy(map[int]PieceInterface)

	CalculateMoves(map[int]PieceInterface, map[int]PieceInterface, int, map[int]struct{}, bool) map[int]struct{}
	Copy(bool) PieceInterface
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

func CheckPieceKindFromAny(piece PieceInterface) PieceKind {
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

func addAttackedBy(piece PieceInterface, attackedPiece PieceInterface, position int) {
	if CheckPieceKindFromAny(attackedPiece) == PieceKindKing {
		return
	}
	attackedBy := attackedPiece.GetAttackedBy()
	attackedBy[position] = piece
	attackedPiece.SetAttackedBy(attackedBy)
}

func addProtectedBy(piece PieceInterface, protectedPiece PieceInterface, position int) {
	if CheckPieceKindFromAny(protectedPiece) == PieceKindKing {
		return
	}
	protectedBy := protectedPiece.GetProtectedBy()
	protectedBy[position] = piece
	protectedPiece.SetProtectedBy(protectedBy)
}

func copyProtectingAndAttacking(protecting map[int]PieceInterface, protectedBy map[int]PieceInterface, attacking map[int]PieceInterface) (map[int]PieceInterface, map[int]PieceInterface, map[int]PieceInterface) {
	copyProtecting := make(map[int]PieceInterface)
	for k, v := range protecting {
		copyProtecting[k] = v.Copy(false)
	}

	copyProtectedBy := make(map[int]PieceInterface)
	if protectedBy != nil {
		for k, v := range protectedBy {
			copyProtectedBy[k] = v.Copy(false)
		}
	}

	copyAttacking := make(map[int]PieceInterface)
	if attacking != nil {
		for k, v := range attacking {
			copyAttacking[k] = v.Copy(false)
		}
	}

	return copyProtecting, copyProtectedBy, copyAttacking
}

func mergeMaps(m1 map[int]struct{}, m2 map[int]struct{}) map[int]struct{} {
	for k, _ := range m2 {
		m1[k] = value
	}
	return m1
}
