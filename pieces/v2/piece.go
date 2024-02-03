package v2

type Piece struct {
	value int

	White            bool
	LastPosition     int
	Options          map[int]struct{}
	EnPassantOptions map[int]int
	HasBeenMoved     bool

	PinnedToKing     bool
	PinnedByPosition int
	PinnedByPiece    *Piece
}

type King struct {
	Piece

	CheckingPieces map[int]*Piece
	Checked        bool
}
