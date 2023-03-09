package ws

import "nhooyr.io/websocket"

type CommentHandler struct {
	conn []*websocket.Conn
}

func (c *CommentHandler) SendComment() error {
	return nil
}

func (c *CommentHandler) AppendWS(ws *websocket.Conn) {

	c.conn = append(c.conn, ws)
}

func (c *CommentHandler) CloseWS(i int) {
	c.conn[i] = c.conn[len(c.conn)-1]
	c.conn = c.conn[:len(c.conn)-1]
}
