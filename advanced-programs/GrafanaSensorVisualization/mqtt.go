package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername("mqtt")
	password := ""
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	uri, err := url.Parse("tcp://10.0.0.22:1883")
	if err != nil {
		log.Fatal(err)
	}
	topic := "sensors"

	go listen(uri, topic)

	client := connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		fmt.Println(t)
		var min int64 = 10
		var max int64 = 100
		var random int64 = (rand.Int63n(max-min) + min)
		nsec := time.Now().UnixNano()
		payload := "weather,location=us-midwest temperature=" + strconv.FormatInt(random, 10) + " " + strconv.FormatInt(nsec, 10)
		client.Publish("sensors", 0, false, payload)
	}
}
