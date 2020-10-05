package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("nsqd:4150", config)
	if err != nil {
		log.Fatalln("create nsq producer fail")
		return
	}

	for {
		time.Sleep(time.Second)
		now := time.Now().String()

		// Synchronously publish a single message to the specified topic.
		// Messages can also be sent asynchronously and/or in batches.
		if err := p.Publish("MY_TOPIC", []byte(now)); err != nil {
			log.Println("nsq publish fail")
		}
	}
}
