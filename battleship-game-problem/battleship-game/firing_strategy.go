package battleship_game

type firingStrategy interface {
	getPos(xlow, xhigh, ylow, yhigh int) (int, int)
}
