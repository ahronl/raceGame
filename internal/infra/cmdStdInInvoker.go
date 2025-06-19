package infra

import (
	"fmt"
	"game/raceCar/internal/entity"
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
			s.game.SetRoad("track")
		case "highway":
			s.game.SetRoad("highway")
		case "city":
			s.game.SetRoad("city")
		case "race":
			s.game.RaceCars()
		case "alfa":
			s.game.BuildAlfa()
		case "bmw":
			s.game.BuildBMW()
		default:
			fmt.Println("Invalid command")
		}
	}
}
