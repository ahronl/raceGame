package entity

type Road struct {
	traction float32
}

func NewRaceTrack() *Road {
	return &Road{traction: 0.9}
}

func NewHighWay() *Road {
	return &Road{traction: 0.5}
}

func NewCityRoad() *Road {
	return &Road{traction: 0.2}
}

func (r *Road) GetTraction() float32 {
	return r.traction
}
