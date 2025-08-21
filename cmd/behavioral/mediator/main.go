package main

import "github.com/barantoraman/DesignPatternsTraining/internal/behavioral/mediator"

func main() {
	manager := mediator.NewStationManger()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: manager,
	}

	freightTrain := &mediator.FreightTrain{
		Mediator: manager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
