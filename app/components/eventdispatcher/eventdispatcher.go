package eventdispatcher

import "fmt"

var eventDispatcher *EventDispatcher

func init() {
	eventDispatcher = &EventDispatcher{
		listeners: make(map[string][]*Listener),
	}
}

type EventDispatcher struct {
	listeners map[string][]*Listener
}

func Dispatcher(event EventInterface) error {
	eventName := event.GetName()
	listeners, ok := eventDispatcher.listeners[eventName]
	if !ok || len(listeners) == 0 {
		return fmt.Errorf("no listeners. eventName=%s", eventName)
	}
	return eventDispatcher.callListeners(event, listeners)
}

func (ed *EventDispatcher) callListeners(event EventInterface, listeners []*Listener) error {
	asyncListners := []*Listener{}
	for _, listener := range listeners {
		if listener.async {
			asyncListners = append(asyncListners, listener)
		} else {
			err := listener.onListen(event)
			if err != nil {
				return err
			}
		}
	}

	eventName := event.GetName()
	for index, listener := range asyncListners {
		go func(l *Listener, eventName string, index int) {
			err := l.onListen(event)
			if err != nil {
				Warnf("on onListen error. eventName=%s||listener index=%d", event.GetName(), index)
			}
		}(listener, eventName, index)
	}

	return nil
}

//添加listener到event-dispatcher
func AddListener(eventName string, listener *Listener) {
	if _, ok := eventDispatcher.listeners[eventName]; !ok {
		eventDispatcher.listeners[eventName] = make([]*Listener, 0)
	}
	eventDispatcher.listeners[eventName] = append(eventDispatcher.listeners[eventName], listener)
}
