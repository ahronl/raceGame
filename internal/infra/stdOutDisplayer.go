package infra

import (
	"example/hello/internal/entity"
	"fmt"
)

type stdOut struct{}

func NewStdoutDisplayer() *stdOut {
	return &stdOut{}
}

func (s *stdOut) ShowCarDrivingAt(speed float32) {
	fmt.Printf("car is speeding at %v \n", speed)
}

func (s *stdOut) ShowNewCar(car entity.Car) {
	fmt.Printf("a new car was born %+v \n", car)
}
