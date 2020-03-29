package main

import "fmt"

// Nested Struct
type Configuration struct {
	Env   string
	Proxy Proxy
}

type Proxy struct {
	Address string
	Port    string
}

func main() {
	// Creating an instance of a nested struct
	c := &Configuration{
		Env: "DEBUG:TRUE",
		Proxy: Proxy{
			Address: "addr",
			Port:    "port",
		},
	}
	fmt.Println(c)
	fmt.Println(c.Proxy.Address)
}
