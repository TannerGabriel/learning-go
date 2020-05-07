package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect("nats://localhost:4444/")

	if err != nil {
		fmt.Println(err)
	}

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Responding to a request message
	nc.Subscribe("help", func(m *nats.Msg) {
		fmt.Println("Help")
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	// sub, _ := nc.SubscribeSync("foo")
	// timeout, _ := time.ParseDuration("30s")
	// m, _ := sub.NextMsg(timeout)
	// fmt.Println(m)

	// // Channel Subscriber
	// ch := make(chan *nats.Msg, 64)
	// sub, _ := nc.ChanSubscribe("foo", ch)
	// msg := <-ch
	// fmt.Println(msg)

	// Unsubscribe
	// sub.Unsubscribe()

	// Drain
	// sub.Drain()

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Requests
	msg, _ := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	fmt.Printf("Msg received on [%s] : %s\n", msg.Subject, string(msg.Data))

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()

}
