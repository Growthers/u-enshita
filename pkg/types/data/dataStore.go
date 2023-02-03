package data

import "time"

type StoreJSON struct {
	Event struct {
		StartAt time.Time `json:"startAt"`
		EndAt   time.Time `json:"endAt"`
	} `json:"event"`
	Speakers []struct {
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
	} `json:"speakers"`
}
