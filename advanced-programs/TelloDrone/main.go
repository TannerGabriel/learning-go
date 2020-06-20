package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

var (
	intensity int = 20
)

func main() {
	keys := keyboard.NewDriver()
	drone := tello.NewDriver("8888")

	work := func() {
		keys.On(keyboard.Key, handleKeyboardInput(drone))
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		work,
	)

	robot.Start()

}

func resetDronePostion(drone *tello.Driver) {
	drone.Forward(0)
	drone.Backward(0)
	drone.Up(0)
	drone.Down(0)
	drone.Left(0)
	drone.Right(0)
	drone.Clockwise(0)
}

func handleKeyboardInput(drone *tello.Driver) func(interface{}) {
	return func(data interface{}) {
		key := data.(keyboard.KeyEvent).Key

		switch key {
		case keyboard.A:
			fmt.Println("Going left.")
			drone.Left(intensity)
		case keyboard.D:
			fmt.Println("Going right.")
			drone.Right(intensity)
		case keyboard.W:
			fmt.Println("Going up.")
			drone.Up(intensity)
		case keyboard.S:
			fmt.Println("Going down.")
			drone.Down(intensity)
		case keyboard.U:
			fmt.Println("Going forward.")
			drone.Forward(intensity)
		case keyboard.J:
			fmt.Println("Going backward.")
			drone.Backward(intensity)
		case keyboard.K:
			fmt.Println("Rotating counter-clockwise.")
			drone.CounterClockwise(intensity)
		case keyboard.H:
			fmt.Println("Rotating clockwise.")
			drone.Clockwise(intensity)
		case keyboard.L:
			fmt.Println("Landing drone")
			drone.Land()
		case keyboard.T:
			fmt.Println("Take off drone")
			drone.TakeOff()
		default:
			resetDronePostion(drone)
		}
	}
}
