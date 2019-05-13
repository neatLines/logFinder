package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/neatLines/logFinder/server/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

//
type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//跨域 待删除
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var (
	wg sync.WaitGroup
)

//
func (c *WsController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	ok, quit := make(chan int, 1), make(chan int, 1)
	go initChan(quit, ok)
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		var filterModel models.Filter
		json.Unmarshal(msg, &filterModel)
		quit <- 1
		<-ok
		go sendMsg(ws, quit, ok, filterModel)

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))
	}
}

func sendMsg(conn *websocket.Conn, quit chan int, ok chan int, filterModel models.Filter) {
	fmt.Println("start one")
	x := make(map[string]bool)
	for _, name := range filterModel.Hostname {
		x[name] = true
	}
	withNanos := "2006-01-02 15:04:05"
	startTime, _ := time.Parse(withNanos, filterModel.StartTime)
	endTime, _ := time.Parse(withNanos, filterModel.EndTime)
	if filterModel.NeedFlush {
		endTime, _ = time.Parse(withNanos, "")
	}
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	//设置分区
	partitionList, err := consumer.Partitions(filterModel.AppName)
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	exit := make(chan int, len(partitionList))
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(filterModel.AppName, int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer, exit chan int) {
			for msg := range pc.Messages() {
				select {
				case <-exit:
					return //收到信号就退出线程
				default:
					if x[string(msg.Key)] {
						var msgV models.MQ
						json.Unmarshal(msg.Value, &msgV)
						withNanos := "2006-01-02 15:04:05"
						thisTime, _ := time.Parse(withNanos, msgV.Time)
						if (!endTime.IsZero() && (startTime.After(thisTime) || endTime.Before(thisTime))) || (endTime.IsZero() && startTime.After(thisTime)) {
						} else {
							if filterModel.Filter != "" {
								match := true
								for _, fil := range strings.Split(filterModel.Filter, "\\bb") {
									if ok, _ := regexp.Match(fil, []byte(msgV.Value)); !ok {
										match = false
									}
								}
								if !match {
									break
								}
							}
							ret := models.Message{Message: msgV.Value, Time: msgV.Time}
							err := conn.WriteJSON(ret)
							if err != nil {
								log.Printf("client.WriteJSON error: %v", err)
								conn.Close()
								break
							}
						}
					}
				}
			}

		}(pc, exit)
	}
	//time.Sleep(time.Hour)
	<-quit
	for partition := range partitionList {
		exit <- partition
	}
	ok <- 1
	consumer.Close()
	fmt.Println("consumer closed")
	fmt.Println("exit one")
}

func initChan(quit chan int, ok chan int) {
	for {
		select {
		case <-quit:
			ok <- 1
			return //收到信号就退出线程
		default:
		}
	}
}
