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
			delivered, err := msg.Sub.Delivered()
			if err != nil {
				fmt.Println("Delivered error message in : ", err.Error())
			}
			fmt.Println("Delivered message in : ", delivered, "subscribers")

			fmt.Println("Queue in :", msg.Sub.Queue)
			fmt.Println("Type in :", msg.Sub.Type())
			fmt.Println("Reply in :", msg.Reply)
			fmt.Println("Header in :", msg.Header)
			fmt.Println("Subject in :", msg.Sub.Subject)

			shouldResponse := "Done"
			if string(msg.Data) == shouldResponse {
				fmt.Println("Response as we need : ", i)
			} else {
				fmt.Println("Response is not what we need : ", i)
			}
		} else {
			fmt.Println("No responder, still waiting...")
		}

		fmt.Println("")
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
