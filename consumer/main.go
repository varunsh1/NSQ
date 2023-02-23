package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	// print consumed message body
	fmt.Println(string(m.Body))

	return nil
}

func main() {
	// cmd flag initialisation
	topic := flag.String("topic", "topic", "used to get consumer topic from flag")
	lookupd := flag.String("lookupd", "lookupd", "used to get nsq lookupd from flag")
	flag.Parse()

	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer(*topic, "test_consumer", config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer
	consumer.AddHandler(&myMessageHandler{})

	// Use nsqlookupd to discover nsqd instances.
	err = consumer.ConnectToNSQLookupd(*lookupd)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
