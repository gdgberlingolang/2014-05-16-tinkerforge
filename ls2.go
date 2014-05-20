	// create a subscriber enumerate (call for all devices)
	subscriber := enumerate.NewEnumerate("enum", func(r *enumerate.EnumerateResult, err error) {
		if handleError(err) {
			return
		}
		fmt.Printf("%s\n", r.String())
	})
	// attach subscriber to the manager
	err = manager.Subscribe(subscriber, "local")
	// go on with the program, waiting for a key
	fmt.Printf("Press return for stop.\n")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()
}

func handleError(err error) bool {
	if err != nil {
		fmt.Printf("Error occur: %s\n", err.Error())
		return true
	}
	return false
}
