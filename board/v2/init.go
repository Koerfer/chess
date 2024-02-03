package v2

import (
	v2 "chess/pieces/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func (a *App) initWhiteBoard() {
	a.whitesTurn = true
	a.whiteBoard = make(map[int]any)
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
	a.blackBoard = make(map[int]any)
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
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
				EnPassantOptions: make(map[int]int),
			}
		case v2.PieceKindKnight:
			a.whiteBoard[pos] = &v2.Knight{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindBishop:
			a.whiteBoard[pos] = &v2.Bishop{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindRook:
			a.whiteBoard[pos] = &v2.Rook{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindQueen:
			a.whiteBoard[pos] = &v2.Queen{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindKing:
			a.whiteBoard[pos] = &v2.King{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
				CheckingPieces: make(map[int]any),
			}
		case v2.PieceKindInvalid:
			log.Fatal("invalid piece kind when initialising board")
		}
	case false:
		switch kind {
		case v2.PieceKindPawn:
			a.blackBoard[pos] = &v2.Pawn{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
				EnPassantOptions: make(map[int]int),
			}
		case v2.PieceKindKnight:
			a.blackBoard[pos] = &v2.Knight{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindBishop:
			a.blackBoard[pos] = &v2.Bishop{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindRook:
			a.blackBoard[pos] = &v2.Rook{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindQueen:
			a.blackBoard[pos] = &v2.Queen{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
			}
		case v2.PieceKindKing:
			a.blackBoard[pos] = &v2.King{
				Piece: &v2.Piece{
					White:        white,
					LastPosition: pos,
					Options:      make(map[int]struct{}),
				},
				CheckingPieces: make(map[int]any),
			}
		case v2.PieceKindInvalid:
			log.Fatal("invalid piece kind when initialising board")
		}
	}

}

func (a *App) initImages() {
	chessboard, err := os.Open("board/v2/images/chessboard.jpeg")
	if err != nil {
		log.Fatalf("unable to open chessboard image: %v", err)
	}
	jpegBoard, err := jpeg.Decode(chessboard)
	if err != nil {
		log.Fatalf("unable to decode chessboard image: %v", err)
	}
	boardImage := ebiten.NewImageFromImage(jpegBoard)

	option, err := os.Open("board/v2/images/option.png")
	if err != nil {
		log.Fatalf("unable to open option image: %v", err)
	}
	optionDecoded, err := png.Decode(option)
	if err != nil {
		log.Fatalf("unable to decode option image: %v", err)
	}
	optionImage := ebiten.NewImageFromImage(optionDecoded)

	lastPosition, err := os.Open("board/v2/images/last_position.png")
	if err != nil {
		log.Fatalf("unable to last position option image: %v", err)
	}
	lastPositionDecoded, err := png.Decode(lastPosition)
	if err != nil {
		log.Fatalf("unable to decode last position image: %v", err)
	}
	lastPositionImage := ebiten.NewImageFromImage(lastPositionDecoded)

	newPosition, err := os.Open("board/v2/images/new_position_marker.png")
	if err != nil {
		log.Fatalf("unable to last position option image: %v", err)
	}
	newPositionDecoded, err := png.Decode(newPosition)
	if err != nil {
		log.Fatalf("unable to decode last position image: %v", err)
	}
	newPositionImage := ebiten.NewImageFromImage(newPositionDecoded)

	whitePawn, err := os.Open("board/v2/images/white_pawn.png")
	if err != nil {
		log.Fatalf("unable to open White pawn image: %v", err)
	}
	whitePawnDecoded, err := png.Decode(whitePawn)
	if err != nil {
		log.Fatalf("unable to decode White pawn image: %v", err)
	}
	whitePawnImage := ebiten.NewImageFromImage(whitePawnDecoded)

	whiteKing, err := os.Open("board/v2/images/white_king.png")
	if err != nil {
		log.Fatalf("unable to open White king image: %v", err)
	}
	whiteKingDecoded, err := png.Decode(whiteKing)
	if err != nil {
		log.Fatalf("unable to decode White king image: %v", err)
	}
	whiteKingImage := ebiten.NewImageFromImage(whiteKingDecoded)

	whiteRook, err := os.Open("board/v2/images/white_rook.png")
	if err != nil {
		log.Fatalf("unable to open White rook image: %v", err)
	}
	whiteRookDecoded, err := png.Decode(whiteRook)
	if err != nil {
		log.Fatalf("unable to decode White pawn image: %v", err)
	}
	whiteRookImage := ebiten.NewImageFromImage(whiteRookDecoded)

	whiteQueen, err := os.Open("board/v2/images/white_queen.png")
	if err != nil {
		log.Fatalf("unable to open White queen image: %v", err)
	}
	whiteQueenDecoded, err := png.Decode(whiteQueen)
	if err != nil {
		log.Fatalf("unable to decode White queen image: %v", err)
	}
	whiteQueenImage := ebiten.NewImageFromImage(whiteQueenDecoded)

	whiteKnight, err := os.Open("board/v2/images/white_knight.png")
	if err != nil {
		log.Fatalf("unable to open White knight image: %v", err)
	}
	whiteKnightDecoded, err := png.Decode(whiteKnight)
	if err != nil {
		log.Fatalf("unable to decode White knight image: %v", err)
	}
	whiteKnightImage := ebiten.NewImageFromImage(whiteKnightDecoded)

	whiteBishop, err := os.Open("board/v2/images/white_bishop.png")
	if err != nil {
		log.Fatalf("unable to open White bishop image: %v", err)
	}
	whiteBishopDecoded, err := png.Decode(whiteBishop)
	if err != nil {
		log.Fatalf("unable to decode White bishop image: %v", err)
	}
	whiteBishopImage := ebiten.NewImageFromImage(whiteBishopDecoded)

	blackPawn, err := os.Open("board/v2/images/black_pawn.png")
	if err != nil {
		log.Fatalf("unable to open black pawn image: %v", err)
	}
	blackPawnDecoded, err := png.Decode(blackPawn)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackPawnImage := ebiten.NewImageFromImage(blackPawnDecoded)

	blackKing, err := os.Open("board/v2/images/black_king.png")
	if err != nil {
		log.Fatalf("unable to open black king image: %v", err)
	}
	blackKingDecoded, err := png.Decode(blackKing)
	if err != nil {
		log.Fatalf("unable to decode black king image: %v", err)
	}
	blackKingImage := ebiten.NewImageFromImage(blackKingDecoded)

	blackRook, err := os.Open("board/v2/images/black_rook.png")
	if err != nil {
		log.Fatalf("unable to open black rook image: %v", err)
	}
	blackRookDecoded, err := png.Decode(blackRook)
	if err != nil {
		log.Fatalf("unable to decode black pawn image: %v", err)
	}
	blackRookImage := ebiten.NewImageFromImage(blackRookDecoded)

	blackQueen, err := os.Open("board/v2/images/black_queen.png")
	if err != nil {
		log.Fatalf("unable to open black queen image: %v", err)
	}
	blackQueenDecoded, err := png.Decode(blackQueen)
	if err != nil {
		log.Fatalf("unable to decode black queen image: %v", err)
	}
	blackQueenImage := ebiten.NewImageFromImage(blackQueenDecoded)

	blackKnight, err := os.Open("board/v2/images/black_knight.png")
	if err != nil {
		log.Fatalf("unable to open black knight image: %v", err)
	}
	blackKnightDecoded, err := png.Decode(blackKnight)
	if err != nil {
		log.Fatalf("unable to decode black knight image: %v", err)
	}
	blackKnightImage := ebiten.NewImageFromImage(blackKnightDecoded)

	blackBishop, err := os.Open("board/v2/images/black_bishop.png")
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
