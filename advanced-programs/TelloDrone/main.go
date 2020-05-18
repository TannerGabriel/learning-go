package main

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
	"gocv.io/x/gocv"
)

var (
	intensity int = 20
	frameX        = 960
	frameY        = 720
	frameSize     = frameX * frameY * 3
)

func main() {
	keys := keyboard.NewDriver()
	drone := tello.NewDriver("8888")

	window := gocv.NewWindow("Tello")

	ffmpeg := exec.Command("ffmpeg", "-hwaccel", "auto", "-hwaccel_device", "opencl", "-i", "pipe:0",
		"-pix_fmt", "bgr24", "-s", strconv.Itoa(frameX)+"x"+strconv.Itoa(frameY), "-f", "rawvideo", "pipe:1")
	ffmpegIn, _ := ffmpeg.StdinPipe()
	ffmpegOut, _ := ffmpeg.StdoutPipe()

	work := func() {
		keys.On(keyboard.Key, handleKeyboardInput(drone))

		if err := ffmpeg.Start(); err != nil {
			fmt.Println(err)
			return
		}

		drone.On(tello.ConnectedEvent, func(data interface{}) {
			fmt.Println("Connected")
			drone.StartVideo()
			drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
			drone.SetExposure(0)

			gobot.Every(100*time.Millisecond, func() {
				drone.StartVideo()
			})
		})

		drone.On(tello.VideoFrameEvent, func(data interface{}) {
			pkt := data.([]byte)
			if _, err := ffmpegIn.Write(pkt); err != nil {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		work,
	)

	robot.Start()

	// now handle video frames from ffmpeg stream in main thread, to be macOS/Windows friendly
	for {
		buf := make([]byte, frameSize)
		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
			fmt.Println(err)
			continue
		}
		img, _ := gocv.NewMatFromBytes(frameY, frameX, gocv.MatTypeCV8UC3, buf)
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
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
