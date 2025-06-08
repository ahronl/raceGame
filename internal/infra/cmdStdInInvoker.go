package infra

import (
	"example/hello/internal/entity"
	"fmt"
)

type stdIn struct {
	game entity.Game
}

func NewCmdStdInInvoker(game entity.Game) *stdIn {
	return &stdIn{
		game: game,
	}
}

func (s *stdIn) RunLoop() {
	for {
		var cmd string
		fmt.Print("Enter a command: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case "exit":
			return
		case "track":
			s.game.SetRoad()
		case "race":
			s.game.RaceCars()
		case "alfa":
			s.game.BuildAlfa()
		case "bmw":
			s.game.BuildBMW()
		default:
		}
	}
}
