package entity

type Car interface {
	Drive()
	CalculateHoursePower() float32
	ReplaceTyres(tyres string)
}

// this is a port
type Displayer interface {
	ShowCarDrivingAt(speed float32)
	ShowNewCar(car Car)
}
