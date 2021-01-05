package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += " Tire_C +"
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += " Tire_P +"
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += " Engine_C +"
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += " Engine_P +"
			outPlaneChan <- plane
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car   " + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane " + strconv.Itoa(i)}
		i++
	}
}

func main() {
	Carchan1 := make(chan Car)
	Carchan2 := make(chan Car)
	Carchan3 := make(chan Car)

	Planechan1 := make(chan Plane)
	Planechan2 := make(chan Plane)
	Planechan3 := make(chan Plane)

	go StartCarWork(Carchan1)
	go StartPlaneWork(Planechan1)

	go MakeTire(Carchan1, Planechan1, Carchan2, Planechan2)
	go MakeEngine(Carchan2, Planechan2, Carchan3, Planechan3)

	for {
		select {
		case result := <-Carchan3:
			fmt.Println(result.val)
		case result := <-Planechan3:
			fmt.Println(result.val)
		}
	}
}
