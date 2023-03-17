package ws

import (
	"context"
	"github.com/growthers/mu-enshita/pkg/types/api"
	"github.com/labstack/echo/v4"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func CommentHandler(c echo.Context) {
	conn, err := websocket.Accept(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan api.CreateCommentWebSocketJSON, 8)
	defer conn.Close(websocket.StatusInternalError, "")
	CommentBroadCast.AppendChan(ch)

	ctx, cancel := context.WithCancel(context.Background())
	//受信処理
	go func() {
		var comment api.CreateCommentRequestJSON
		for {
			err := wsjson.Read(ctx, conn, comment)
			//TODO 切断処理をちゃんとかく
			if err != nil {
				CommentBroadCast.CloseChan(ch)
				conn.Close(websocket.StatusInternalError, "")
				//チャンネルを削除
				cancel()
				break
			}

			commentData := api.CreateCommentWebSocketJSON{
				From:        comment.From,
				DataSource:  comment.DataSource,
				CommentText: comment.CommentText,
			}

			CommentBroadCast.SendComment(commentData)
		}
	}()
	//送信処理
	go func() {
		for {
			select {
			case comment := <-ch:
				err := wsjson.Write(ctx, conn, comment)
				if err != nil {
					conn.Close(websocket.StatusInternalError, "")
					CommentBroadCast.CloseChan(ch)
					cancel()
					break
				}
			}
		}

	}()
	for {
		select {
		case <-ctx.Done():
			break

		}
	}
}
