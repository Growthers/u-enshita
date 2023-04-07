package api

type WebsocketType interface {
	CreateCommentWebSocketJSON | ChangeSpeakerWebSocketJSON
}

// Websocketでやり取りするとき用の型
type WebsocketResponse[T WebsocketType] struct {
	Type    string `json:"type"`
	Payload T      `json:"payload"`
}
