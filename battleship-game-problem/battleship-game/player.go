package battleship_game

import (
	"errors"
	"fmt"
)

type player struct {
	ships         []ship
	boundary      boundary
	shipCount     int
	liveShipCount int
	shipPoints    map[Point]int
}

func newPlayer(boundary boundary) *player {
	return &player{
		ships:         []ship{},
		boundary:      boundary,
		shipCount:     0,
		liveShipCount: 0,
		shipPoints:    map[Point]int{},
	}
}

func (p *player) addShip(point Point, size int) error {
	x1 := point.x + size
	y1 := point.y + size
	if !p.boundary.contains(point) || !p.boundary.contains(NewPoint(x1, y1)) {
		return errors.New("invalid input")
	}
	newShip := newShip(point, size)
	p.ships = append(p.ships, newShip)

	p.incShipCount()
	p.addShipPoints(point, size)
	return nil
}

func (p *player) incShipCount() {
	p.shipCount = p.shipCount + 1
	p.liveShipCount = p.liveShipCount + 1
}

func (p *player) addShipPoints(point Point, size int) {
	x := point.x
	y := point.y
	for i := x; i <= x+size; i++ {
		for j := y; j <= y+size; j++ {
			newPoint := NewPoint(i, j)
			p.shipPoints[newPoint] = p.shipCount - 1
		}
	}

}

func (p *player) hasFailed() bool {
	return p.shipCount > 0 && p.liveShipCount == 0
}

func (p *player) destroyShip(point Point) {
	shipId, ok := p.shipPoints[point]
	if !ok {
		fmt.Println("Missed Bombing")
		return
	}

	p.ships[shipId].setState(dead)
	p.liveShipCount = p.liveShipCount - 1
	fmt.Println("Killed ship", shipId)

}
