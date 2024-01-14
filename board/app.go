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

	whiteBoard     map[*Piece]struct{}
	blackBoard     map[*Piece]struct{}
	whitePositions map[int]*Piece
	blackPositions map[int]*Piece

	whitesTurn bool
}

func (a *App) initWhiteBoard() {
	a.whiteBoard = make(map[*Piece]struct{})
	a.whiteBoard[&Piece{
		kind:            Rook,
		white:           true,
		currentPosition: 56,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Knight,
		white:           true,
		currentPosition: 57,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Bishop,
		white:           true,
		currentPosition: 58,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Queen,
		white:           true,
		currentPosition: 59,
	}] = value
	a.whiteBoard[&Piece{
		kind:            King,
		white:           true,
		currentPosition: 60,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Bishop,
		white:           true,
		currentPosition: 61,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Knight,
		white:           true,
		currentPosition: 62,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Rook,
		white:           true,
		currentPosition: 63,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 48,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 49,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 50,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 51,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 52,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 53,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 54,
	}] = value
	a.whiteBoard[&Piece{
		kind:            Pawn,
		white:           true,
		currentPosition: 55,
	}] = value
}

func (a *App) initBlackBoard() {
	a.blackBoard = make(map[*Piece]struct{})
	a.blackBoard[&Piece{
		kind:            Rook,
		white:           false,
		currentPosition: 0,
	}] = value
	a.blackBoard[&Piece{
		kind:            Knight,
		white:           false,
		currentPosition: 1,
	}] = value
	a.blackBoard[&Piece{
		kind:            Bishop,
		white:           false,
		currentPosition: 2,
	}] = value
	a.blackBoard[&Piece{
		kind:            Queen,
		white:           false,
		currentPosition: 3,
	}] = value
	a.blackBoard[&Piece{
		kind:            King,
		white:           false,
		currentPosition: 4,
	}] = value
	a.blackBoard[&Piece{
		kind:            Bishop,
		white:           false,
		currentPosition: 5,
	}] = value
	a.blackBoard[&Piece{
		kind:            Knight,
		white:           false,
		currentPosition: 6,
	}] = value
	a.blackBoard[&Piece{
		kind:            Rook,
		white:           false,
		currentPosition: 7,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 8,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 9,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 10,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 11,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 12,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 13,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 14,
	}] = value
	a.blackBoard[&Piece{
		kind:            Pawn,
		white:           false,
		currentPosition: 15,
	}] = value
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

func findSelectedPiece(board map[*Piece]struct{}, position int) *Piece {
	for piece := range board {
		if piece.currentPosition == position {
			return piece
		}
	}

	return nil
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

		var piece *Piece
		switch a.whitesTurn {
		case true:
			piece = findSelectedPiece(a.whiteBoard, position)
		case false:
			piece = findSelectedPiece(a.blackBoard, position)
		}

		for option := range piece.options {
			if position != option {
				continue
			}

			piece.lastPosition = piece.currentPosition
			piece.currentPosition = option
			return nil
		}

		if piece == nil {
			return nil
		}

		piece.calculateOptions(a.whiteBoard, a.blackBoard)
	}

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
