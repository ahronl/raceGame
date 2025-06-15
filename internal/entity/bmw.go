package entity

import (
	"math/rand"
	"time"
)

type Bmw struct {
	engine    string
	tyres     string
	road      *Road
	displayer Displayer
}

func NewBmw(r *Road, d Displayer) Car {
	return &Bmw{
		engine:    "v6",
		tyres:     "michelin",
		road:      r,
		displayer: d,
	}
}

func (car *Bmw) Drive() {
	speed := car.CalculateHoursePower() * car.road.GetTraction() * car.getRandom()
	car.displayer.ShowCarDrivingAt(speed)
}

func (car *Bmw) getRandom() float32 {
	rand.Seed(time.Now().UnixNano())
	return float32(5 + rand.Intn(6))
}

func (car *Bmw) CalculateHoursePower() float32 {

	engineScore := 1
	tyresScore := 1

	if car.engine == "v8" {
		engineScore = 8
	}

	if car.tyres == "nexen" {
		tyresScore = 2
	}

	if car.tyres == "michelin" {
		tyresScore = 10
	}

	power := engineScore + tyresScore
	return float32(power)
}

func (car *Bmw) ReplaceTyres(tyres string) {
	car.tyres = tyres
}
