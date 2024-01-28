package board

import "chess/enginev2"

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	a.initWhiteBoard()
	a.initBlackBoard()
	a.initImages()
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
	a.engine = enginev2.Engine{}
	a.engine.Init(3)
}
