package main

import (
	"game/raceCar/internal/entity"
	"game/raceCar/internal/infra"
)

func main() {
	displayer := infra.NewStdoutDisplayer()      //driven adapter
	game := entity.NewGame(displayer)            // game and buissness logic
	cmdInvoker := infra.NewCmdStdInInvoker(game) //driving adapter
	cmdInvoker.RunLoop()
}
