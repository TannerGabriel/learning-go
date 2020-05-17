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

		if key == keyboard.A {
			fmt.Println("Going left.")
			drone.Left(intensity)
		} else if key == keyboard.D {
			fmt.Println("Going right.")
			drone.Right(intensity)
		} else if key == keyboard.W {
			fmt.Println("Going up.")
			drone.Up(intensity)
		} else if key == keyboard.S {
			fmt.Println("Going down.")
			drone.Down(intensity)
		} else if key == keyboard.U {
			fmt.Println("Going forward.")
			drone.Forward(intensity)
		} else if key == keyboard.J {
			fmt.Println("Going backward.")
			drone.Backward(intensity)
		} else if key == keyboard.K {
			fmt.Println("Rotating counter-clockwise.")
			drone.CounterClockwise(intensity)
		} else if key == keyboard.H {
			fmt.Println("Rotating clockwise.")
			drone.Clockwise(intensity)
		} else if key == keyboard.L {
			fmt.Println("Landing drone")
			drone.Land()
		} else if key == keyboard.T {
			fmt.Println("Take off drone")
			drone.TakeOff()
		} else {
			resetDronePostion(drone)
		}
	}
}
