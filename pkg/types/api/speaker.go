package api

type GetSpeakerResponseJSON struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Social []struct {
		ServiceName string `json:"serviceName"`
		AccountName string `json:"accountName"`
	} `json:"social"`
	Duration  uint64 `json:"duration"`
	BreakTime bool   `json:"breakTime"`
	Order     uint64 `json:"order"`
	Cancel    bool   `json:"cancel"`
}

type UpdateSpeakerRequestJSON struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Social []struct {
		ServiceName string `json:"serviceName"`
		AccountName string `json:"accountName"`
	} `json:"social"`
	Duration  uint64 `json:"duration"`
	BreakTime bool   `json:"breakTime"`
	Order     uint64 `json:"order"`
	Cancel    bool   `json:"cancel"`
}

type UpdateSpeakerResponseJSON struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Social []struct {
		ServiceName string `json:"serviceName"`
		AccountName string `json:"accountName"`
	} `json:"social"`
	Duration  uint64 `json:"duration"`
	BreakTime bool   `json:"breakTime"`
	Order     uint64 `json:"order"`
	Cancel    bool   `json:"cancel"`
}

type ChangeSpeakerWebSocketJSON struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Social []struct {
		ServiceName string `json:"serviceName"`
		AccountName string `json:"accountName"`
	} `json:"social"`
	Duration  uint64 `json:"duration"`
	BreakTime bool   `json:"breakTime"`
}
