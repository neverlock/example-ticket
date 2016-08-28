package main

import (
	"fmt"

	"github.com/kataras/iris"
)

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
	// for Allow origin you can make use of the middleware
	//iris.Config().Websocket.Headers["Access-Control-Allow-Origin"] = "*"

	var myChatRoom = "room1"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {

		c.Join(myChatRoom)

		c.On("chat", func(message string) {
			// to all except this connection ->
			//c.To(websocket.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message)

			// to the client ->
			//c.Emit("chat", "Message from myself: "+message)

			//send the message to the whole room,
			//all connections are inside this room will receive this message
			//c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
			js := fmt.Sprintf("{\"From\":\"%s\",\"MSG\":\"%s\"}", c.ID(), message)
			c.To(myChatRoom).Emit("chat", js)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8080")
}

func (u BookingAPI) Get() {
	u.Write("Get from /booking")
}
