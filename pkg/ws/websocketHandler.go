package ws

import (
	"context"
	"github.com/growthers/mu-enshita/pkg/types/api"
	"github.com/labstack/echo/v4"
	"log"
	"nhooyr.io/websocket"
)

func WebSocketHandler(c echo.Context) {
	conn, err := websocket.Accept(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan api.CreateCommentWebSocketJSON, 8)
	defer conn.Close(websocket.StatusInternalError, "")
	CommentBroadcast.AppendChan(ch)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//受信処理

	go WebsocketReceive(ctx, conn)
}

func WebsocketReceive(ctx context.Context, c *websocket.Conn) error {

	return nil
}

func WebsocketSend() error {
	return nil
}
