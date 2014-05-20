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
