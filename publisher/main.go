package main

import (
	"fmt"
	"strconv"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	i := 0
	for {
		// Simple Publisher
		t := time.Now()
		msg, err := nc.Request("foo", []byte("Hello World "+strconv.Itoa(i)+" sent in : "+t.Format("2006-01-02 15:04:05")), 5*time.Second)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Sending hello world", i)
		}

		if msg != nil {
			shouldResponse := "Done"
			if string(msg.Data) == shouldResponse {
				fmt.Println("Response as we need : ", i)
			} else {
				fmt.Println("Response is not what we need : ", i)
			}
		} else {
			fmt.Println("No responder, still waiting...")
		}

		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
