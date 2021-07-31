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

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		t := time.Now()
		fmt.Printf("%s when %s\n", t.Format("2006-01-02 15:04:05"), (m.Data))
	})

	time.Sleep(1 * time.Minute)
}
