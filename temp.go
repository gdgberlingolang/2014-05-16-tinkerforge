package main

import (
	"bricker"
	"bricker/device/bricklet/temperature"
	"bricker/manager/chanconnection"
	"fmt"
)

func main() {
	quit := make(chan struct{}) // channel for done

	// creates a manager (event handler etc.)
	manager := chanconnection.NewChanConnection()
	defer manager.Done()

	// create a bricker (communication with brickd)
	bricker, err := bricker.NewBrickerUnbuffered("localhost:4223")
	if err != nil {
		// TODO: handle error
		fmt.Printf("No connection: %s\n", err.Error())
		return
	}
	defer bricker.Done()

	// attach the bricker to the manger
	err = manager.AddBricker(bricker, "local")
	if err != nil { // no bricker, no fun
		// TODO: handle error
		fmt.Printf("No connection: %s\n", err.Error())
		return
	}

	// create a subscriber gettemperature (call for the temperature once)
	subscriber := temperature.NewGetTemperature("gettemp", 42362, func(r *temperature.TemperatureResult, err error) {
		defer close(quit)
		if err != nil { // no bricker, no fun
			// TODO: handle error
			fmt.Printf("Error occur: %s\n", err.Error())
			return
		}
		fmt.Printf("%s\n", r.String())
	})

	// attach subscriber to the manager
	err = manager.Subscribe(subscriber, "local")

	// endless loop, waiting for closing the quit channel
	for {
		select {
		case <-quit:
			return
		}
	}
}
