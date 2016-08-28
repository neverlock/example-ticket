package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris"
	"github.com/neverlock/utility/random"
	"golang.org/x/net/websocket"
)

/* Native messages no need to import the iris-ws.js to the ./templates.client.html
Use of: OnMessage and EmitMessage
*/

type clientPage struct {
	Title string
	Host  string
}

type BookingAPI struct {
	*iris.Context
}

func main() {

	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})
	iris.API("/booking", BookingAPI{})

	// the path which the websocket client should listen/registed to ->
	iris.Config.Websocket.Endpoint = "/my_endpoint"

	ws := iris.Websocket // get the websocket server

	ws.OnConnection(func(c iris.WebsocketConnection) {

		c.OnMessage(func(data []byte) {
			message := string(data)
			//c.To(iris.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + message))
			//c.EmitMessage([]byte("Me: " + message))
			height := random.Int(1, 100)
			width := random.Int(1, 300)

			js := fmt.Sprintf("{\"From\":\"%s\",\"H\":%d,\"W\":%d,\"MSG\":\"%s\"}", c.ID(), height, width, message)
			c.To(iris.Broadcast).EmitMessage([]byte(js))
			c.EmitMessage([]byte(js))
			//c.To(myChatRoom).Emit("chat", js)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8080")
}

func (u BookingAPI) Get() {
	u.Write("Get from /booking")
	origin := "http://104.238.149.36:8080/"
	url := "ws://104.238.149.36:8080/my_endpoint"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	/*
		Fm := "Test from"
		Msg := "dummy msg"
		js1 := fmt.Sprintf("{\"From\":\"%s\",\"MSG\":\"%s\"}", Fm, Msg)
	*/
	//height := random.Int(1, 100)
	//width := random.Int(1, 300)

	//js := fmt.Sprintf("{\"From\":\"%s\",\"H\":%d,\"W\":%d,\"MSG\":\"%s\"}", c.ID(), height, width, message)

	if _, err := ws.Write([]byte("")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
