package pieces

type Piece struct {
	Kind         PieceKind
	White        bool
	LastPosition int
	Options      map[int]struct{}
	HasBeenMoved bool
	Checked      bool
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
