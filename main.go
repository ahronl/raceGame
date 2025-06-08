package main

import (
	"example/hello/internal/entity"
	"example/hello/internal/infra"
)

func main() {
	displayer := infra.NewStdoutDisplayer()      //driven adapter
	game := entity.NewGame(displayer)            // game and buissness logic
	cmdInvoker := infra.NewCmdStdInInvoker(game) //driving adapter
	cmdInvoker.RunLoop()
}
