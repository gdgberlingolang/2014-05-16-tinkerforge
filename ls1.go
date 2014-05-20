package main
import (
	"bricker"
	"bricker/device/enumerate"
	"bricker/manager/chanconnection"
	"bufio"
	"fmt"
	"os"
)
func main() {
	manager := chanconnection.NewChanConnection()
	defer manager.Done()
	bricker, err := bricker.NewBrickerUnbuffered("localhost:4223")
	if handleError(err) {
		return
	}
	defer bricker.Done()
	err = manager.AddBricker(bricker, "local")
	if handleError(err) {
		return
	}

