package v2

import (
	v2 "chess/pieces/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"image/jpeg"
	"image/png"
	"os"
)

func (a *App) initWhiteBoard() {
	a.whitesTurn = true
	a.whiteBoard = make(map[int]v2.PieceInterface)
	a.addPiece(56, v2.PieceKindRook, true)
	a.addPiece(57, v2.PieceKindKnight, true)
	a.addPiece(58, v2.PieceKindBishop, true)
	a.addPiece(59, v2.PieceKindQueen, true)
	a.addPiece(60, v2.PieceKindKing, true)
	a.addPiece(61, v2.PieceKindBishop, true)
	a.addPiece(62, v2.PieceKindKnight, true)
	a.addPiece(63, v2.PieceKindRook, true)

	a.addPiece(48, v2.PieceKindPawn, true)
	a.addPiece(49, v2.PieceKindPawn, true)
	a.addPiece(50, v2.PieceKindPawn, true)
	a.addPiece(51, v2.PieceKindPawn, true)
	a.addPiece(52, v2.PieceKindPawn, true)
	a.addPiece(53, v2.PieceKindPawn, true)
	a.addPiece(54, v2.PieceKindPawn, true)
	a.addPiece(55, v2.PieceKindPawn, true)
}

func (a *App) initBlackBoard() {
	a.blackBoard = make(map[int]v2.PieceInterface)
	a.addPiece(0, v2.PieceKindRook, false)
	a.addPiece(1, v2.PieceKindKnight, false)
	a.addPiece(2, v2.PieceKindBishop, false)
	a.addPiece(3, v2.PieceKindQueen, false)
	a.addPiece(4, v2.PieceKindKing, false)
	a.addPiece(5, v2.PieceKindBishop, false)
	a.addPiece(6, v2.PieceKindKnight, false)
	a.addPiece(7, v2.PieceKindRook, false)

	a.addPiece(8, v2.PieceKindPawn, false)
	a.addPiece(9, v2.PieceKindPawn, false)
	a.addPiece(10, v2.PieceKindPawn, false)
	a.addPiece(11, v2.PieceKindPawn, false)
	a.addPiece(12, v2.PieceKindPawn, false)
	a.addPiece(13, v2.PieceKindPawn, false)
	a.addPiece(14, v2.PieceKindPawn, false)
	a.addPiece(15, v2.PieceKindPawn, false)
}

