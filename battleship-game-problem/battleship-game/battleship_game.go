package battleship_game

import (
	"fmt"
	"sync"
)

type BattleShipGame struct {
	N               int
	ships           []ship
	players         []*player
	firingStrategy  firingStrategy
	currentPlayerId int
}

var once sync.Once

var game *BattleShipGame

func InitGame(n int) *BattleShipGame {
	once.Do(func() {
		game = &BattleShipGame{
			N:               n,
			ships:           []ship{},
			firingStrategy:  newRandomFiringStrategy(),
			currentPlayerId: 0,
		}
	})
	game.setPlayers()

	return game
}

func (b *BattleShipGame) AddShip(playerId int, point Point, size int) error {
	return b.players[playerId].addShip(point, size)
}

func (b *BattleShipGame) StartGame() {
	fmt.Println("Game Started")

	for !b.hasGameEnded() {
		fmt.Println("****************")
		fmt.Println(b.currentPlayerId, "shooting")
		oppPlayer := b.getOppositionPlayer()
		b.shoot(oppPlayer)
		b.setCurrentPlayerId()
		fmt.Println("****************")
	}

	fmt.Println("Game Ended")
}

func (b *BattleShipGame) setPlayers() {
	boundary1 := boundary{
		topLeft:     NewPoint(0, 0),
		bottomRight: NewPoint(b.N, b.N/2),
	}

	player1 := newPlayer(boundary1)

	boundary2 := boundary{
		topLeft:     NewPoint(0, b.N/2),
		bottomRight: NewPoint(b.N, b.N),
	}

	player2 := newPlayer(boundary2)

	b.players = []*player{player1, player2}
}

func (b *BattleShipGame) setCurrentPlayerId() {
	if b.currentPlayerId == 0 {
		b.currentPlayerId = 1
	} else {
		b.currentPlayerId = 0
	}
}

func (b *BattleShipGame) hasGameEnded() bool {
	for i, player := range b.players {
		if player.hasFailed() {
			fmt.Println("Player lost:", i)
			return true
		}
	}
	return false
}

func (b *BattleShipGame) shoot(player *player) {
	shootingPoint := b.firingStrategy.getPos(player.boundary)
	fmt.Println("shooting location:", shootingPoint.x, shootingPoint.y)
	player.destroyShip(shootingPoint)
}

func (b *BattleShipGame) getOppositionPlayer() *player {
	if b.currentPlayerId == 0 {
		return b.players[1]
	}

	return b.players[0]
}
