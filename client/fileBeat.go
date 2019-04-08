package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
)

var filename = flag.String("f", "", "filename")

var url = flag.String("u", "", "url")

var appname = flag.String("n", "", "appname")

var host, _ = os.Hostname()

var hostname = flag.String("m", host, "hostname")

func produceMessage(url string, topicName string, hostname string, msgChan chan string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{url}, config)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topicName,
		Partition: int32(-1),
		Key:       sarama.StringEncoder(hostname),
	}

	for message := range msgChan {
		msg.Value = sarama.ByteEncoder(message)
		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}
}

func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("Listening %s, Send to %s, appname: %s", *filename, *url, *appname))
	if t, err := tail.TailFile(*filename, tail.Config{Follow: true}); err != nil {
		fmt.Println("tail failed", err)
		return
	} else {
		defer t.Done()
		msgChan := make(chan string, 100)
		go produceMessage(*url, *appname, *hostname, msgChan)
		for line := range t.Lines {
			msgChan <- line.Text
		}
		<-msgChan
	}
}
