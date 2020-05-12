package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()

	// Funcs are automatically invoked in their own goroutine, asynchronously
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("TZ=America/New_York 30 04 * * * *", func() { fmt.Println("Runs at 04:30 New York time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Runs every hour") })
	c.AddFunc("@every 0h0m1s", func() { sayHelloTo("Someone!") })
	c.Start()

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Added time to see output
	time.Sleep(10 * time.Second)

	c.Stop() // Stop the scheduler (does not stop any jobs already running)

}

func sayHelloTo(name string) {
	fmt.Println("Hello ", name)
}
