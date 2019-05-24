package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type KAFKAMSG struct {
	Value   string
	Time    string
	CPUINFO float64
	MEMINFO float64
}

var filename = flag.String("f", "", "filename")

var url = flag.String("u", "", "url")

var webPort = flag.String("wp", "", "webport")

var kafkaPort = flag.String("kp", "", "kafkaport")

var appname = flag.String("n", "", "appname")

var host, _ = os.Hostname()

var hostname = flag.String("m", host, "hostname")

var token = flag.String("t", "", "token")

func produceMessage(url string, topicName string, hostname string, msgChan chan string) {
	v, _ := mem.VirtualMemory()
	cc, _ := cpu.Percent(time.Second, false)
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			<-ticker.C
			v, _ = mem.VirtualMemory()
			cc, _ = cpu.Percent(time.Second, false)
		}
	}()
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
		km := KAFKAMSG{message, time.Now().Format("2006-01-02 15:04:05"), cc[0], v.UsedPercent}
		s, _ := json.Marshal(km)
		msg.Value = sarama.ByteEncoder(s)
		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}
}

//
func httpPostJSON(appname string, hostname string, urlbase string) {
	jsonStr := []byte(`{ "appname": "` + appname + `", "hostname": "` + hostname + `" }`)
	url := "http://" + urlbase + "/v1/hosts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// handle error
	}

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statuscode)
	fmt.Println(hea)

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
		go httpPostJSON(*appname, *hostname, *url+":"+*webPort)
		go produceMessage(*url+":"+*kafkaPort, *appname, *hostname, msgChan)
		for line := range t.Lines {
			msgChan <- line.Text
		}
		<-msgChan
	}
}
