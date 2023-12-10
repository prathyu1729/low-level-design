package battleship_game

type firingStrategy interface {
	getPos(boundary boundary) Point
}
