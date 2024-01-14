package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	ScreenWidth  = 870
	ScreenHeight = 849
)

var (
	bI      *ebiten.Image
	optionI *ebiten.Image
	wpI     *ebiten.Image
	wbI     *ebiten.Image
	wkiI    *ebiten.Image
	wqI     *ebiten.Image
	wrI     *ebiten.Image
	wknI    *ebiten.Image
	bpI     *ebiten.Image
	bbI     *ebiten.Image
	bkiI    *ebiten.Image
	bqI     *ebiten.Image
	brI     *ebiten.Image
	bknI    *ebiten.Image
)

type App struct {
	touchIDs  []ebiten.TouchID
	op        ebiten.DrawImageOptions
	initiated bool

	whiteBoard map[int]*Piece
	blackBoard map[int]*Piece

	whitesTurn    bool
	selectedPiece *Piece
}

func (a *App) initWhiteBoard() {
	a.whiteBoard = make(map[int]*Piece)
	a.whiteBoard[56] = &Piece{
		kind:    Rook,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[57] = &Piece{
		kind:    Knight,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[58] = &Piece{
		kind:    Bishop,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[59] = &Piece{
		kind:    Queen,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[60] = &Piece{
		kind:    King,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[61] = &Piece{
		kind:    Bishop,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[62] = &Piece{
		kind:    Knight,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[63] = &Piece{
		kind:    Rook,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[48] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[49] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[50] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[51] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[52] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[53] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[54] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
	a.whiteBoard[55] = &Piece{
		kind:    Pawn,
		white:   true,
		options: make(map[int]struct{}),
	}
}

func (a *App) initBlackBoard() {
	a.blackBoard = make(map[int]*Piece)
	a.blackBoard[0] = &Piece{
		kind:    Rook,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[1] = &Piece{
		kind:    Knight,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[2] = &Piece{
		kind:    Bishop,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[3] = &Piece{
		kind:    Queen,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[4] = &Piece{
		kind:    King,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[5] = &Piece{
		kind:    Bishop,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[6] = &Piece{
		kind:    Knight,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[7] = &Piece{
		kind:    Rook,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[8] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[9] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[10] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[11] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[12] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[13] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[14] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
	a.blackBoard[15] = &Piece{
		kind:    Pawn,
		white:   false,
		options: make(map[int]struct{}),
	}
}

func (a *App) initImages() {
	chessboard, err := os.Open("board/chessboard.jpeg")
	if err != nil {
		log.Fatalf("unable to open chessboard image: %v", err)
	}
	jpegBoard, err := jpeg.Decode(chessboard)
	if err != nil {
		log.Fatalf("unable to decode chessboard image: %v", err)
	}
	boardImage := ebiten.NewImageFromImage(jpegBoard)

	option, err := os.Open("board/option.png")
	if err != nil {
		log.Fatalf("unable to open option image: %v", err)
	}
	optionDecoded, err := png.Decode(option)
	if err != nil {
		log.Fatalf("unable to decode option image: %v", err)
	}
	optionImage := ebiten.NewImageFromImage(optionDecoded)

	whitePawn, err := os.Open("board/white_pawn.png")
	if err != nil {
		log.Fatalf("unable to open white pawn image: %v", err)
	}
	whitePawnDecoded, err := png.Decode(whitePawn)
	if err != nil {
		log.Fatalf("unable to decode white pawn image: %v", err)
	}
	whitePawnImage := ebiten.NewImageFromImage(whitePawnDecoded)

	whiteKing, err := os.Open("board/white_king.png")
	if err != nil {
		log.Fatalf("unable to open white king image: %v", err)
	}
	whiteKingDecoded, err := png.Decode(whiteKing)
	if err != nil {
		log.Fatalf("unable to decode white king image: %v", err)
	}
	whiteKingImage := ebiten.NewImageFromImage(whiteKingDecoded)

	whiteRook, err := os.Open("board/white_rook.png")
	if err != nil {
		log.Fatalf("unable to open white rook image: %v", err)
	}
	whiteRookDecoded, err := png.Decode(whiteRook)
	if err != nil {
		log.Fatalf("unable to decode white pawn image: %v", err)
	}
	whiteRookImage := ebiten.NewImageFromImage(whiteRookDecoded)

	whiteQueen, err := os.Open("board/white_queen.png")
	if err != nil {
		log.Fatalf("unable to open white queen image: %v", err)
	}
	whiteQueenDecoded, err := png.Decode(whiteQueen)
	if err != nil {
		log.Fatalf("unable to decode white queen image: %v", err)
	}
	whiteQueenImage := ebiten.NewImageFromImage(whiteQueenDecoded)

	whiteKnight, err := os.Open("board/white_knight.png")
	if err != nil {
		log.Fatalf("unable to open white knight image: %v", err)
	}
	whiteKnightDecoded, err := png.Decode(whiteKnight)
	if err != nil {
		log.Fatalf("unable to decode white knight image: %v", err)
	}
	whiteKnightImage := ebiten.NewImageFromImage(whiteKnightDecoded)

	whiteBishop, err := os.Open("board/white_bishop.png")
	if err != nil {
		log.Fatalf("unable to open white bishop image: %v", err)
	}
	whiteBishopDecoded, err := png.Decode(whiteBishop)
	if err != nil {
		log.Fatalf("unable to decode white bishop image: %v", err)
	}
	whiteBishopImage := ebiten.NewImageFromImage(whiteBishopDecoded)

	blackPawn, err := os.Open("board/black_pawn.png")
	if err != nil {
		log.Fatalf("unable to open black pawn image: %v", err)
	}
	blackPawnDecoded, err := png.Decode(blackPawn)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackPawnImage := ebiten.NewImageFromImage(blackPawnDecoded)

	blackKing, err := os.Open("board/black_king.png")
	if err != nil {
		log.Fatalf("unable to open black king image: %v", err)
	}
	blackKingDecoded, err := png.Decode(blackKing)
	if err != nil {
		log.Fatalf("unable to decode black king image: %v", err)
	}
	blackKingImage := ebiten.NewImageFromImage(blackKingDecoded)

	blackRook, err := os.Open("board/black_rook.png")
	if err != nil {
		log.Fatalf("unable to open black rook image: %v", err)
	}
	blackRookDecoded, err := png.Decode(blackRook)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackRookImage := ebiten.NewImageFromImage(blackRookDecoded)

	blackQueen, err := os.Open("board/black_queen.png")
	if err != nil {
		log.Fatalf("unable to open black queen image: %v", err)
	}
	blackQueenDecoded, err := png.Decode(blackQueen)
	if err != nil {
		log.Fatalf("unable to decode black queen image: %v", err)
	}
	blackQueenImage := ebiten.NewImageFromImage(blackQueenDecoded)

	blackKnight, err := os.Open("board/black_knight.png")
	if err != nil {
		log.Fatalf("unable to open black knight image: %v", err)
	}
	blackKnightDecoded, err := png.Decode(blackKnight)
	if err != nil {
		log.Fatalf("unable to decode black knight image: %v", err)
	}
	blackKnightImage := ebiten.NewImageFromImage(blackKnightDecoded)

	blackBishop, err := os.Open("board/black_bishop.png")
	if err != nil {
		log.Fatalf("unable to open black bishop image: %v", err)
	}
	blackBishopDecoded, err := png.Decode(blackBishop)
	if err != nil {
		log.Fatalf("unable to decode black bishop image: %v", err)
	}
	blackBishopImage := ebiten.NewImageFromImage(blackBishopDecoded)

	b := boardImage.Bounds().Size()
	bI = ebiten.NewImage(b.X, b.Y)
	o := optionImage.Bounds().Size()
	optionI = ebiten.NewImage(o.X, o.Y)

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

func (a *App) initTheRest() {
	a.initWhiteBoard()
	a.initBlackBoard()

	a.whitesTurn = true
	a.initImages()
}

func (a *App) Update() error {
	if !a.initiated {
		a.init()
	}
	a.touchIDs = ebiten.AppendTouchIDs(a.touchIDs[:0])

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		X := int(math.Floor(float64(x) / 108.75))
		Y := int(math.Floor(float64(y) / 106.125))
		position := X + Y*8
		if position >= 64 {
			return nil
		}

		var board map[int]*Piece
		switch a.whitesTurn {
		case true:
			board = a.whiteBoard
		case false:
			board = a.blackBoard
		}

		if piece, ok := board[position]; ok {
			piece.calculateOptions(a.whiteBoard, a.blackBoard, position)
			piece.lastPosition = position
			a.selectedPiece = piece
		}

		if a.selectedPiece == nil {
			return nil
		}

		for option := range a.selectedPiece.options {
			if position != option {
				continue
			}

			board[option] = a.selectedPiece
			delete(board, a.selectedPiece.lastPosition)
			a.selectedPiece = nil
			return nil
		}
	}

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
