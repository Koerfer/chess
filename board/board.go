package board

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	a.initTheRest()
}
