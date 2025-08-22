package mediator

import "fmt"

type Train interface {
	Arrive()
	Depart()
	permitArrival()
}

type PassengerTrain struct {
	Mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if g.Mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}
