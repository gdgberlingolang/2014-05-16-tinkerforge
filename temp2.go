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
		fmt.Printf("Could not attach bricker: %s\n", err.Error())
		return
	}
