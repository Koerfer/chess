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
		Kind:         pieces.Rook,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 56,
	}
	a.whiteBoard[57] = &pieces.Piece{
		Kind:         pieces.Knight,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 57,
	}
	a.whiteBoard[58] = &pieces.Piece{
		Kind:         pieces.Bishop,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 58,
	}
	a.whiteBoard[59] = &pieces.Piece{
		Kind:         pieces.Queen,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 59,
	}
	a.whiteBoard[60] = &pieces.Piece{
		Kind:           pieces.King,
		White:          true,
		Options:        make(map[int]struct{}),
		CheckingPieces: make(map[int]*pieces.Piece),
		LastPosition:   60,
	}
	a.whiteBoard[61] = &pieces.Piece{
		Kind:         pieces.Bishop,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 61,
	}
	a.whiteBoard[62] = &pieces.Piece{
		Kind:         pieces.Knight,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 62,
	}
	a.whiteBoard[63] = &pieces.Piece{
		Kind:         pieces.Rook,
		White:        true,
		Options:      make(map[int]struct{}),
		LastPosition: 63,
	}
	a.whiteBoard[48] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     8,
	}
	a.whiteBoard[49] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     49,
	}
	a.whiteBoard[50] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     50,
	}
	a.whiteBoard[35] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     35,
	}
	a.whiteBoard[36] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     36,
	}
	a.whiteBoard[53] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     53,
	}
	a.whiteBoard[54] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     54,
	}
	a.whiteBoard[55] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            true,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     55,
	}
}

func (a *App) initBlackBoard() {
	a.blackBoard = make(map[int]*pieces.Piece)
	a.blackBoard[0] = &pieces.Piece{
		Kind:         pieces.Rook,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 0,
	}
	a.blackBoard[1] = &pieces.Piece{
		Kind:         pieces.Knight,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 1,
	}
	a.blackBoard[2] = &pieces.Piece{
		Kind:         pieces.Bishop,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 2,
	}
	a.blackBoard[3] = &pieces.Piece{
		Kind:         pieces.Queen,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 3,
	}
	a.blackBoard[4] = &pieces.Piece{
		Kind:           pieces.King,
		White:          false,
		Options:        make(map[int]struct{}),
		CheckingPieces: make(map[int]*pieces.Piece),
		LastPosition:   4,
	}
	a.blackBoard[5] = &pieces.Piece{
		Kind:         pieces.Bishop,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 5,
	}
	a.blackBoard[6] = &pieces.Piece{
		Kind:         pieces.Knight,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 6,
	}
	a.blackBoard[7] = &pieces.Piece{
		Kind:         pieces.Rook,
		White:        false,
		Options:      make(map[int]struct{}),
		LastPosition: 7,
	}
	a.blackBoard[8] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     8,
	}
	a.blackBoard[9] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     9,
	}
	a.blackBoard[18] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     18,
	}
	a.blackBoard[11] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     11,
	}
	a.blackBoard[20] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     20,
	}
	a.blackBoard[13] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     13,
	}
	a.blackBoard[14] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     14,
	}
	a.blackBoard[15] = &pieces.Piece{
		Kind:             pieces.Pawn,
		White:            false,
		Options:          make(map[int]struct{}),
		EnPassantOptions: make(map[int]int),
		LastPosition:     15,
	}
}

