package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 0h1m", func() { sayHelloTo("Someone!") })
	c.Start()
	time.Sleep(1 * time.Minute)
}

func sayHelloTo(name string) string {
	return "Hello " + name
}
