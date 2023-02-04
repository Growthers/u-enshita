package data

type StoreJSON struct {
	Event    EventStoreJSON     `json:"event"`
	Speakers []SpeakerStoreJSON `json:"speakers"`
}

type EventStoreJSON struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

type SpeakerStoreJSON struct {
	ID     string `json:"id"`
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