//func (a *App) initWhiteBoard() {
//	a.whitesTurn = true
//	a.whiteBoard = make(map[int]*pieces.Piece)
//	a.whiteBoard[56] = &pieces.Piece{
//		Kind:         pieces.Rook,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 56,
//	}
//	a.whiteBoard[57] = &pieces.Piece{
//		Kind:         pieces.Knight,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 57,
//	}
//	a.whiteBoard[58] = &pieces.Piece{
//		Kind:         pieces.Bishop,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 58,
//	}
//	a.whiteBoard[59] = &pieces.Piece{
//		Kind:         pieces.Queen,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 59,
//	}
//	a.whiteBoard[60] = &pieces.Piece{
//		Kind:           pieces.King,
//		White:          true,
//		Options:        make(map[int]struct{}),
//		CheckingPieces: make(map[int]*pieces.Piece),
//		LastPosition:   60,
//	}
//	a.whiteBoard[61] = &pieces.Piece{
//		Kind:         pieces.Bishop,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 61,
//	}
//	a.whiteBoard[62] = &pieces.Piece{
//		Kind:         pieces.Knight,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 62,
//	}
//	a.whiteBoard[63] = &pieces.Piece{
//		Kind:         pieces.Rook,
//		White:        true,
//		Options:      make(map[int]struct{}),
//		LastPosition: 63,
//	}
//	a.whiteBoard[48] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     8,
//	}
//	a.whiteBoard[49] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     49,
//	}
//	a.whiteBoard[50] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     50,
//	}
//	a.whiteBoard[51] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     51,
//	}
//	a.whiteBoard[52] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     52,
//	}
//	a.whiteBoard[53] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     53,
//	}
//	a.whiteBoard[54] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     54,
//	}
//	a.whiteBoard[55] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            true,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     55,
//	}
//}
//
//func (a *App) initBlackBoard() {
//	a.blackBoard = make(map[int]*pieces.Piece)
//	a.blackBoard[0] = &pieces.Piece{
//		Kind:         pieces.Rook,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 0,
//	}
//	a.blackBoard[1] = &pieces.Piece{
//		Kind:         pieces.Knight,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 1,
//	}
//	a.blackBoard[2] = &pieces.Piece{
//		Kind:         pieces.Bishop,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 2,
//	}
//	a.blackBoard[3] = &pieces.Piece{
//		Kind:         pieces.Queen,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 3,
//	}
//	a.blackBoard[4] = &pieces.Piece{
//		Kind:           pieces.King,
//		White:          false,
//		Options:        make(map[int]struct{}),
//		CheckingPieces: make(map[int]*pieces.Piece),
//		LastPosition:   4,
//	}
//	a.blackBoard[5] = &pieces.Piece{
//		Kind:         pieces.Bishop,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 5,
//	}
//	a.blackBoard[6] = &pieces.Piece{
//		Kind:         pieces.Knight,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 6,
//	}
//	a.blackBoard[7] = &pieces.Piece{
//		Kind:         pieces.Rook,
//		White:        false,
//		Options:      make(map[int]struct{}),
//		LastPosition: 7,
//	}
//	a.blackBoard[8] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     8,
//	}
//	a.blackBoard[9] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     9,
//	}
//	a.blackBoard[10] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     10,
//	}
//	a.blackBoard[11] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     11,
//	}
//	a.blackBoard[12] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     12,
//	}
//	a.blackBoard[13] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     13,
//	}
//	a.blackBoard[14] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     14,
//	}
//	a.blackBoard[15] = &pieces.Piece{
//		Kind:             pieces.Pawn,
//		White:            false,
//		Options:          make(map[int]struct{}),
//		EnPassantOptions: make(map[int]int),
//		LastPosition:     15,
//	}
//}

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

	lastPosition, err := os.Open("board/images/last_position.png")
	if err != nil {
		log.Fatalf("unable to last position option image: %v", err)
	}
	lastPositionDecoded, err := png.Decode(lastPosition)
	if err != nil {
		log.Fatalf("unable to decode last position image: %v", err)
	}
	lastPositionImage := ebiten.NewImageFromImage(lastPositionDecoded)

	newPosition, err := os.Open("board/images/new_position_marker.png")
	if err != nil {
		log.Fatalf("unable to last position option image: %v", err)
	}
	newPositionDecoded, err := png.Decode(newPosition)
	if err != nil {
		log.Fatalf("unable to decode last position image: %v", err)
	}
	newPositionImage := ebiten.NewImageFromImage(newPositionDecoded)

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
