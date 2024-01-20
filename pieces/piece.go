package pieces

type Piece struct {
	Kind             PieceKind
	White            bool
	LastPosition     int
	Options          map[int]struct{}
	EnPassantOptions map[int]int
	HasBeenMoved     bool
	Checked          bool
	CheckingPieces   map[int]*Piece
	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    *Piece
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