func (a *App) addPiece(pos int, kind v2.PieceKind, white bool) {
	switch white {
	case true:
		switch kind {
		case v2.PieceKindPawn:
			a.whiteBoard[pos] = &v2.Pawn{
				White:            white,
				LastPosition:     pos,
				Options:          make(map[int]struct{}),
				Protecting:       make(map[int]v2.PieceInterface),
				Value:            1,
				AttackedBy:       make(map[int]v2.PieceInterface),
				ProtectedBy:      make(map[int]v2.PieceInterface),
				EnPassantOptions: make(map[int]int),
			}
		case v2.PieceKindKnight:
			a.whiteBoard[pos] = &v2.Knight{
				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        2,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindBishop:
			a.whiteBoard[pos] = &v2.Bishop{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        3,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindRook:
			a.whiteBoard[pos] = &v2.Rook{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        5,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindQueen:
			a.whiteBoard[pos] = &v2.Queen{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        9,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindKing:
			a.whiteBoard[pos] = &v2.King{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        100,

				CheckingPieces: make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind when initialising board")
		}
	case false:
		switch kind {
		case v2.PieceKindPawn:
			a.blackBoard[pos] = &v2.Pawn{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        1,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),

				EnPassantOptions: make(map[int]int),
			}
		case v2.PieceKindKnight:
			a.blackBoard[pos] = &v2.Knight{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        2,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindBishop:
			a.blackBoard[pos] = &v2.Bishop{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        3,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindRook:
			a.blackBoard[pos] = &v2.Rook{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        5,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindQueen:
			a.blackBoard[pos] = &v2.Queen{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        9,
				AttackedBy:   make(map[int]v2.PieceInterface),
				ProtectedBy:  make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindKing:
			a.blackBoard[pos] = &v2.King{

				White:        white,
				LastPosition: pos,
				Options:      make(map[int]struct{}),
				Protecting:   make(map[int]v2.PieceInterface),
				Value:        100,

				CheckingPieces: make(map[int]v2.PieceInterface),
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind when initialising board")
		}
	}

}

func (a *App) initImages() {
	chessboard, err := os.Open("board/v2/images/chessboard.jpeg")
	if err != nil {
		panic("unable to open chessboard image")
	}
	jpegBoard, err := jpeg.Decode(chessboard)
	if err != nil {
		panic("unable to decode chessboard image")
	}
	boardImage := ebiten.NewImageFromImage(jpegBoard)

	option, err := os.Open("board/v2/images/option.png")
	if err != nil {
		panic("unable to open option image")
	}
	optionDecoded, err := png.Decode(option)
	if err != nil {
		panic("unable to decode option image")
	}
	optionImage := ebiten.NewImageFromImage(optionDecoded)

	lastPosition, err := os.Open("board/v2/images/last_position.png")
	if err != nil {
		panic("unable to last position option image")
	}
	lastPositionDecoded, err := png.Decode(lastPosition)
	if err != nil {
		panic("unable to decode last position image")
	}
	lastPositionImage := ebiten.NewImageFromImage(lastPositionDecoded)

	newPosition, err := os.Open("board/v2/images/new_position_marker.png")
	if err != nil {
		panic("unable to last position option image")
	}
	newPositionDecoded, err := png.Decode(newPosition)
	if err != nil {
		panic("unable to decode last position image")
	}
	newPositionImage := ebiten.NewImageFromImage(newPositionDecoded)

	whitePawn, err := os.Open("board/v2/images/white_pawn.png")
	if err != nil {
		panic("unable to open White pawn image")
	}
	whitePawnDecoded, err := png.Decode(whitePawn)
	if err != nil {
		panic("unable to decode White pawn image")
	}
	whitePawnImage := ebiten.NewImageFromImage(whitePawnDecoded)

	whiteKing, err := os.Open("board/v2/images/white_king.png")
	if err != nil {
		panic("unable to open White king image")
	}
	whiteKingDecoded, err := png.Decode(whiteKing)
	if err != nil {
		panic("unable to decode White king image")
	}
	whiteKingImage := ebiten.NewImageFromImage(whiteKingDecoded)

	whiteRook, err := os.Open("board/v2/images/white_rook.png")
	if err != nil {
		panic("unable to open White rook image")
	}
	whiteRookDecoded, err := png.Decode(whiteRook)
	if err != nil {
		panic("unable to decode White pawn image")
	}
	whiteRookImage := ebiten.NewImageFromImage(whiteRookDecoded)

	whiteQueen, err := os.Open("board/v2/images/white_queen.png")
	if err != nil {
		panic("unable to open White queen image")
	}
	whiteQueenDecoded, err := png.Decode(whiteQueen)
	if err != nil {
		panic("unable to decode White queen image")
	}
	whiteQueenImage := ebiten.NewImageFromImage(whiteQueenDecoded)

	whiteKnight, err := os.Open("board/v2/images/white_knight.png")
	if err != nil {
		panic("unable to open White knight image")
	}
	whiteKnightDecoded, err := png.Decode(whiteKnight)
	if err != nil {
		panic("unable to decode White knight image")
	}
	whiteKnightImage := ebiten.NewImageFromImage(whiteKnightDecoded)

	whiteBishop, err := os.Open("board/v2/images/white_bishop.png")
	if err != nil {
		panic("unable to open White bishop image")
	}
	whiteBishopDecoded, err := png.Decode(whiteBishop)
	if err != nil {
		panic("unable to decode White bishop image")
	}
	whiteBishopImage := ebiten.NewImageFromImage(whiteBishopDecoded)

	blackPawn, err := os.Open("board/v2/images/black_pawn.png")
	if err != nil {
		panic("unable to open black pawn image")
	}
	blackPawnDecoded, err := png.Decode(blackPawn)
	if err != nil {
		panic("unable to decode black pawn image")
	}
	blackPawnImage := ebiten.NewImageFromImage(blackPawnDecoded)

	blackKing, err := os.Open("board/v2/images/black_king.png")
	if err != nil {
		panic("unable to open black king image")
	}
	blackKingDecoded, err := png.Decode(blackKing)
	if err != nil {
		panic("unable to decode black king image")
	}
	blackKingImage := ebiten.NewImageFromImage(blackKingDecoded)

	blackRook, err := os.Open("board/v2/images/black_rook.png")
	if err != nil {
		panic("unable to open black rook image")
	}
	blackRookDecoded, err := png.Decode(blackRook)
	if err != nil {
		panic("unable to decode black pawn image")
	}
	blackRookImage := ebiten.NewImageFromImage(blackRookDecoded)

	blackQueen, err := os.Open("board/v2/images/black_queen.png")
	if err != nil {
		panic("unable to open black queen image")
	}
	blackQueenDecoded, err := png.Decode(blackQueen)
	if err != nil {
		panic("unable to decode black queen image")
	}
	blackQueenImage := ebiten.NewImageFromImage(blackQueenDecoded)

	blackKnight, err := os.Open("board/v2/images/black_knight.png")
	if err != nil {
		panic("unable to open black knight image")
	}
	blackKnightDecoded, err := png.Decode(blackKnight)
	if err != nil {
		panic("unable to decode black knight image")
	}
	blackKnightImage := ebiten.NewImageFromImage(blackKnightDecoded)

	blackBishop, err := os.Open("board/v2/images/black_bishop.png")
	if err != nil {
		panic("unable to open black bishop image")
	}
	blackBishopDecoded, err := png.Decode(blackBishop)
	if err != nil {
		panic("unable to decode black bishop image")
	}
	blackBishopImage := ebiten.NewImageFromImage(blackBishopDecoded)

	b := boardImage.Bounds().Size()
	bI = ebiten.NewImage(b.X, b.Y)
	o := optionImage.Bounds().Size()
	optionI = ebiten.NewImage(o.X, o.Y)
	lp := lastPositionImage.Bounds().Size()
	lastPositionI = ebiten.NewImage(lp.X, lp.Y)
	np := newPositionImage.Bounds().Size()
	newPositionI = ebiten.NewImage(np.X, np.Y)

	wp := whitePawnImage.Bounds().Size()
	wr := whiteRookImage.Bounds().Size()
	wkn := whiteKnightImage.Bounds().Size()
	wki := whiteKingImage.Bounds().Size()
	wq := whiteQueenImage.Bounds().Size()
	wb := whiteBishopImage.Bounds().Size()
	wpI = ebiten.NewImage(wp.X, wp.Y)
	wrI = ebiten.NewImage(wr.X, wr.Y)
	wknI = ebiten.NewImage(wkn.X, wkn.Y)
	wkiI = ebiten.NewImage(wki.X, wki.Y)
	wqI = ebiten.NewImage(wq.X, wq.Y)
	wbI = ebiten.NewImage(wb.X, wb.Y)

	bp := blackPawnImage.Bounds().Size()
	br := blackRookImage.Bounds().Size()
	bkn := blackKnightImage.Bounds().Size()
	bki := blackKingImage.Bounds().Size()
	bq := blackQueenImage.Bounds().Size()
	bb := blackBishopImage.Bounds().Size()
	bpI = ebiten.NewImage(bp.X, bp.Y)
	brI = ebiten.NewImage(br.X, br.Y)
	bknI = ebiten.NewImage(bkn.X, bkn.Y)
	bkiI = ebiten.NewImage(bki.X, bki.Y)
	bqI = ebiten.NewImage(bq.X, bq.Y)
	bbI = ebiten.NewImage(bb.X, bb.Y)

	op := &ebiten.DrawImageOptions{}
	bI.DrawImage(boardImage, op)
	optionI.DrawImage(optionImage, op)
	lastPositionI.DrawImage(lastPositionImage, op)
	newPositionI.DrawImage(newPositionImage, op)

	wpI.DrawImage(whitePawnImage, op)
	wrI.DrawImage(whiteRookImage, op)
	wknI.DrawImage(whiteKnightImage, op)
	wkiI.DrawImage(whiteKingImage, op)
	wqI.DrawImage(whiteQueenImage, op)
	wbI.DrawImage(whiteBishopImage, op)

	bpI.DrawImage(blackPawnImage, op)
	brI.DrawImage(blackRookImage, op)
	bknI.DrawImage(blackKnightImage, op)
	bkiI.DrawImage(blackKingImage, op)
	bqI.DrawImage(blackQueenImage, op)
	bbI.DrawImage(blackBishopImage, op)
}
