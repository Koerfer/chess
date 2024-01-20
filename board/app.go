package board

import (
	"chess/pieces"
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

	whiteBoard map[int]*pieces.Piece
	blackBoard map[int]*pieces.Piece

	whitesTurn    bool
	selectedPiece *pieces.Piece
}

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
		Kind:    pieces.King,
		White:   true,
		Options: make(map[int]struct{}),
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
		Kind:    pieces.King,
		White:   false,
		Options: make(map[int]struct{}),
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

		var board map[int]*pieces.Piece
		switch a.whitesTurn {
		case true:
			board = a.whiteBoard
		case false:
			board = a.blackBoard
		}

		if piece, ok := board[position]; ok {
			piece.LastPosition = position
			a.selectedPiece = piece
		}

		if a.selectedPiece == nil {
			return nil
		}

		for option, take := range a.selectedPiece.EnPassantOptions {
			if position != option {
				continue
			}

			switch a.whitesTurn {
			case true:
				delete(a.blackBoard, take)
			case false:
				delete(a.whiteBoard, take)
			}

			board[option] = a.selectedPiece
			delete(board, a.selectedPiece.LastPosition)
			a.selectedPiece = nil

			a.whitesTurn = !a.whitesTurn
			a.calculateAllPositions(a.whiteBoard, a.blackBoard)
			return nil
		}

		for option := range a.selectedPiece.Options {
			if position != option {
				continue
			}

			a.TakeOrPromote(position)

			if a.selectedPiece.Kind == pieces.King && !a.selectedPiece.HasBeenMoved {
				castled := a.castle(option, board)
				if castled {
					return nil
				}
			}

			if a.selectedPiece.Kind == pieces.King || a.selectedPiece.Kind == pieces.Rook {
				a.selectedPiece.HasBeenMoved = true
			}

			board[option] = a.selectedPiece
			delete(board, a.selectedPiece.LastPosition)
			a.selectedPiece = nil

			a.whitesTurn = !a.whitesTurn
			a.calculateAllPositions(a.whiteBoard, a.blackBoard)
			return nil
		}
	}

	return nil
}

func (a *App) castle(option int, board map[int]*pieces.Piece) bool {
	switch option {
	case 2:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[3] = board[0]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 0)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 6:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[5] = board[7]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 7)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 58:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[59] = board[56]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 56)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	case 62:
		a.selectedPiece.HasBeenMoved = true
		board[option] = a.selectedPiece
		board[61] = board[63]
		delete(board, a.selectedPiece.LastPosition)
		delete(board, 63)
		a.selectedPiece = nil
		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) calculateAllPositions(whiteBoard map[int]*pieces.Piece, blackBoard map[int]*pieces.Piece) {
	forbiddenSquares := make(map[int]struct{})

	switch a.whitesTurn {
	case true:
		for position, piece := range blackBoard {
			forbiddenCaptures, _ := piece.CalculateOptions(whiteBoard, blackBoard, position, nil, false)
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if piece.Kind != pieces.Pawn {
				for forbidden := range piece.Options {
					forbiddenSquares[forbidden] = struct{}{}
				}
			}
		}

		for position, piece := range whiteBoard {
			piece.CalculateOptions(whiteBoard, blackBoard, position, forbiddenSquares, true)
		}
	case false:
		for position, piece := range whiteBoard {
			forbiddenCaptures, _ := piece.CalculateOptions(whiteBoard, blackBoard, position, nil, false)
			for forbidden := range forbiddenCaptures {
				forbiddenSquares[forbidden] = struct{}{}
			}
			if piece.Kind != pieces.Pawn {
				for forbidden := range piece.Options {
					forbiddenSquares[forbidden] = struct{}{}
				}
			}
		}

		for position, piece := range blackBoard {
			piece.CalculateOptions(whiteBoard, blackBoard, position, forbiddenSquares, true)
		}
	}

}

func (a *App) TakeOrPromote(position int) {
	if a.selectedPiece.Kind == pieces.Pawn {
		end := 7
		if a.selectedPiece.White == true {
			end = 0
		}
		if position/8 == end {
			a.selectedPiece.Kind = pieces.Queen // todo: add convert to better Piece logic
		}
	}

	switch a.whitesTurn {
	case true:
		delete(a.blackBoard, position)
	case false:
		delete(a.whiteBoard, position)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
