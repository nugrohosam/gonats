package main

import (
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("Some error happen")
		return
	}

	sub, errSub := nc.SubscribeSync("foo")
	if errSub != nil {
		fmt.Printf("Err sub\n")
	}

	m, errNextMsg := sub.NextMsg(1 * time.Millisecond)
	if errNextMsg != nil {
		fmt.Printf("Err next msg\n")
	}

	fmt.Println(m.Metadata())

	time.Sleep(1 * time.Minute)
}
