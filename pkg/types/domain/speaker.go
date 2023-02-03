package domain

type Speaker struct {
	Name   string
	Title  string
	Social []struct {
		ServiceName string
		AccountName string
	}
	Duration  uint64
	BreakTime bool
	Order     uint64
	Cancel    bool
}

func NewSpeaker(name string, title string, social []struct {
	ServiceName string
	AccountName string
}, duration uint64, breakTime bool, order uint64, cancel bool) *Speaker {
	return &Speaker{Name: name, Title: title, Social: social, Duration: duration, BreakTime: breakTime, Order: order, Cancel: cancel}
}
