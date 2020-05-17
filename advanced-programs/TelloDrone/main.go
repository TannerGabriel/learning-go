package main

import (
	"time"

	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	// "gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	// keys := keyboard.NewDriver()
	drone := tello.NewDriver("8888")

	work := func() {
		// keys.On(keyboard.Key, handleKeyboardInput(drone))

		//TakeOff the Drone.
		gobot.After(5*time.Second, func() {
			drone.TakeOff()
			fmt.Println("Tello Taking Off...")
		})

		//Land the Drone.
		gobot.After(15*time.Second, func() {
			drone.Land()
			fmt.Println("Tello Landing...")
		})

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()

}
