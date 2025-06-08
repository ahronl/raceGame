package entity

import (
	"math/rand"
	"time"
)

type AlfaRomeo struct {
	engine    string
	tyres     string
	road      *Road
	displayer Displayer
}

func NewAlfa(r *Road, d Displayer) Car {
	return &AlfaRomeo{
		engine:    "v8",
		tyres:     "nexen",
		road:      r,
		displayer: d,
	}
}

func (car *AlfaRomeo) Drive() {
	speed := car.CalculateHoursePower() * car.road.GetTraction() * car.getRandom()
	car.displayer.ShowCarDrivingAt(speed)
}

func (car *AlfaRomeo) getRandom() float32 {
	rand.Seed(time.Now().UnixNano())
	return float32(5 + rand.Intn(6))
}

func (car *AlfaRomeo) CalculateHoursePower() float32 {

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

func (car *AlfaRomeo) ReplaceTyres(tyres string) {
	car.tyres = tyres
}
