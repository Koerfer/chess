package v2

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	//a.initTestWhiteBoard() // todo remove
	//a.initTestBlackBoard() // todo remove
	a.initWhiteBoard()
	a.initBlackBoard()

	a.initImages()
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
	a.engine = v2.Engine{}
	a.engine.Init(3, true)
}
