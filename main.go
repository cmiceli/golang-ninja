package main

import (
	"fmt"
	"github.com/ninjasphere/go-ninja/logger"

	"github.com/ninjasphere/driver-block/arduino"
)

var log = logger.GetLogger("PiCrust")

var path = "/dev/ttyAMA0"
var speed = 9600

func arduinoPrinter(a arduino.DeviceData) {
	fmt.Printf("Got data: %s %d %d %v\n", a.G, a.V, a.D, a)
}

func main() {

	_, err := arduino.Connect(path, speed, arduinoPrinter)

	if err != nil {
		fmt.Printf("Couldn't connect to arduino: %s\n", err)
	}
}
