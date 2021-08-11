package main

import (
	"flag"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	servers := "nats://localhost:4222"
	nc, err := nats.Connect(servers)

	if err != nil {
		fmt.Println("Some error happen")
		return
	}

	scaleNumber := flag.String("scale", "1", "--")
	flag.Parse()

	// Simple Async Subscriber
	nc.QueueSubscribe("foo", "group_type", func(m *nats.Msg) {
		fmt.Println("Here :", *scaleNumber, "from", m.Sub.Queue, string(m.Data))
		m.Respond([]byte("Number of scale : " + *scaleNumber))
	})

	time.Sleep(4 * time.Minute)
}
