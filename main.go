package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"

	"github.com/neverlock/utility/random"

	wsx "golang.org/x/net/websocket"
)

/* Native messages no need to import the iris-ws.js to the ./templates.client.html
Use of: OnMessage and EmitMessage
*/

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()

	app.StaticWeb("/js", "./static/js")

	app.Get("/", func(ctx iris.Context) {
		ctx.View("client.html", clientPage{"Client Page", ctx.Host()})
	})

	// https://github.com/kataras/iris/tree/master/_examples/README.md#mvc
	mvc.New(app.Party("/booking")).Handle(new(BookingController))

	// https://github.com/kataras/iris/blob/master/_examples/README.md#websockets
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
	}) // get the websocket server

	// the path which the websocket client should listen/registed to ->
	app.Any("/my_endpoint", ws.Handler())

	ws.OnConnection(func(c websocket.Connection) {

		c.OnMessage(func(data []byte) {
			message := string(data)
			//c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + message))
			//c.EmitMessage([]byte("Me: " + message))
			height := random.Int(1, 100)
			width := random.Int(1, 300)

			js := fmt.Sprintf("{\"From\":\"%s\",\"H\":%d,\"W\":%d,\"MSG\":\"%s\"}", c.ID(), height, width, message)
			c.To(websocket.Broadcast).EmitMessage([]byte(js))
			c.EmitMessage([]byte(js))
			//c.To(myChatRoom).Emit("chat", js)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	app.Run(iris.Addr(":8080"))
}

type BookingController struct{}

func (c *BookingController) Get() string {
	origin := "http://104.238.149.36:8080/"
	url := "ws://104.238.149.36:8080/my_endpoint"
	ws, err := wsx.Dial(url, "", origin)
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

	return "Get from /booking"
}
