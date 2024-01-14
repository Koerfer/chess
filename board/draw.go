package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func (a *App) drawOptions(screen *ebiten.Image, piece *Piece) {
	a.op.GeoM.Reset()

	for option := range piece.options {
		a.op.GeoM.Reset()

		a.op.GeoM.Translate(108.75*float64(option%8), 106.125*math.Floor(float64(option/8)))
		screen.DrawImage(optionI, &a.op)
	}

}

func (a *App) Draw(screen *ebiten.Image) {
	a.op.GeoM.Reset()

	screen.DrawImage(bI, &a.op)
	a.drawPieces(screen)
	if a.selectedPiece != nil {
		a.drawOptions(screen, a.selectedPiece)
	}
}

func (a *App) drawPieces(screen *ebiten.Image) {
	a.op.GeoM.Reset()

	for position, piece := range a.whiteBoard {
		a.op.GeoM.Reset()
		switch piece.kind {
		case Pawn:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wpI, &a.op)
		case Knight:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wknI, &a.op)
		case Bishop:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wbI, &a.op)
		case Rook:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wrI, &a.op)
		case Queen:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wqI, &a.op)
		case King:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(wkiI, &a.op)
		}
	}

	for position, piece := range a.blackBoard {
		a.op.GeoM.Reset()
		switch piece.kind {
		case Pawn:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(bpI, &a.op)
		case Knight:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(bknI, &a.op)
		case Bishop:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(bbI, &a.op)
		case Rook:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(brI, &a.op)
		case Queen:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(bqI, &a.op)
		case King:
			a.op.GeoM.Translate(108.75*float64(position%8), 106.125*math.Floor(float64(position/8)))
			screen.DrawImage(bkiI, &a.op)
		}
	}
}
