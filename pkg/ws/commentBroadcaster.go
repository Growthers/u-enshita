package ws

import (
	"github.com/growthers/mu-enshita/pkg/types/api"
	"sync"
)

var CommentBroadcast CommentBroadcaster

type CommentBroadcaster struct {
	mu       sync.Mutex
	nodeList map[chan api.CreateCommentWebSocketJSON]struct{}
}

func (c *CommentBroadcaster) SendComment(comment api.CreateCommentWebSocketJSON) {
	for k, _ := range c.nodeList {
		k <- comment
	}
}

func (c *CommentBroadcaster) AppendChan(ch chan api.CreateCommentWebSocketJSON) {
	c.mu.Lock()
	c.nodeList[ch] = struct{}{}
	c.mu.Unlock()
}

// CloseChan もしかしたらMutexを挟むべき？
func (c *CommentBroadcaster) CloseChan(ch chan api.CreateCommentWebSocketJSON) {
	c.mu.Lock()
	delete(c.nodeList, ch)
	c.mu.Unlock()
}
