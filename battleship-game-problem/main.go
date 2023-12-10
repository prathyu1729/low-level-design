package main

import (
	battleship_game "battleship-game/battleship-game"
	"fmt"
)

func main() {

	game := battleship_game.InitGame(4)
	if err := game.AddShip(0, battleship_game.NewPoint(0, 0), 2); err != nil {
		fmt.Println(err.Error())
	}

	if err := game.AddShip(0, battleship_game.NewPoint(2, 0), 2); err != nil {
		fmt.Println(err.Error())
	}

	if err := game.AddShip(1, battleship_game.NewPoint(0, 2), 1); err != nil {
		fmt.Println(err.Error())
	}

	game.StartGame()

}
