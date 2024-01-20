package board

import (
	"chess/pieces"
	"github.com/hajimehoshi/ebiten/v2"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func (a *App) initWhiteBoard() {
	a.whitesTurn = true
	a.whiteBoard = make(map[int]*pieces.Piece)
	a.whiteBoard[56] = &pieces.Piece{
		Kind:    pieces.Rook,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[57] = &pieces.Piece{
		Kind:    pieces.Knight,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[58] = &pieces.Piece{
		Kind:    pieces.Bishop,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[59] = &pieces.Piece{
		Kind:    pieces.Queen,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[60] = &pieces.Piece{
		Kind:           pieces.King,
		White:          true,
		Options:        make(map[int]struct{}),
		CheckingPieces: make(map[int]*pieces.Piece),
	}
	a.whiteBoard[61] = &pieces.Piece{
		Kind:    pieces.Bishop,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[62] = &pieces.Piece{
		Kind:    pieces.Knight,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[63] = &pieces.Piece{
		Kind:    pieces.Rook,
		White:   true,
		Options: make(map[int]struct{}),
	}
	a.whiteBoard[48] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[49] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[50] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[51] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[52] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[53] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[54] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.whiteBoard[55] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
}

func (a *App) initBlackBoard() {
	a.blackBoard = make(map[int]*pieces.Piece)
	a.blackBoard[0] = &pieces.Piece{
		Kind:    pieces.Rook,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[1] = &pieces.Piece{
		Kind:    pieces.Knight,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[2] = &pieces.Piece{
		Kind:    pieces.Bishop,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[3] = &pieces.Piece{
		Kind:    pieces.Queen,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[4] = &pieces.Piece{
		Kind:           pieces.King,
		White:          false,
		Options:        make(map[int]struct{}),
		CheckingPieces: make(map[int]*pieces.Piece),
	}
	a.blackBoard[5] = &pieces.Piece{
		Kind:    pieces.Bishop,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[6] = &pieces.Piece{
		Kind:    pieces.Knight,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[7] = &pieces.Piece{
		Kind:    pieces.Rook,
		White:   false,
		Options: make(map[int]struct{}),
	}
	a.blackBoard[8] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[9] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[10] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[11] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[12] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[13] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[14] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
	a.blackBoard[15] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
	}
}

func (a *App) initImages() {
	chessboard, err := os.Open("board/images/chessboard.jpeg")
	if err != nil {
		log.Fatalf("unable to open chessboard image: %v", err)
	}
	jpegBoard, err := jpeg.Decode(chessboard)
	if err != nil {
		log.Fatalf("unable to decode chessboard image: %v", err)
	}
	boardImage := ebiten.NewImageFromImage(jpegBoard)

	option, err := os.Open("board/images/option.png")
	if err != nil {
		log.Fatalf("unable to open option image: %v", err)
	}
	optionDecoded, err := png.Decode(option)
	if err != nil {
		log.Fatalf("unable to decode option image: %v", err)
	}
	optionImage := ebiten.NewImageFromImage(optionDecoded)

	whitePawn, err := os.Open("board/images/white_pawn.png")
	if err != nil {
		log.Fatalf("unable to open White pawn image: %v", err)
	}
	whitePawnDecoded, err := png.Decode(whitePawn)
	if err != nil {
		log.Fatalf("unable to decode White pawn image: %v", err)
	}
	whitePawnImage := ebiten.NewImageFromImage(whitePawnDecoded)

	whiteKing, err := os.Open("board/images/white_king.png")
	if err != nil {
		log.Fatalf("unable to open White king image: %v", err)
	}
	whiteKingDecoded, err := png.Decode(whiteKing)
	if err != nil {
		log.Fatalf("unable to decode White king image: %v", err)
	}
	whiteKingImage := ebiten.NewImageFromImage(whiteKingDecoded)

	whiteRook, err := os.Open("board/images/white_rook.png")
	if err != nil {
		log.Fatalf("unable to open White rook image: %v", err)
	}
	whiteRookDecoded, err := png.Decode(whiteRook)
	if err != nil {
		log.Fatalf("unable to decode White pawn image: %v", err)
	}
	whiteRookImage := ebiten.NewImageFromImage(whiteRookDecoded)

	whiteQueen, err := os.Open("board/images/white_queen.png")
	if err != nil {
		log.Fatalf("unable to open White queen image: %v", err)
	}
	whiteQueenDecoded, err := png.Decode(whiteQueen)
	if err != nil {
		log.Fatalf("unable to decode White queen image: %v", err)
	}
	whiteQueenImage := ebiten.NewImageFromImage(whiteQueenDecoded)

	whiteKnight, err := os.Open("board/images/white_knight.png")
	if err != nil {
		log.Fatalf("unable to open White knight image: %v", err)
	}
	whiteKnightDecoded, err := png.Decode(whiteKnight)
	if err != nil {
		log.Fatalf("unable to decode White knight image: %v", err)
	}
	whiteKnightImage := ebiten.NewImageFromImage(whiteKnightDecoded)

	whiteBishop, err := os.Open("board/images/white_bishop.png")
	if err != nil {
		log.Fatalf("unable to open White bishop image: %v", err)
	}
	whiteBishopDecoded, err := png.Decode(whiteBishop)
	if err != nil {
		log.Fatalf("unable to decode White bishop image: %v", err)
	}
	whiteBishopImage := ebiten.NewImageFromImage(whiteBishopDecoded)

	blackPawn, err := os.Open("board/images/black_pawn.png")
	if err != nil {
		log.Fatalf("unable to open black pawn image: %v", err)
	}
	blackPawnDecoded, err := png.Decode(blackPawn)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackPawnImage := ebiten.NewImageFromImage(blackPawnDecoded)

	blackKing, err := os.Open("board/images/black_king.png")
	if err != nil {
		log.Fatalf("unable to open black king image: %v", err)
	}
	blackKingDecoded, err := png.Decode(blackKing)
	if err != nil {
		log.Fatalf("unable to decode black king image: %v", err)
	}
	blackKingImage := ebiten.NewImageFromImage(blackKingDecoded)

	blackRook, err := os.Open("board/images/black_rook.png")
	if err != nil {
		log.Fatalf("unable to open black rook image: %v", err)
	}
	blackRookDecoded, err := png.Decode(blackRook)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackRookImage := ebiten.NewImageFromImage(blackRookDecoded)

	blackQueen, err := os.Open("board/images/black_queen.png")
	if err != nil {
		log.Fatalf("unable to open black queen image: %v", err)
	}
	blackQueenDecoded, err := png.Decode(blackQueen)
	if err != nil {
		log.Fatalf("unable to decode black queen image: %v", err)
	}
	blackQueenImage := ebiten.NewImageFromImage(blackQueenDecoded)

	blackKnight, err := os.Open("board/images/black_knight.png")
	if err != nil {
		log.Fatalf("unable to open black knight image: %v", err)
	}
	blackKnightDecoded, err := png.Decode(blackKnight)
	if err != nil {
		log.Fatalf("unable to decode black knight image: %v", err)
	}
	blackKnightImage := ebiten.NewImageFromImage(blackKnightDecoded)

	blackBishop, err := os.Open("board/images/black_bishop.png")
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
