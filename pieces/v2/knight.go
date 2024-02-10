package v2

type Knight struct {
	Value          int
	EvaluatedValue int
	White          bool
	LastPosition   int
	Options        map[int]struct{}
	Protecting     map[int]PieceInterface
	AttackedBy     map[int]PieceInterface
	ProtectedBy    map[int]PieceInterface

	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    PieceInterface
}

func (k *Knight) GetValue() int {
	return k.Value
}
func (k *Knight) GetEvaluatedValue() int {
	return k.EvaluatedValue
}
func (k *Knight) GetWhite() bool {
	return k.White
}
func (k *Knight) GetLastPosition() int {
	return k.LastPosition
}
func (k *Knight) GetOptions() map[int]struct{} {
	return k.Options
}
func (k *Knight) GetProtecting() map[int]PieceInterface {
	return k.Protecting
}
func (k *Knight) GetAttackedBy() map[int]PieceInterface {
	return k.AttackedBy
}
func (k *Knight) GetProtectedBy() map[int]PieceInterface {
	return k.ProtectedBy
}
func (k *Knight) SetValue(value int) {
	k.Value = value
}
func (k *Knight) SetEvaluatedValue(evaluatedValue int) {
	k.EvaluatedValue = evaluatedValue
}
func (k *Knight) SetWhite(white bool) {
	k.White = white
}
func (k *Knight) SetLastPosition(lastPosition int) {
	k.LastPosition = lastPosition
}
func (k *Knight) SetOptions(options map[int]struct{}) {
	k.Options = options
}
func (k *Knight) SetProtecting(protecting map[int]PieceInterface) {
	k.Protecting = protecting
}
func (k *Knight) SetAttackedBy(attackedBy map[int]PieceInterface) {
	k.AttackedBy = attackedBy
}
func (k *Knight) SetProtectedBy(protectedBy map[int]PieceInterface) {
	k.ProtectedBy = protectedBy
}

func (k *Knight) CalculateMoves(whiteBoard map[int]PieceInterface, blackBoard map[int]PieceInterface, position int, forbiddenSquares map[int]struct{}, fixLastPosition bool) map[int]struct{} {
	if fixLastPosition {
		k.LastPosition = position
	}

	myBoard := whiteBoard
	opponentBoard := blackBoard
	if k.White == false {
		myBoard = blackBoard
		opponentBoard = whiteBoard
	}

	right := position % 8
	down := position / 8
	up := -8
	left := -1
	if right-2 >= 0 { // left 2 ok
		if down-1 >= 0 { // up 1 ok
			newPosition := position + left*2 + up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 { // down 1 ok
			newPosition := position + left*2 - up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if right+2 <= 7 {
		if down-1 >= 0 {
			newPosition := position - left*2 + up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if down+1 <= 7 {
			newPosition := position - left*2 - up
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down+2 <= 7 {
		if right-1 >= 0 {
			newPosition := position - up*2 + left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position - up*2 - left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	if down-2 >= 0 {
		if right-1 >= 0 {
			newPosition := position + up*2 + left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
		if right+1 <= 7 {
			newPosition := position + up*2 - left
			k.calculateOptions(position, newPosition, opponentBoard, forbiddenSquares)
		}
	}

	k.deleteOptions(myBoard, position)
	if k.PinnedToKing {
		k.Options = make(map[int]struct{})
		k.Protecting = make(map[int]PieceInterface)
	}

	return forbiddenSquares
}

func (k *Knight) calculateOptions(position int, newPosition int, opponentBoard map[int]PieceInterface, forbiddenSquares map[int]struct{}) {
	opponent, ok := opponentBoard[newPosition]
	if ok && !k.PinnedToKing {
		addAttackedBy(k, opponent, position)
	}
	if ok && CheckPieceKindFromAny(opponent) == PieceKindKing {
		king := opponent.(*King)
		king.Checked = true
		king.CheckingPieces[position] = king
	}
	k.Options[newPosition] = value
	forbiddenSquares[newPosition] = value
}

func (k *Knight) deleteOptions(board map[int]PieceInterface, position int) {
	var toRemove []int
	for option := range k.Options {
		if protectedPiece, ok := board[option]; ok {
			k.Protecting[option] = protectedPiece
			addProtectedBy(k, protectedPiece, position)
			toRemove = append(toRemove, option)
		}
	}
	for _, toDelete := range toRemove {
		delete(k.Options, toDelete)
	}
}

func (k *Knight) Copy(deep bool) PieceInterface {
	if k == nil {
		return nil
	}
	copyCat := &Knight{
		Value:            k.Value,
		EvaluatedValue:   k.EvaluatedValue,
		White:            k.White,
		LastPosition:     k.LastPosition,
		Options:          k.Options,
		PinnedToKing:     k.PinnedToKing,
		PinnedByPosition: k.PinnedByPosition,
	}
	if deep {
		copyCat.PinnedByPiece = k.PinnedByPiece.Copy(false)
		copyCat.Protecting, copyCat.ProtectedBy, copyCat.AttackedBy = copyProtectingAndAttacking(k.Protecting, k.ProtectedBy, k.AttackedBy)
	}
	return copyCat
}
