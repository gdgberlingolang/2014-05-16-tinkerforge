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
