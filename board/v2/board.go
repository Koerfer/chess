package v2

import v2 "chess/engine/v2"

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	a.initWhiteBoard()
	a.initBlackBoard()

	a.initImages()
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
	a.engine = v2.Engine{}
}
