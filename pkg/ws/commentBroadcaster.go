package ws

import (
	"github.com/growthers/mu-enshita/pkg/types/api"
)

var CommentBroadCast CommentBroadcaster

type CommentBroadcaster struct {
	nodeList map[chan api.CreateCommentWebSocketJSON]struct{}
}

func (c *CommentBroadcaster) SendComment(comment api.CreateCommentWebSocketJSON) {
	for k, _ := range c.nodeList {
		k <- comment
	}
}

func (c *CommentBroadcaster) AppendChan(ch chan api.CreateCommentWebSocketJSON) {
	c.nodeList[ch] = struct{}{}
}

// もしかしたらMutexを挟むべき？
func (c *CommentBroadcaster) CloseChan(ch chan api.CreateCommentWebSocketJSON) {
	delete(c.nodeList, ch)
}
