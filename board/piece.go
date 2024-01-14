package board

type Piece struct {
	kind            PieceKind
	white           bool
	currentPosition int
	lastPosition    int
	options         map[int]struct{}
}

type PieceKind int8

const (
	Pawn   PieceKind = 1
	Knight PieceKind = 2
	Bishop PieceKind = 3
	Rook   PieceKind = 5
	Queen  PieceKind = 9
	King   PieceKind = 100
)
