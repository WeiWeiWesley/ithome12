package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	c, err := nsq.NewConsumer("MY_TOPIC", "ch1", config)
	if err != nil {
		log.Fatalln("create nsq consumer fail", err.Error())
		return
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	c.AddHandler(&myMessageHandler{})

	err = c.ConnectToNSQLookupd("nsqlookupd:4161")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer c.Stop()

	select {}
}

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func processMessage(msg []byte) error {
	log.Println(string(msg))
	return nil
}
