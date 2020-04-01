package eventdispatcher

import (
	"fmt"
	"testing"
	"time"
)

type TimeoutEvent struct {
	Name     string
	Url      string
	Duration time.Duration
}

type TimeoutEventData struct {
	Url      string
	Duration time.Duration
}

func (e TimeoutEvent) GetName() string {
	return e.Name
}

func (e TimeoutEvent) GetData() interface{} {
	return TimeoutEventData{
		Url:      e.Url,
		Duration: e.Duration,
	}
}

func TestEventDispatcher_AddListener(t *testing.T) {
	e := TimeoutEvent{
		Name:     "event!timeout",
		Url:      "/worker/get-products",
		Duration: 1 * time.Second,
	}

	listener1 := NewListener(func(event EventInterface) error {
		dataInterface := event.GetData()
		data := dataInterface.(TimeoutEventData)
		fmt.Println(data.Url)
		return nil
	}, false)

	AddListener(e.Name, listener1)
	if Dispatcher(e) != nil {
		t.Fatalf("出现错误")
	}
}
