package entity

// this is a port
type Game interface {
	SetRoad(string)
	RaceCars()
	BuildAlfa()
	BuildBMW()
}

type GameState struct {
	cars      []Car
	road      *Road
	displayer Displayer
}

func NewGame(displayer Displayer) Game {
	return &GameState{
		cars:      []Car{},
		road:      NewRaceTrack(),
		displayer: displayer,
	}
}

func (u *GameState) RaceCars() {
	for _, c := range u.cars {
		c.Drive()
	}
}

func (u *GameState) BuildAlfa() {
	alfa := NewAlfa(u.road, u.displayer) // a car must be on a road in order to drive
	u.cars = append(u.cars, alfa)
	u.displayer.ShowNewCar(alfa)
}

func (u *GameState) BuildBMW() {
	bmw := NewBmw(u.road, u.displayer)
	u.cars = append(u.cars, bmw)
	u.displayer.ShowNewCar(bmw)
}

func (u *GameState) SetRoad(roadtype string) {
	if roadtype == "track" {
		u.road = NewRaceTrack()
	} else if roadtype == "highway" {
		u.road = NewHighWay()
	} else if roadtype == "city" {
		u.road = NewCityRoad()
	}

	u.displayer.ShowRoad(*u.road)
}
