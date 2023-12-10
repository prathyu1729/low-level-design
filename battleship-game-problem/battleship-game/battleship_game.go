package battleship_game

import "sync"

type BattleShipGame struct {
	N               int
	ships           []ship
	players         map[int]player
	firingStrategy  firingStrategy
	currentPlayerId int
}

var once sync.Once

var game *BattleShipGame

func InitGame(n int) *BattleShipGame {
	once.Do(func() {
		game = &BattleShipGame{
			N:              n,
			ships:          []ship{},
			players:        initPlayers(),
			firingStrategy: newRandomFiringStrategy(),
		}
	})

	game.currentPlayerId = game.getCurrentPlayerId()
	return game
}

func initPlayers() map[int]player {
	panic("implement me")
}

func (b *BattleShipGame) getCurrentPlayerId() int {
	panic("implement me")
}

