package controllers

import (
	"fmt"
	"log"
	"net/http"

	"time"

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
		quit <- 1
		<-ok
		go sendMsg(ws, quit, ok)

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))
	}
}

func sendMsg(conn *websocket.Conn, quit chan int, ok chan int) {
	fmt.Println("start one")
	for {

		select {
		case <-quit:
			ok <- 1
			fmt.Println("exit one")
			return //收到信号就退出线程
		default:

			//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
			time.Sleep(time.Second * 3)
			msg := models.Message{Message: "这是向页面发送的数据 ", Time: time.Now().Format("2006-01-02 15:04:05")}
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				conn.Close()
				break
			}
		}
	}
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